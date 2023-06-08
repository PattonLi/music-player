<template>
  <div class="flex-col flex h-full pb-28 justify-center">
    <div class="flex items-center">
      <!-- 左边专辑封面 -->
      <div class="w-1/2">
        <div class="flex justify-center">
          <div class="w-1/2 h-1/2">
            <img :src="albumLogo" alt="" />
          </div>
        </div>
      </div>

      <!-- 右边歌词歌曲信息 -->
      <div class="w-1/2">
        <div class="flex justify-center h-full">
          <div class="flex flex-col items-center mr-20">
            <!-- 歌曲标题 -->
            <div class="h-5 mt-9 mb-5">
              <span class="font-bold text-3xl">歌曲名</span>
            </div>
            <div class="h-5">
              <span>{{ songPos }}</span>
            </div>
            <div class="h-5 mb-12">
              <span>{{ artistPos }}</span>
            </div>
            <!-- 歌词显示 -->
            <div class="flex-1">
              <SongLyric></SongLyric>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import albumLogo from '@/assets/images/albumLogo.png'
import { usePlayerStore } from '@/stores/player'
import { storeToRefs } from 'pinia'
import { apiGetLyric } from '@/utils/api/song'
import SongLyric from './SongLyric.vue'

const lyric = ref()
const get = () => {
  lyric.value = apiGetLyric(song.value.lyricUrl!)
}

const songPos = computed(() => {
  return song.value.artist ? '歌手：' + song.value.artist : ''
})
const artistPos = computed(() => {
  return song.value.album ? '专辑：' + song.value.album : ''
})

const { song } = storeToRefs(usePlayerStore())
</script>
