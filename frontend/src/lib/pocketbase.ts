import PocketBase from 'pocketbase';

import { writable } from 'svelte/store';

export const pb = new PocketBase("http://frontend.docker.localhost/pocketbase");

export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange(() => {
  currentUser.set(pb.authStore.model);
});
