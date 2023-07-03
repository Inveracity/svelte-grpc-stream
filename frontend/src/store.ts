import { writable } from 'svelte/store';
import type { Message } from './types';

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

function createMessages() {
  const { subscribe, set, update } = writable<Array<Message>>([]);

  return {
    subscribe,
    add: (msg: Message) => update(n => [...n, msg]),
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

function createUsername() {
  const { subscribe, update } = writable<string>('');

  return {
    subscribe,
    set: (channel: string) => update(_ => channel),
  };
}

export const username = createUsername();