<template>
  <v-container>
      <v-sheet class="d-flex flex-row-reverse">
        <v-sheet v-show="store.online" class="ma-2 pa-2" @click="$router.push('/user/logout')">
          <v-btn>退出</v-btn>
        </v-sheet>
        <v-sheet v-show="!store.online" class="ma-2 pa-2" @click="$router.push('/user/login')">
          <v-btn>登录</v-btn>
        </v-sheet>
        <v-sheet v-show="!store.online" class="ma-2 pa-2 " @click="$router.push('/user/register')">
          <v-btn>注册</v-btn>
        </v-sheet>
        <v-sheet class="ma-2 pa-2" @click="$router.push('/')">
          <v-btn>首页</v-btn>
        </v-sheet>
      </v-sheet>
  </v-container>
</template>

<script>
import logo from '@/assets/logo.svg?raw'
import { useAppStore } from "@/stores/app";

export default {
  setup() {
    const store = useAppStore();
    return { store,logo}
  },

  data: () => ({
  }),

  mounted() {
  },

  methods: {
  },
};
</script>
