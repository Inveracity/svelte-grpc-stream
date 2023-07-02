import { writable } from 'svelte/store';
import type { IMessage } from './types';

function createMessages() {
  const { subscribe, set, update } = writable<Array<IMessage>>([]);

  return {
    subscribe,
    add: (msg: IMessage) => update(n => [...n, msg]),
    reset: () => set([]),
  };
}

export const messages = createMessages();

function createChannelSelector() {
  const { subscribe, update } = writable<string>('general');

  return {
    subscribe,
    set: (channel: string) => update(_ => channel),
  };
}

export const channel = createChannelSelector();
