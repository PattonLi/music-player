<template>
  <div v-if="suggestData">
    <!-- 主体 -->
    <div v-for="(order, index) in suggestData.order" :key="index">
      <!-- 推荐搜索条目是什么类型 -->
      <div class="pt-2 pb-1.5 px-2.5">{{ getTitle(order) }}</div>
      <div v-if="order === 'songs'">
        <div
        v-for="item in suggestData.songs"
        :key="item.id"
        class="py-1.5 px-2.5 hover-bg-main text-xs cursor-pointer"
        @click="play(item.id)"
      >
        <span class="text-emerald-500">{{ item.name }}</span>
        <span class="pl-1"> - {{ item.artists.first()?.name }}</span>
        </div>
      </div>
      

      <div
        v-for="item in suggestData.albums"
        :key="item.id"
        v-if="order === 'albums'"
        class="py-1.5 px-2.5 hover-bg-main text-xs cursor-pointer"
        @click="routerPush('album', item.id)"
      >
        <span class="text-emerald-500">{{ item.name }}</span>
        <span class="pl-1"> - {{ item.artist.name }}</span>
      </div>

      <div
        v-for="item in suggestData.artists"
        :key="item.id"
        v-if="order === 'artists'"
        class="py-1.5 px-2.5 hover-bg-main text-xs cursor-pointer flex items-center"
        @click="routerPush('artistDetail', item.id)"
      >
        <el-avatar size="small" :src="item.img1v1Url" />
        <span class="text-emerald-500 ml-2">{{ item.name }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useSearchStore } from '@/stores/search'
import { usePlayerStore } from '@/stores/player'

const { suggestData, showSearchView } = storeToRefs(useSearchStore())

const { play } = usePlayerStore()
const router = useRouter()

const routerPush = (name: string, id: number) => {
  router.push({ name: name, query: { id: id } }).then(() => {
    showSearchView.value = false
  })
}

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
</script>
<style lang="scss"></style>
