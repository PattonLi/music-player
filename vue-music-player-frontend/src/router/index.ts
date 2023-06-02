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
      component: () => import('@/views/appRoot/AppRoot.vue'),
      redirect: { name: Pages.discover },
      children: [
        {
          //首页，进入时的默认界面
          path: 'discover',
          name: 'discover',
          meta: {
            menu: 'discover'
          },
          component: () => import('@/views/discover/BaseDiscover.vue')
        },
        {
          //电台组件
          path: 'radio',
          name: 'radio',
          meta: {
            menu: 'radio'
          },
          component: () => import('@/views/radio/RadioStation.vue')
        },
        {
          //音乐视频组件
          path: 'video',
          name: 'video',
          meta: {
            menu: 'video'
          },
          component: () => import('@/views/mv/MusicVideo.vue')
        },
        {
          path: 'artistDetail',
          name: 'artistDetail',
          component: () => import('@/views/artist/AppArtistDetail.vue')
        },
        {
          path: 'albumDetail',
          name: 'albumDetail',
          component: () => import('@/views/album/AppAlbumDetail.vue')
        },
        {
          //音乐库组件
          path: 'library',
          name: 'library',
          meta: {
            menu: 'library'
          },
          component: () => import('@/views/library/AppMusicLibrary.vue')
        }
      ]
    }
  ]
})

export default router
