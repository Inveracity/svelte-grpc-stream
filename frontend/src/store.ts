import { writable } from 'svelte/store';
import type { Message } from './types';

const createNotifier = () => {
	const { subscribe, set, update } = writable<Message[]>([]);

	return {
		subscribe,
		write: (item: Message | undefined) =>
			update((notifs) => {
				if (item) {
					notifs.push(item);
				}
				return notifs;
			}),
		reset: () => set([])
	};
};

export const notifier = createNotifier();

function createStatus() {
	const { subscribe, set } = writable('');

	return {
		subscribe,
		connected: () => set('connected'),
		disconnected: () => set('disconnected'),
		pending: () => set('pending'),
		error: (errormsg: string) => set('error: ' + errormsg)
	};
}

export const status = createStatus();
