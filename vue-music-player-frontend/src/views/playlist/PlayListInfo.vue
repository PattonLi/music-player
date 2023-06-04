<template>
  <div class="flex items-stretch">
    <!-- 封面 -->
    <img
      :src="playlist.picUrl"
      alt="专辑封面"
      class="w-44 h-44 object-cover rounded-xl flex-shrink-0"
    />

    <!-- 信息 -->
    <div class="pl-5 flex flex-col justify-between py-1 flex-1">
      <div>
        <!-- 歌单名 -->
        <div class="text-2xl font-bold">{{ playlist.playList }}</div>
        <div class="flex items-center text-xs text-dc pb-2 pt-3">
          <!-- 创建用户信息 -->
          <span class="ml-2">{{ playlist.user }}</span>
          <!-- 标签 -->
          <span class="ml-5 flex text-dc hover-text">
            <span>{{ playlist.tag }}</span>
          </span>
        </div>

        <!-- 描述 -->
        <div class="text-xs text-gray-500 leading-normal">
          <span :text="playlist.profile" :end="90" />
        </div>
      </div>

      <!-- 功能键 -->
      <div class="justify-self-stretch mt-5 gap-x-2 flex items-center">
        <button class="w-32 button" @click="playAll">
          <IconPark :icon="PlayOne" size="22" class="mr-1" />
          <span>播放全部</span>
        </button>
        <button class="w-32 button-outline">
          <IconPark :icon="Like" size="18" class="mr-1" />
          <span>收藏</span>
        </button>
        <button class="button-outline w-8">
          <IconPark :icon="More" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { PlayOne, Like, More } from '@icon-park/vue-next'

import IconPark from '@/components/common/IconPark.vue'
import type { PlayList } from '@/models/playlist'

defineProps<{
  playlist: PlayList
  playAll?: () => void
}>()
</script>
