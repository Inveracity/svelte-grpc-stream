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

export const chat_cache = persisted(
  'chatmessages', // storage
  { lastTs: '0' }, // default value
  { storage: 'session' } // options
)

const transport = new GrpcWebFetchTransport({
  baseUrl: env.PUBLIC_RELAY_URL
});

let controller = new AbortController();

export const Connect = async (serverId: string, userId: string, timestamp: string, jwt: string) => {
  // While the connection is attempting to open, let the UI show a pending state
  status.pending();

  // If the client disconnected, the abort controller is no longer valid and a new one must be created
  if (controller.signal.aborted) {
    controller = new AbortController();
  }

  // The abort controller is used to signal the server to close the stream
  const opts = transport.mergeOptions({ abort: controller.signal });

  // Get the last timestamp from the cache
  const lastTs = get(chat_cache).lastTs

  // Create a new subscription to the server
  const sub = new ChatServiceClient(transport).connect({
    serverId: serverId,
    userId: userId,
    lastTs: timestamp,
    jwt: jwt,
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
};

export const SendMessage = (msg: OutgoingMessage) => {
  const client = new ChatServiceClient(transport);

  const request: ChatMessage = {
    channelId: msg.channelId,
    userId: msg.userId,
    text: msg.text,
    ts: "0", // The server will set the timestamp
    jwt: msg.jwt,
  };

  client.send(request).then((response) => {

    console.log(response.status.code);
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

  // Do not write messages with timestamp 0 to the UI
  if (msg.ts === "0") {
    return true;
  }

  if (msg.userId === "server") {
    return true;
  }

  if (msg.channelId === "system" && msg.userId !== pb.authStore.model?.name) {
    return filter_system_messages(msg);
  }

  // Deduplicate messages with the same timestamp
  if (msg.ts === lastTs) {
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
  if (msg.text.startsWith("channel_add")) {
    const channel_name = msg.text.split(" ")[1]
    console.log(channel_name)
    channels.add(channel_name);
  }

  return true;
}
