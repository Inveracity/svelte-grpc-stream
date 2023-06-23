import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import type { SendRequest, SubscribeRequest } from './proto/notifications/v1/notifications';
import { NotificationServiceClient } from './proto/notifications/v1/notifications.client';
import { notifier } from './store';
import { status } from './store';

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
	transport.mergeOptions({abort: controller.signal});

	const client = new NotificationServiceClient(transport);

	// This request tells the server to open a stream to the client
	const request: SubscribeRequest = {
		channelId: channelId,
		userId: userId,
	};
	const call = client.subscribe(request);

	// While the connection is attempting to open, let the UI show a pending state
	status.pending();

	// If the connection fails, let the UI show an error state
	call.status.catch((e: Error) => {
		status.error(e.message);
	});

	// Listen for messages from the server
	for await (const msg of call.responses) {
		status.connected();
		console.log(msg)
		const message = `${msg.channelId}/${msg.userId}: ${msg.text}`;
		notifier.write(message);
	}

	await call.status;
	await call.trailers;

	status.disconnected();
	controller.abort();
};

// The client can actively unsubscribe letting the server know to close the stream
export const Unsubscribe = async () => {
	status.disconnected();
	controller.abort();
};

export const SendNotification = (channelId: string, userId: string, text: string) => {
	const client = new NotificationServiceClient(transport);

	const request: SendRequest  = {
			channelId: channelId,
			userId: userId,
			text: text,
	};

	client.send(request).then((response) => {

		console.log(response.status.code);
	}).catch((e) => {
		console.log(e);
	});
}
