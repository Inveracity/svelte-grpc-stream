import { writable } from 'svelte/store';

function createChannelSelector() {
  const { subscribe, update } = writable<string>('general');

  return {
    subscribe,
    set: (channel: string) => update(_ => channel),
  };
}

export const channel = createChannelSelector();
