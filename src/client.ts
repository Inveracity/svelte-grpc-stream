import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import type { NotifyRequest } from './gen/notifications/v1/notifications';
import { NotificationServiceClient } from "./gen/notifications/v1/notifications.client";
import { notifier } from './store';

const transport = new GrpcWebFetchTransport({
  baseUrl: "notifier.docker.localhost",
});

const client = new NotificationServiceClient(transport)

export const Subscribe = async () => {
    const request: NotifyRequest = {}
    const call = client.notify(request)
    for await (const msg of call.responses ) {
        notifier.write(msg.notifications?.name)
        console.log("got message", msg)
    }
    const status = await call.status
    const trailers = await call.trailers
}
