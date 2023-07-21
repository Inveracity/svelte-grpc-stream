import { writable } from 'svelte/store';

function createUserList() {
  const { subscribe, update } = writable<string[]>([]);
  return {
    subscribe,
    add: (user: string) => update(users => [...users, user]),
    remove: (user: string) => update(users => users.filter(c => c !== user)),
    set: (users: string[]) => update(_ => users),
  };
}

export const users = createUserList();
