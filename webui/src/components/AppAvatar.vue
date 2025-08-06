<template>
  <v-menu min-width="200px" v-if="store.online">
    <template v-slot:activator="{ props }">
      <v-btn icon v-bind="props">
        <v-avatar color="brown" size="default" :image="store.user.avatar"></v-avatar>
      </v-btn>
    </template>

    <v-card>
      <v-card-text>
        <div class="mx-auto ">
          <h3>{{ store.user.username }}</h3>
          <p class="text-caption mt-1">{{ store.user.email }}</p>
          <v-divider class="my-3"></v-divider>
          <v-btn @click="$router.push('/user/logout')" variant="text" rounded>退出登录</v-btn>
        </div>
      </v-card-text>
    </v-card>
  </v-menu>

   <div v-if="!store.online">
     <v-btn text="注册" variant="text" @click="$router.push('/user/register')"></v-btn>
     <v-btn text="登录" variant="text"  @click="$router.push('/user/login')"></v-btn>
   </div>
</template>

<script>
import { useAppStore } from "@/stores/app";
export default {
  setup() {
    const store = useAppStore();
    return { store }
  },

  data: () => ({
    //
  }),
}
</script>
