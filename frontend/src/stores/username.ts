import { writable } from 'svelte/store';

function createUsername() {
  const { subscribe, update } = writable<string>('');

  return {
    subscribe,
    set: (channel: string) => update(_ => channel),
  };
}

export const username = createUsername();
