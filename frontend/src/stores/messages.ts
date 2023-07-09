import { writable } from 'svelte/store';
import type { Message } from '../types';

const democontent: Array<Message> = []

function createMessages() {
  const { subscribe, set, update } = writable<Array<Message>>(democontent);

  return {
    subscribe,
    add: (msg: Message) => update(n => [...n, msg]),
    reset: () => set([]),
  };
}

export const messages = createMessages();
