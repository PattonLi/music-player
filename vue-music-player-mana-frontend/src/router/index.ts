import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    //大盘数据
    {
      path: '/',
      name: 'welcome',
      component: () => import('@/views/data-analyse/DataAnalyse.vue')
    },
    //登录
    {
      path: '/login',
      component: () => import('@/views/login/MyLogin.vue'),
      name: 'login'
    },
    //歌曲相关
    {
      path: '/addSong',
      component: () => import('@/views/song-mana/AddSong.vue'),
      name: 'addSong'
    }

    //用户相关

    //日志相关
  ]
})

export default router
