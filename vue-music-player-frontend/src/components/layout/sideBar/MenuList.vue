<template>
  <!-- 图标 -->
  <div class="mx-14 pt-4">
    <img style="width: 80px; height: 80px" src="@/assets/images/music-logo.svg" />
  </div>
  <!--菜单栏拖动条-->
  <ElScrollbar>
    <!-- 边距样式 -->
    <div class="mt-8 pl-8 pr-6 space-y-2" v-for="(menuItem, key) in menus" :key="key">
      <!-- 菜单栏 -->
      <div class="menu-title text-main">{{ menuItem.name }}</div>
      <!-- 菜单 -->
      <div
        class="menu-item hover-bg-main text-main"
        v-for="menu in menuItem.menus"
        :key="menu.key"
        :class="{ active: currentKey === menu.key }"
        @click="click(menu)"
      >
        <!-- 图标 -->
        <IconPark :icon="menu.icon" size="18" :theme="menu.theme" />
        <!-- 菜单名称 -->
        <span class="ml-2">{{ menu.name }}</span>
      </div>
    </div>
  </ElScrollbar>
</template>

<script setup lang="ts">
//图标组件
import IconPark from '@/components/common/IconPark.vue'
import { Pages } from '@/router/pages'
//引入菜单栏图标
import {
  Planet,
  Music,
  VideoOne,
  Fm,
  Like,
  Computer,
  DownloadThree,
  Time
} from '@icon-park/vue-next'

//menu接口类型
interface IMenu {
  name: string
  key: string
  icon: any
  theme: 'outline' | 'filled' | 'two-tone' | 'multi-color'
}
//菜单组类型
interface IMenus {
  name: string
  menus: IMenu[]
}

//初始化菜单内容
const menus: IMenus[] = [
  //菜单栏1
  {
    name: '在线音乐',
    menus: [
      {
        name: '推荐',
        key: Pages.discover,
        icon: Planet,
        theme: 'outline'
      },
      {
        name: '音乐馆',
        key: Pages.library,
        icon: Music,
        theme: 'outline'
      },
      {
        name: '视频',
        key: Pages.mv,
        icon: VideoOne,
        theme: 'outline'
      },
      {
        name: '雷达',
        key: Pages.radio,
        icon: Fm,
        theme: 'outline'
      }
    ]
  },
  //菜单栏2
  {
    name: '我的音乐',
    menus: [
      {
        name: '我喜欢',
        key: Pages.like,
        icon: Like,
        theme: 'outline'
      },
      {
        name: '本地歌曲',
        key: Pages.local,
        icon: Computer,
        theme: 'outline'
      },
      {
        name: '下载歌曲',
        key: Pages.download,
        icon: DownloadThree,
        theme: 'outline'
      },
      {
        name: '最近播放',
        key: Pages.recentPlay,
        icon: Time,
        theme: 'outline'
      }
    ]
  }
]

const route = useRoute()
const currentKey = ref(route.meta.menu)
const router = useRouter()

//监视路由元数据
watch(
  () => route.meta.menu,
  (menu) => {
    currentKey.value = menu
  }
)

const click = async (menu: IMenu) => {
  await router.push({ name: menu.key, replace: true })
}
</script>

<style lang="scss" scoped>
.menu-title {
  @apply text-sm pl-5 pr-4 pb-2 text-gray-500;
}

.menu-item {
  @apply text-sm pl-5 pr-3 pt-2 pb-2 rounded cursor-pointer transition-colors flex items-center;
}

// 当选中菜单时 TODO
.active {
  @apply bg-gradient-to-r from-teal-500 to-emerald-300 text-slate-50 cursor-default;
}
</style>
