import { createRouter, createWebHistory } from 'vue-router'
//import struct page
import { Pages } from '@/router/Pages'
//import components

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: Pages.home,
      component: () => import('@/App.vue'),
      redirect: { name: Pages.discover },
      children: [
        {
          //首页，进入时的默认界面
          path: 'discover',
          name: 'discover',
          component: () => import('@/views/discover/BaseDiscover.vue')
        }
      ]
    }
  ]
})

export default router
