import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import type { SubscribeRequest } from './proto/notifications/v1/notifications';
import { NotificationServiceClient } from "./proto/notifications/v1/notifications.client";
import { notifier } from './store';
import { status } from './store';

let controller = new AbortController()

export const Subscribe = async (subscriberId: string) => {
  // If the client disconnected, the abort controller is no longer valid and a new one must be created
  if (controller.signal.aborted) {
    controller = new AbortController()
  }

  // Create a new transport and client
  const transport = new GrpcWebFetchTransport({
    baseUrl: "http://relay.docker.localhost",
    abort: controller.signal,
  });
  
  const client = new NotificationServiceClient(transport)

  // This request tells the server to open a stream to the client
  const request: SubscribeRequest = {subid: subscriberId}
  const call = client.notify(request)

  // While the connection is attempting to open, let the UI show a pending state
  status.pending()
  
  // If the connection fails, let the UI show an error state
  call.status.catch((e) => {
    status.error(e.message)
  })
  
  // Listen for messages from the server
  for await (const msg of call.responses ) {
    status.connected()
    const message = `${msg.notifications?.sender}: ${msg.notifications?.text}`
    notifier.write(message)
  }
  
  await call.status
  await call.trailers

  status.disconnected()
  controller.abort() 
}

// The client can actively unsubscribe letting the server know to close the stream
export const Unsubscribe = async () => {
  status.disconnected()
  controller.abort()
}