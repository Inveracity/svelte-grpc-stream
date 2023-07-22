import { get } from 'svelte/store'
import { env } from '$env/dynamic/public';

import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { persisted } from 'svelte-local-storage-store'
import { DateTime } from 'luxon';

import { messages } from '$lib/stores/messages';
import { status } from '$lib/stores/status';

import { ChatServiceClient } from '$lib/proto/chat/v1/chat.client';
import type { ChatMessage } from '$lib/proto/chat/v1/chat';
import type { Message, OutgoingMessage } from './types';
import { currentUser, pb } from './pocketbase';
import { channels } from './stores/channel';
import { users, type User } from './stores/users';

export const chat_cache = persisted(
  'chatmessages', // storage
  { lastTs: '0' }, // default value
  { storage: 'session' } // options
)

const transport = new GrpcWebFetchTransport({
  baseUrl: env.PUBLIC_RELAY_URL
});

let controller = new AbortController();

export const Connect = async (serverId: string, userId: string, timestamp: string) => {
  // While the connection is attempting to open, let the UI show a pending state
  status.pending();

  // If the client disconnected, the abort controller is no longer valid and a new one must be created
  if (controller.signal.aborted) {
    controller = new AbortController();
  }

  // The abort controller is used to signal the server to close the stream
  const opts = transport.mergeOptions({ abort: controller.signal, meta: { jwt: pb.authStore.token } });

  // Get the last timestamp from the cache
  const lastTs = get(chat_cache).lastTs

  // Create a new subscription to the server
  const sub = new ChatServiceClient(transport).connect({
    serverId: serverId,
    userId: userId,
    lastTs: timestamp,
  }, opts);


  // If the connection fails, let the UI show an error state
  sub.status.catch((e: Error) => {
    status.error(e.message);
  });

  // Listen for messages from the server
  try {
    for await (const msg of sub.responses) {
      // Filter out messages that should not be written to the UI
      if (filtered(msg, lastTs)) {
        continue;
      }

      // Format timestamp
      const ts = timestampToDate(msg.ts)
      const message: Message = {
        message: msg.text,
        timestamp: ts,
        channel: msg.channelId,
        user: msg.userId,
      }

      messages.add(message);
      chat_cache.set({ lastTs: msg.ts })

    }
  } catch (_) {
    // Stream closed, force user to log in again
    currentUser.set(null);
  }

  // await sub.headers;
  // await sub.trailers;
  chat_cache.set({ lastTs: "0" })
  status.disconnected();
};

// The client can actively Disconnect letting the server know to close the stream
export const Disconnect = async () => {
  chat_cache.set({ lastTs: "0" })
  messages.reset();
  controller.abort();
  users.upd({ name: pb.authStore.model?.name, presence: false });
};

export const SendMessage = (msg: OutgoingMessage) => {
  const client = new ChatServiceClient(transport);
  const opts = transport.mergeOptions({ meta: { jwt: pb.authStore.token } })
  const request: ChatMessage = {
    channelId: msg.channelId,
    userId: msg.userId,
    text: msg.text,
    ts: "0", // The server will set the timestamp
  };

  client.send(request, opts).then((_) => {
    // nothing
  }).catch((e) => {
    console.log(e);
  });
}

const filtered = (msg: ChatMessage, lastTs: string): boolean => {
  // Do not write server messages to the UI
  if (msg.userId === "server" && msg.text === "connected") {
    status.connected();
    return true;
  }

  if (msg.channelId === "system") {
    return filter_system_messages(msg);
  }

  // Do not write messages with timestamp 0 to the UI
  // Deduplicate messages with the same timestamp
  // Ignore all other server messages
  if (msg.ts === "0" || msg.ts === lastTs || msg.userId === "server") {
    return true;
  }

  return false
}

const timestampToDate = (timestamp: string): string => {
  try {
    const nano = parseInt(timestamp)
    return DateTime.fromMillis(nano / 1000000).toFormat("HH:mm")
  } catch (e) {
    console.log(e);
    return timestamp;
  }
}

const filter_system_messages = (msg: ChatMessage): boolean => {
  // Tell UI to show new channel when another user adds one
  if (msg.text.startsWith("channel_add") && msg.userId !== pb.authStore.model?.name) {
    const channel_name = msg.text.split(" ")[1]
    channels.add(channel_name);
  }

  if (msg.text.startsWith("connected")) {
    const user: User = { name: msg.userId, presence: true }
    users.upd(user);
  }

  if (msg.text.startsWith("disconnected")) {
    const user: User = { name: msg.userId, presence: false }
    users.upd(user);
    messages.add({ channel: "general", message: `${msg.userId} disconnected`, timestamp: "0", user: "server" })
  }

  return true;
}
