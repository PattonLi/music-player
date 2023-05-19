import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue')
    },
    {
      path: '/login',
      component: () => import('@/views/login/MyLogin.vue'),
      name: 'login'
    },
    {
      path: '/addsong',
      component: () => import('@/views/song-mana/AddSong.vue'),
      name: 'addsong'
    }
  ]
})

export default router
