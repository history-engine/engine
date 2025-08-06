/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'
import { setupLayouts } from 'virtual:generated-layouts'

const routes = [
  {
    path: '/',
    component: () => import('@/pages/index.vue')
  },
  {
    path: '/user/login',
    component: () => import('@/pages/user/login.vue')
  },
  {
    path: '/user/register',
    component: () => import('@/pages/user/register.vue')
  },
  {
    path: '/about',
    component: () => import('@/pages/about.vue')
  },
  //个人设置相关
  {
    path: '/setting',
    component: () => import('@/pages/setting/index.vue')
  },
  {
    path: '/setting/profile',
    component: () => import('@/pages/setting/profile.vue')
  },
  {
    path: '/setting/host',
    component: () => import('@/pages/setting/host.vue')
  },
  {
    path: '/setting/alias',
    component: () => import('@/pages/setting/alias.vue')
  },
  {
    path: '/setting/storage',
    component: () => import('@/pages/setting/storage.vue')
  },
  {
    path: '/setting/filetype',
    component: () => import('@/pages/setting/filetype.vue')
  },
  {
    path: '/setting/page',
    component: () => import('@/pages/setting/page.vue')
  },
  // 管理员相关
  {
    path: '/admin',
    name: 'AdminIndex',
    component: () => import('@/pages/admin/index.vue')
  },
  {
    path: '/admin/user',
    name: 'AdminUser',
    component: () => import('@/pages/admin/user.vue')
  },
  {
    path: '/admin/icon',
    name: 'AdminIcon',
    component: () => import('@/pages/admin/icon.vue')
  },
  {
    path: '/admin/misc',
    name: 'AdminMisc',
    component: () => import('@/pages/admin/misc.vue')
  },
  {
    path: '/admin/file',
    name: 'AdminFile',
    component: () => import('@/pages/admin/file.vue')
  },
  {
    path: '/stats',
    component: () => import('@/pages/stats.vue')
  }
]
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  extendRoutes: setupLayouts,
  routes: routes
})

export default router
