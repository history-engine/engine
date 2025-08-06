<template>
  <v-container class="">
    <h1>logout</h1>
  </v-container>
</template>

<style scoped lang="sass">
</style>

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
      url: "/api/user/logout",
    }).then(res => {
      if (res.code == 0) {
        this.store.logout()
        this.$router.push('/user/login')
      }
    }).catch(err => {
      console.log('退出失败：' + err)
    });
  }
}
</script>
