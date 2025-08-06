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
  { title: '用户管理', path: '/admin/user' },
  { title: 'Icon管理', path: '/admin/icon' },
  { title: '杂项设置', path: '/admin/misc' },
  { title: '文件管理', path: '/admin/file' }
]);

const isActive = (path) => {
  return route.path === path;
};

const navigateTo = (path) => {
  router.push(path);
};

// 如果当前路径是 '/admin'，则重定向到 '/admin/user'
onMounted(() => {
  if (route.path === '/admin') {
    router.push('/admin/user');
  }
});
</script>

<style scoped>
.active {
  background-color: #e0e0e0; /* 高亮背景色 */
  font-weight: bold;
}
</style>