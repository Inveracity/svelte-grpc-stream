import { writable } from 'svelte/store';

function createChannelSelector() {
  const { subscribe, update } = writable<string>('general');

  return {
    subscribe,
    set: (channel: string) => update(() => channel)
  };
}

function createChannelList() {
  const { subscribe, update } = writable<string[]>([]);
  return {
    subscribe,
    add: (channel: string) => update((channels) => [...channels, channel]),
    remove: (channel: string) => update((channels) => channels.filter((c) => c !== channel)),
    set: (channels: string[]) => update(() => channels)
  };
}

function toggleChannelList() {
  const { subscribe, update } = writable<boolean>(true);
  return {
    subscribe,
    toggle: () => update((v) => !v)
  };
}

export const channel = createChannelSelector();
export const channels = createChannelList();
export const showChannelList = toggleChannelList();
