import { writable } from 'svelte/store';

const createNotifier = () => {
	const { subscribe, set, update } = writable(['']);

	return {
		subscribe,
		write: (item: string | undefined) =>
			update((notifs) => {
				notifs.push(item ? item : '');
				return notifs;
			}),
		reset: () => set([''])
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
