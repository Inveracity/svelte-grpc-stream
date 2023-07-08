import { writable } from 'svelte/store';
import type { Message } from '../types';

const democontent: Array<Message> = [
  {
    message: 'Hello World',
    user: 'user1',
    channel: 'general',
    timestamp: "2023-01-01 12:00:00",
  },
  {
    message: 'Hello World',
    user: 'chris',
    channel: 'general',
    timestamp: "2023-01-01 12:00:00",
  },
]

function createMessages() {
  const { subscribe, set, update } = writable<Array<Message>>(democontent);

  return {
    subscribe,
    add: (msg: Message) => update(n => [...n, msg]),
    reset: () => set([]),
  };
}

export const messages = createMessages();