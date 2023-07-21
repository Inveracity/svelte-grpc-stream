import { env } from '$env/dynamic/public';
import PocketBase from 'pocketbase';
import { writable } from 'svelte/store';
import { channels } from './stores/channel';
import { users, type User } from './stores/users';

export const pb = new PocketBase(env.PUBLIC_POCKETBASE_URL);
export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange(() => {
  currentUser.set(pb.authStore.model);
});

export const fetchChannels = async () => {
  const records = await pb.collection('channels').getFullList({
    sort: 'created',
  });

  // convert records to array and set in channels store
  channels.set(records.map((record) => {
    return record.name;
  }));
}

export const writeChannel = async (channelName: string) => {
  await pb.collection('channels').create({
    name: channelName,
  });
}

export const fetchUsers = async () => {
  const records: User[] = await pb.collection('users').getFullList({
    sort: 'created',
  });

  // convert records to array and set in channels store
  users.set(records.map((record: User) => {
    const user: User = { name: record.name, presence: false };
    return user;
  }));
}
