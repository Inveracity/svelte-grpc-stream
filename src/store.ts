import { writable } from 'svelte/store';

function createNotifier() {
	const { subscribe, set, update } = writable([""]);

	return {
		subscribe,
		write: (item: string | undefined) => update(notifs => {
			notifs.push(item ? item : ""); return notifs
		}),
		reset: () => set([""])
	};
}

function createStatus() {
	const { subscribe, set } = writable("");

	return {
		subscribe,
		connected: () => set("connected"),
		disconnected: () => set("disconnected"),
	};
}

export const notifier = createNotifier();
export const status = createStatus();
