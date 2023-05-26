<template>
  <!--菜单栏拖动条-->
  <ElScrollbar>
    <div class="mt-10 pl-6 pr-6 space-y-1.5" v-for="(menuItem, key) in menus" :key="key">
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
        <span class="ml-1">{{ menu.name }}</span>
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
  @apply text-xs pl-4 pr-4 pb-2;
}

.menu-item {
  @apply text-sm pl-4 pr-4 pt-1.5 pb-1.5 rounded cursor-pointer transition-colors flex items-center;
}

.active {
  @apply bg-gradient-to-r from-teal-400 to-emerald-400 text-slate-50 cursor-default;
}
</style>
