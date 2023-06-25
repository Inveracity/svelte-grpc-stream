import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { NotificationServiceClient } from './proto/notifications/v1/notifications.client';
import type { Notification } from './proto/notifications/v1/notifications';
import { notifier } from './store';
import { status } from './store';
import { persisted } from 'svelte-local-storage-store'
import { get } from 'svelte/store'
import { DateTime } from 'luxon';
import type { Message } from './types';

export const notifications_cache = persisted(
  'notifications', // storage
  { lastTs: '0' }, // default value
  { storage: 'session' } // options
)

const transport = new GrpcWebFetchTransport({
  baseUrl: 'http://relay.docker.localhost'
});

let controller = new AbortController();

export const Subscribe = async (channelId: string, userId: string, timestamp: string) => {
  // If the client disconnected, the abort controller is no longer valid and a new one must be created
  if (controller.signal.aborted) {
    controller = new AbortController();
  }

  // The abort controller is used to signal the server to close the stream
  const opts = transport.mergeOptions({ abort: controller.signal });

  // Get the last timestamp from the cache
  let lastTs = get(notifications_cache).lastTs

  // Create a new subscription to the server
  const sub = new NotificationServiceClient(transport).subscribe({ channelId, userId, lastTs }, opts);

  // While the connection is attempting to open, let the UI show a pending state
  status.pending();

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

      notifier.write(message);
      notifications_cache.set({ lastTs: msg.ts })

    }
  } catch (e: any) {
    console.log("Stream closed");
  }

  status.disconnected();
};

// The client can actively unsubscribe letting the server know to close the stream
export const Unsubscribe = async () => {
  controller.abort();
};

export const SendNotification = (channelId: string, userId: string, text: string) => {
  const client = new NotificationServiceClient(transport);

  const request: Notification = {
    channelId: channelId,
    userId: userId,
    text: text,
    ts: "0", // The server will set the timestamp
  };

  client.send(request).then((response) => {

    console.log(response.status.code);
  }).catch((e) => {
    console.log(e);
  });
}

const filtered = (msg: Notification, lastTs: string): boolean => {
  // Do not write server messages to the UI
  if (msg.userId === "server" && msg.text === "connected") {
    status.connected();
    return true;
  }

  // Do not write messages with timestamp 0 to the UI
  if (msg.ts === "0") {
    return true;
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
    return DateTime.fromMillis(nano / 1000000).toFormat("HH:mm:ss")
  } catch (e) {
    console.log(e);
    return timestamp;
  }
}
