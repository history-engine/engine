// Utilities
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    keyword: "",
    online: false,
    user: {},
  }),

  actions: {
    login(user) {
      if (user && user.id > 0) {
        this.online = true;
        this.user = user;
      }
    },
    logout() {
      this.online = false;
      this.user = {};
    },
  },
})
