<template>
  <div v-if="suggestData">
    <!-- 主体 -->
    <div v-for="(order, index) in suggestData.order" :key="index">
      <!-- 推荐搜索条目是什么类型 -->
      <div class="pt-2 pb-1.5 px-2.5">{{ getTitle(order) }}</div>

      <!-- 歌曲条目 -->
      <div v-if="order === 'songs'">
        <div
          v-for="item in suggestData.songs"
          :key="item.id"
          class="suggestItem hover-bg-main"
          @click="play(item.id)"
        >
          <!-- 歌曲名 -->
          <span class="text-emerald-500">{{ item.name }}</span>
          <span class="pl-1"> - {{ item.artist }}</span>
        </div>
      </div>

      <!-- 专辑条目 -->
      <div v-if="order === 'albums'">
        <div
          v-for="item in suggestData.albums"
          :key="item.id"
          class="suggestItem hover-bg-main"
          @click="jump('album', item.id)"
        >
          <span class="text-emerald-500">{{ item.name }}</span>
          <span class="pl-1"> - {{ item.name }}</span>
        </div>
      </div>

      <div v-if="order === 'artists'">
        <div
          v-for="item in suggestData.artists"
          :key="item.id"
          class="suggestItem hover-bg-main"
          @click="jump('artist', item.id)"
        >
          <span class="text-emerald-500">{{ item.name }}</span>
          <span class="text-emerald-500 ml-2">{{ item.name }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useSearchStore } from '@/stores/search'
import { usePlayerStore } from '@/stores/player'
import { routerPush } from '@/utils/navigator/router'

const { suggestData, showSearchView } = storeToRefs(useSearchStore())
const { play } = usePlayerStore()

//显示类型
const getTitle = (name: string) => {
  switch (name) {
    case 'songs':
      return '单曲'
    case 'albums':
      return '专辑'
    case 'artists':
      return '歌手'
    default:
      return name
  }
}

//跳转
const jump = (name: string, id: number) => {
  routerPush(name, id)
  showSearchView.value = false
}
</script>

<style lang="scss" scoped>
.suggestItem {
  @apply py-1.5 px-2.5 text-xs cursor-pointer;
}
</style>
