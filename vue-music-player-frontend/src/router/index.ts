import { createRouter, createWebHistory } from 'vue-router'
//import struct page
import { Pages } from '@/router/pages'
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
        /*--------------------菜单组件 ------------------*/
        {
          //首页，进入时的默认界面
          path: 'discover',
          name: Pages.discover,
          meta: {
            menu: 'discover'
          },
          component: () => import('@/views/discover/AppDiscover.vue')
        },
        {
          //电台组件
          path: 'radio',
          name: Pages.radio,
          meta: {
            menu: 'radio'
          },
          component: () => import('@/views/radio/AppRadio.vue')
        },
        {
          //音乐视频组件
          path: 'video',
          name: Pages.mv,
          meta: {
            menu: 'mv'
          },
          component: () => import('@/views/mv/AppMusicVideo.vue')
        },
        {
          //音乐库组件
          path: 'library',
          name: Pages.library,
          meta: {
            menu: 'library'
          },
          component: () => import('@/views/library/AppMusicLibrary.vue'),
          children: [
            /*--------------------音乐库子组件 ------------------*/
            {
              path: 'select',
              name: Pages.select,
              component: () => import('@/views/library/select/AppSelect.vue')
            },
            {
              //电台组件
              path: 'radio',
              name: Pages.radio,
              component: () => import('@/views/radio/AppRadio.vue')
            },
            {
              //排行榜组件
              path: 'top',
              name: Pages.top,
              component: () => import('@/views/library/toplist/AppTopList.vue')
            },
            {
              //歌手组件
              path: 'artist',
              name: Pages.artist,
              component: () => import('@/views/library/artist/AppArtist.vue')
            },
            {
              //分类歌单组件
              path: 'playlistCata',
              name: Pages.playlistCata,
              component: () => import('@/views/library/playlist-catagory/AppPlayListCata.vue')
            },
            {
              //数字专辑组件组件
              path: 'digitalAlbum',
              name: Pages.digitalAlbum,
              component: () => import('@/views/library/digital-album/AppDigitalAlbum.vue')
            },
            {
              //手机专享组件
              path: 'phoneOnly',
              name: Pages.phoneOnly,
              component: () => import('@/views/library/phone-only/AppPhoneOnly.vue')
            }
          ]
        },

        /*--------------------详情页组件 ------------------*/
        {
          //音乐库组件
          path: 'detail',
          name: 'detail',
          component: () => import('@/views/detail/AppInfo.vue'),
          children: [
            {
              //歌手详情页
              path: 'artistDetail',
              name: Pages.artistDetail,
              component: () => import('@/views/detail/artist/AppArtistDetail.vue')
            },
            {
              //专辑详情页
              path: 'albumDetail',
              name: Pages.albumDetail,
              component: () => import('@/views/detail/album/AppAlbumDetail.vue')
            }
          ]
        },

        /*--------------------搜索结果组件 ------------------*/

        {
          //搜索结果页
          path: 'search/result',
          name: Pages.searchResult,
          component: () => import('@/views/search-result/AppSearchResult.vue')
        },

        /*--------------------播放器组件 ------------------*/
        {
          //mv player组件
          path: 'mvPlayer',
          name: Pages.mvPlayer,
          component: () => import('@/views/mv-player/AppMvPlayer.vue')
        },
        /*--------------------用户中心组件 ------------------*/
        {
          //mv player组件
          path: 'userCenter',
          name: Pages.userCenter,
          component: () => import('@/views/user-center/AppUserCenter.vue')
        },
      ]
    }
  ]
})

export default router
