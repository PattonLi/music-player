<template>
  <!-- 歌曲信息 -->
  <div class="text-xl text-main ml-8 mt-2 flex items-center">
    <div>歌曲信息</div>
  </div>
  <div class="flex ml-8 mt-4">
    <div class="w-20 h-20">
      <img class="rounded-md" :src="song?.picUrl" alt="songImg" />
    </div>
    <div class="flex flex-col ml-5 my-2 justify-between">
      <div>
        <span class="text-xl">{{ song?.name }}</span>
      </div>
      <div class="flex">
        专辑：<span class="text-dc mr-6">{{ song?.album }}</span>
        歌手：<span class="text-dc">{{ song?.artist }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Song } from '@/models/song'
import { AlertError } from '@/utils/alert/AlertPop'
import { apiGetSong } from '@/utils/api/song'

const props = defineProps({
  songId: Number
})
const song = ref<Song>()
onMounted(async () => {
  const data = await apiGetSong(props.songId!)
  if (data.code == 200) {
    song.value = data.song
  } else {
    AlertError('获取歌曲信息失败')
  }
})
</script>
