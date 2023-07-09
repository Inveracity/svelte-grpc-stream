import { writable } from 'svelte/store';

function createChannelSelector() {
  const { subscribe, update } = writable<string>('general');

  return {
    subscribe,
    set: (channel: string) => update(_ => channel),
  };
}

function createChannelList() {
  const { subscribe, update } = writable<string[]>([]);
  return {
    subscribe,
    add: (channel: string) => update(channels => [...channels, channel]),
    remove: (channel: string) => update(channels => channels.filter(c => c !== channel)),
    set: (channels: string[]) => update(_ => channels),
  };
}


export const channel = createChannelSelector();
export const channels = createChannelList();
