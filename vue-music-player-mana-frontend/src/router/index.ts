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
      path: '/song',
      component: () => import('@/views/song-mana/AddSong.vue'),
      name: 'song',
      children: [
        {
          path: '/add',
          component: () => import('@/views/song-mana/AddSong.vue'),
          name: 'add'
        },
        {
          path: '/deleteSong',
          component: () => import('@/views/song-mana/AddSong.vue'),
          name: 'deleteSong'
        },
        {
          path: '/modiSong',
          component: () => import('@/views/song-mana/AddSong.vue'),
          name: 'modiSong'
        }
      ]
    }

    //用户相关

    //日志相关
  ]
})

export default router
