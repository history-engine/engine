<template>
  <v-app>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import http from "@/services/http"
import { useAppStore } from "@/stores/app";

export default {
  setup() {
    const store = useAppStore();
    return { store }
  },

  data: () => ({

  }),

  methods: {
  },

  mounted() {
    http({
      method: 'get',
      url: "/api/user/info",
    }).then(res => {
      if (res.code == 0) {
        this.store.login(res.data);
      }
    }).catch(err => {
      console.log('获取用户信息失败：' + err)
    });
  }
}
</script>
