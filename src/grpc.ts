import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import type { NotifyRequest } from './proto/notifications/v1/notifications';
import { NotificationServiceClient } from "./proto/notifications/v1/notifications.client";
import { notifier } from './store';
import { status } from './store';


// Create a new AbortController for each subscription
let controller = new AbortController()
// Assume the default state is disconnected
let state = "disconnected"
// Subscribe to the status store to update whether the client is connected or not
status.subscribe((value) => {state = value})

export const Subscribe = async (subscriberId: string) => {
  // If the client disconnected, the abort controller is no longer valid and a new one must be created
  if (controller.signal.aborted) {
    controller = new AbortController()
  }

  const transport = new GrpcWebFetchTransport({
    baseUrl: "http://notifier.docker.localhost",
    abort: controller.signal,
  });
  
  const client = new NotificationServiceClient(transport)
  const request: NotifyRequest = {id: subscriberId}
  const call = client.notify(request)

  // Update UI to show that the client is connected
  status.connected()

  // Listen for messages from the server
  for await (const msg of call.responses ) {
    console.log("connected!")
    if (state === "disconnected") { 
      console.log("breaking")
      break 
    }
    
    notifier.write(msg.notifications?.msg)
  }
  
  let callstatus = await call.status;
  console.log(callstatus)
  status.disconnected()
  controller.abort()
}

export const Unsubscribe = async () => {
  status.disconnected()
  controller.abort()
}