import { env } from '$env/dynamic/public';

import PocketBase from 'pocketbase';

import { writable } from 'svelte/store';

export const pb = new PocketBase(env.PUBLIC_POCKETBASE_URL);

export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange(() => {
  currentUser.set(pb.authStore.model);
});
