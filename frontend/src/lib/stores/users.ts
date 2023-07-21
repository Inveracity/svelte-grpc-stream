import { writable } from 'svelte/store';

export interface User {
  name: string;
  presence: boolean;
}

function createUserList() {
  const { subscribe, update } = writable<User[]>([]);
  return {
    subscribe,
    add: (user: User) => update(users => [...users, user]),
    remove: (user: User) => update(users => users.filter(c => c !== user)),
    set: (users: User[]) => update(_ => users),
    upd: (user: User) => update(users => {
      const index = users.findIndex(u => u.name === user.name);
      if (index === -1) {
        console.log("user not found", user.name)
        return users;
      }
      users[index] = user;
      return users;
    }),
  };
}

function toggleUserList() {
  const { subscribe, update } = writable<boolean>(false);
  return {
    subscribe,
    toggle: () => update(v => !v),
  };
}

export const users = createUserList();
export const showUserList = toggleUserList();
