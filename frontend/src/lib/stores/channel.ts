import { writeChannel } from '$lib/pocketbase';
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
    add: (channel: string) => update(channels => {
      // Precaution to limit channels to 10
      if (channels.length === 10) {
        return channels;
      }
      writeChannel(channel);
      // notify other users of the new channel
      return [...channels, channel]
    }),
    remove: (channel: string) => update(channels => channels.filter(c => c !== channel)),
    set: (channels: string[]) => update(_ => channels),
  };
}


export const channel = createChannelSelector();
export const channels = createChannelList();
