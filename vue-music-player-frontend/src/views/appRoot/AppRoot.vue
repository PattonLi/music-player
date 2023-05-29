<template>
  <!-- 顶层容器 -->
  <!-- 全屏幕、flex布局、铺展元素、隐藏溢出 -->
  <div class="w-screen h-screen flex items-stretch overflow-hidden">
    <!-- 菜单栏 -->
    <!-- 固定大小、不被压缩 -->
    <div class="w-56 h-screen flex-shrink-0">
      <AppMenu />
    </div>
    <!-- 主要内容部分 -->
    <!-- 纵向排布 -->
    <div class="flex-1 flex flex-col">
      <!-- 导航栏 -->
      <div class="h-16">
        <AppHeader />
      </div>
      <!-- 内容部分 -->
      <div class="flex-1 overflow-hidden">
        <ElScrollbar>
          <div class="container mx-auto">
            <RouterView />
          </div>
        </ElScrollbar>
      </div>
      <!--播放控制区域-->
      <div class="h-20">
        <AppFooter />
      </div>
    </div>
    <!-- 播放列表popover -->
    <AppPlayList />
  </div>
</template>

<script setup lang="ts">
import AppMenu from '@/components/layout/sideBar/AppMenu.vue'
import AppHeader from '@/components/layout/header/AppHeader.vue'
import AppFooter from '@/components/layout/footer/AppFooter.vue'
import AppPlayList from '@/components/layout/playList/AppPlayList.vue'
</script>

<style lang="scss"></style>
