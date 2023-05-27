<template>
  <!-- 图标 -->
  <div class="mx-12 pt-4">
    <IconPark :icon=Performance theme="outline" size="75" />
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
import { Performance } from "@icon-park/vue-next";

const { menus, currentKey, click } = useMenu()
</script>

<style lang="scss" scoped>
.menu-title {
  @apply text-sm pl-4 pr-4 pb-2 text-gray-500;
}

.menu-item {
  @apply text-sm pl-4 pr-4 pt-1.5 pb-1.5 rounded cursor-pointer transition-colors flex items-center;
}

// 当选中菜单时 TODO
.active {
  @apply text-green-400 cursor-default;
}
</style>
