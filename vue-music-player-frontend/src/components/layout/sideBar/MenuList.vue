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
import { useMenu } from '@/components/layout/sideBar/useMenu'
//图标组件
import IconPark from '@/components/common/IconPark.vue'

const { menus, currentKey, click } = useMenu()
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
