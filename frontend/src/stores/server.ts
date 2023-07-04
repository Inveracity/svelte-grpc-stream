import { writable } from 'svelte/store';

function createServer() {
    const { subscribe, update } = writable<string>('myserver');

    return {
      subscribe,
      set: (server: string) => update(_ => server),
    };
  }

  export const server = createServer();
