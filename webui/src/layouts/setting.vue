<template>
  <v-app>
    <AppBar />

    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col cols="2">
            <v-sheet rounded="lg">
              <v-list rounded="lg" class="pa-2">
                <v-list-item
                  v-for="item in menuItems"
                  :key="item.title"
                  color="grey-lighten-4"
                  :title="item.title"
                  @click="navigateTo(item.path)"
                  :class="{ active: isActive(item.path) }"
                ></v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col>
            <v-sheet min-height="70vh" rounded="lg" class="pa-2">
              <router-view />
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <AppFooter />
  </v-app>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router';
import { onMounted, ref } from 'vue';

const route = useRoute();
const router = useRouter();

// 定义菜单项
const menuItems = ref([
  { title: '个人资料', path: '/setting/profile' },
  { title: '域名匹配', path: '/setting/host' },
  { title: '别名设置', path: '/setting/alias' },
  { title: '后缀过滤', path: '/setting/filetype' },
  { title: '存储设置', path: '/setting/storage' },
  { title: '页面管理', path: '/setting/page' }
]);

const isActive = (path) => {
  return route.path === path;
};

const navigateTo = (path) => {
  router.push(path);
};

// 如果当前路径是 '/setting'，则重定向到 '/setting/profile'
onMounted(() => {
  if (route.path === '/setting') {
    router.push('/setting/profile');
  }
});
</script>

<style scoped>
.active {
  background-color: #e0e0e0; /* 高亮背景色 */
  font-weight: bold;
}
</style>