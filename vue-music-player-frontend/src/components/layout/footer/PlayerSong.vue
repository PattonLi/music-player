<template>
  <div class="flex player-song">
    <!-- 歌曲封面 -->
    <img
      alt=""
      class="w-11 h-11 rounded song_cover"
      :src="song.picUrl || albumLogo"
      @click="changePlayerShow"
    />
    <div class="ml-3 text-sm flex flex-col justify-between">
      <!-- 歌曲信息 -->
      <div class="w-52 2xl:w-96 cursor-pointer truncate">
        <div class="flex">
          <span>{{ song.name || '云音乐' }}</span>
          <span class="ml-2 text-dc">- {{ song.name || `unknow artist` }}</span>
        </div>
      </div>
      <!-- 功能区 -->
      <div class="flex gap-x-3">
        <IconPark :icon="Like" size="18" :stroke-width="3" class="text-slate-400 hover-text" />
        <IconPark :icon="DownTwo" size="18" :stroke-width="3" class="text-slate-400 hover-text" @click="downloadSong(song)"/>
        <IconPark :icon="MoreTwo" size="18" :stroke-width="3" class="text-slate-400 hover-text" />
        <el-badge :value="1000" :max="999" class="badge">
          <IconPark :icon="Comment" size="18" :stroke-width="3" class="text-slate-400 hover-text" />
        </el-badge>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Like, DownTwo, MoreTwo, Comment } from '@icon-park/vue-next'
import { usePlayerStore } from '@/stores/player'
import albumLogo from '@/assets/images/albumLogo.png'
import IconPark from '@/components/common/IconPark.vue'
import { downloadSong } from '@/utils/download/dowmload'
const { changePlayerShow } = usePlayerStore()


const { song } = toRefs(usePlayerStore())
</script>

<style lang="scss">
.song_cover {
  @apply hover:opacity-70 cursor-pointer;
}

.player-song {
  .badge {
    //评论数量样式
    .el-badge__content {
      @apply scale-75 left-1 bg-stone-100 text-slate-500 bg-opacity-80 right-auto;
      @apply dark:bg-stone-900;
    }
  }
}
</style>
