<template>
  <div class="flex items-center justify-between drag h-14">
    <!-- 导航栏相关 -->
    <div class="flex items-center pl-5">
      <!-- 左导航栏 -->
      <IconPark
        :icon="Left"
        :size="iconSize"
        :stroke-width="2"
        class="icon-button"
        @click="router.back()"
      />
      <!-- 右导航栏 -->
      <IconPark :icon="Right" :size="iconSize" :stroke-width="2" class="icon-button" />
      <!-- 搜索框 -->
      <div class="search no-drag ml-2">
        <SearchPop />
      </div>
    </div>

    <!-- 用户信息显示 -->
    <div class="flex items-center mr-5">
      <UserInfo />
      <IconPark :icon="Mail" :stroke-width="2" class="icon-button" />
      <IconPark :icon="Platte" :stroke-width="2" class="icon-button" @click="changeThemeMode" />
      <IconPark :icon="HamburgerButton" :stroke-width="2" class="icon-button" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { HamburgerButton, Left, Mail, Platte, Right, Search } from '@icon-park/vue-next'
import { useRouter } from 'vue-router'
import IconPark from '@/components/common/IconPark.vue'
import UserInfo from '@/components/layout/header/UserInfo.vue'
import SearchPop from '@/components/layout/header/SearchPop.vue'

const iconSize = 22

const router = useRouter()

const mode = ref(false)

const changeThemeMode = () => {
  window.document.documentElement.setAttribute('data-theme', mode.value ? 'dark' : 'light')

  mode.value = !mode.value
}
</script>
<style lang="scss" scoped>
.search {
  @apply text-xs;
  :deep(.el-input__inner) {
    @apply rounded-full text-xs bg-slate-100 dark:bg-stone-900;
  }
}

.icon-button {
  @apply p-2 cursor-pointer;
  @apply hover:text-emerald-400;
}
</style>
