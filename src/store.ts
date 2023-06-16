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

export const notifier = createNotifier();
