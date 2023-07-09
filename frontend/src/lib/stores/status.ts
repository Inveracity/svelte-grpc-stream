import { writable } from 'svelte/store';

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
