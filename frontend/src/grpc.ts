import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { NotificationServiceClient } from './proto/notifications/v1/notifications.client';
import { notifier } from './store';
import { status } from './store';
import { persisted } from 'svelte-local-storage-store'
import { get } from 'svelte/store'

export const notifications_cache = persisted(
	'notifications', // storage
	{lastTs: '0'}, // default value
	{storage: 'session'} // options
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
	const opts = transport.mergeOptions({abort: controller.signal});

	// Get the last timestamp from the cache
	let lastTs = get(notifications_cache).lastTs

	// Create a new subscription to the server
	const sub = new NotificationServiceClient(transport).subscribe({channelId, userId, lastTs}, opts);

	// While the connection is attempting to open, let the UI show a pending state
	status.pending();

	// If the connection fails, let the UI show an error state
	sub.status.catch((e: Error) => {
		status.error(e.message);
	});

	// Listen for messages from the server
	try {
		for await (const msg of sub.responses) {
			status.connected();
			const message = `${msg.channelId}/${msg.userId}: ${msg.text}`;
			notifier.write(message);
			notifications_cache.set({lastTs: msg.ts})
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

	const request = {
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
