<template>
  <v-app-bar flat class="bg-grey-lighten-4">
    <v-container class="mx-auto d-flex align-center justify-center ">
      <v-btn text="首页" variant="text" @click="goHome"></v-btn>
      <v-btn text="统计信息" variant="text" @click="$router.push('/stats')"></v-btn>
      <v-btn text="设置中心" variant="text" @click="$router.push('/setting')"></v-btn>
      <v-btn text="管理中心" variant="text" @click="$router.push('/admin')"></v-btn>

      <v-spacer></v-spacer>

      <AppAvatar />
    </v-container>
  </v-app-bar>
</template>

<script setup>
import AppAvatar from "@/components/AppAvatar.vue";
</script>

<script>
import { useAppStore } from "@/stores/app";
export default {
  setup() {
    const store = useAppStore();
    return { store, logo }
  },

  data: () => ({
    //
  }),

  methods: {
    goHome() {
      window.location.href = "/";
    }
  }
}
</script>
