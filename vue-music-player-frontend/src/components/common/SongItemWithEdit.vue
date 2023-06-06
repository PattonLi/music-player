<template>
  <div
    class="flex song-item items-center w-full hover-bg-main"
    :class="{ playing: id == propSong.songId }"
    @dblclick="play(propSong.songId)"
  >
    <!-- 左部 -->
    <div class="flex-shrink-0 flex-1 flex items-center justify-between pr-5">
      <!-- 歌曲编号 -->
      <div class="text-sm ml-0.5">
        <span>{{ _.padStart(order.toString(), 2, '0') }}</span>
      </div>
      <!-- 喜欢按钮 -->
      <div class="items-center flex flex-1 flex-shrink-0">
        <IconPark
          :icon="Like"
          size="20"
          :stroke-width="3"
          class="text-gray-400 ml-4 mr-2 cursor-pointer hover:text-red-400"
        />
        <!-- 歌曲名 -->
        <div class="truncate text-lg" style="max-width: 40%">
          <small>{{ propSong.name }}</small>
        </div>
        <!-- 随机返回一个1-10整数 -->
        <IconPark
          v-if="true"
          class="ml-3 text-gray-400 cursor-pointer hover-text"
          size="16"
          :icon="PlayTwo"
          @click="routerPushByNameId(Pages.mvDetail, _.random(1, 10))"
        />
      </div>

      <!-- 中间操作栏 -->
      <div class="hidden icon-action flex-shrink-0">
        <div class="flex items-center gap-x-2 text-gray-400 ml-2">
          <IconPark
            title="播放"
            :icon="PlayOne"
            size="24"
            class="hover-text"
            @click="play(propSong.songId)"
          />
          <IconPark :icon="Add" size="20" class="hover-text" />
          <IconPark :icon="DownTwo" size="20" class="hover-text" />
          <IconPark :icon="MoreTwo" size="20" class="hover-text" />
        </div>
      </div>
    </div>

    <!-- 歌手名 -->
    <div class="text-lg flex-shrink-0 w-1/4 items-center">
      <small
        class="truncate hover-text"
        @click="routerPushByNameId(Pages.artistDetail, propSong.artistId)"
      >
        {{ propSong.artist }}
      </small>
    </div>

    <!-- 专辑 -->
    <div class="text-lg flex-shrink-0 w-1/4 items-center">
      <small
        class="truncate hover-text"
        @click="routerPushByNameId(Pages.albumDetail, propSong.artistId)"
      >
        {{ propSong.album }}
      </small>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Add, DownTwo, Like, MoreTwo, PlayOne, PlayTwo } from '@icon-park/vue-next'
import { usePlayerStore } from '@/stores/player'
import IconPark from '@/components/common/IconPark.vue'
import type { Song } from '@/models/song'
import { Pages } from '@/router/pages'
import { routerPushByNameId } from '@/utils/navigator/router'
import { storeToRefs } from 'pinia'
import _ from 'lodash'

defineProps<{
  propSong: Song
  order: number
}>()
const { play } = usePlayerStore()
const { song } = storeToRefs(usePlayerStore())
const id = song.value.songId
</script>

<style lang="scss" scoped>
//内联操作图标
.song-item {
  @apply py-2.5 pl-1;
  &:hover {
    .icon-action {
      @apply inline-block;
    }
  }
}
//在播放时
.playing {
  @apply bg-emerald-50 dark:bg-stone-800;
}
</style>
