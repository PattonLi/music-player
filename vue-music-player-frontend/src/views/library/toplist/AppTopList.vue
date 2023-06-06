<template>
  <div>
    <!-- 官方榜 -->
    <div class="text-xl pb-5 font-bold">{{ topListData[0].type }}</div>
    <!-- 网格 -->
    <div class="grid grid-flow-row grid-cols-2 2xl:grid-cols-4 gap-5">
      <div
        v-for="(item, index) in topListData[0].topList"
        :key="index"
        class="flex bg-dc rounded-lg items-center cursor-pointer"
        @click="toPlaylist(item)"
      >
        <!-- 封面 -->
        <MyCover
          :name="item.name"
          :pic-url="item.picUrl"
          :play-count="item.listeners"
          class="w-36 flex-shrink-0"
          show-play-count
        />

        <div class="px-5 flex-1 flex-shrink-0 flex flex-col">
          <!-- 排行榜名字 -->
          <div class="text-xl font-bold">{{ item.name }}</div>
          <!-- 3条显示 -->
          <div class="text-xs text-main mt-2">
            <div
              v-for="(song, index) in _.sampleSize(personalizedSongs, 3)"
              :key="index"
              class="mt-2"
            >
              <div class="flex">
                <span class="mr-1">{{ index + 1 }}</span>
                <div class="flex-auto w-20 truncate">
                  <span>{{ song.name }}</span>
                  <span>-</span>
                  <span>{{ song.artist }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 地区榜 -->
    <div class="text-xl py-5 font-bold">{{ topListData[1].type }}</div>
    <!-- 网格 -->
    <div class="grid grid-flow-row grid-cols-5 2xl:grid-cols-10 gap-5">
      <div v-for="(item, index) in topListData[1].topList" :key="index" @click="toPlaylist(item)">
        <MyCover :name="item.name" :pic-url="item.picUrl" :play-count="item.listeners" />
        <div class="text-xs mt-2">{{ item.name }}</div>
      </div>
    </div>

    <!-- 全球榜 -->
    <div class="text-xl py-5 font-bold">{{ topListData[2].type }}</div>
    <!-- 网格 -->
    <div class="grid grid-flow-row grid-cols-5 2xl:grid-cols-10 gap-5">
      <div v-for="(item, index) in topListData[2].topList" :key="index" @click="toPlaylist(item)">
        <MyCover :name="item.name" :pic-url="item.picUrl" :play-count="item.listeners" />
        <div class="text-xs mt-2">{{ item.name }}</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import MyCover from '@/components/common/MyCover.vue'
import type { TopList } from '@/models/toplist'
import { Pages } from '@/router/pages'
// import { useTopListStore } from '@/stores/toplist'
import { routerPushByNameId } from '@/utils/navigator/router'
import { useMusicStore } from '@/stores/music'
import { storeToRefs } from 'pinia'
import _ from 'lodash'

// const { topListData } = storeToRefs(useTopListStore())
// const { updateTopListData } = useTopListStore()
const { personalizedSongs } = storeToRefs(useMusicStore())
const { updatePersonalize: UpdatePersonalize } = useMusicStore()

onMounted(async () => {
  // updateTopListData()
  UpdatePersonalize(1)
})

const toPlaylist = (top: TopList) => {
  routerPushByNameId(Pages.playlistDetail, top.topListId)
}

const topListData = ref([
  {
    type: '官方榜',
    topList: [
      {
        topListId: 1,
        listeners: 123.3,
        name: '飙升榜',
        picUrl: 'sss'
      },
      {
        topListId: 2,
        listeners: 123.3,
        name: '热歌榜',
        picUrl: 'sss'
      },
      {
        topListId: 3,
        listeners: 123.3,
        name: '新歌榜',
        picUrl: 'sss'
      },
      {
        topListId: 4,
        listeners: 123.3,
        name: '流行指数榜',
        picUrl: 'sss'
      },
      {
        topListId: 5,
        listeners: 123.3,
        name: '腾讯音乐人原创榜',
        picUrl: 'sss'
      },
      {
        topListId: 6,
        listeners: 123.3,
        name: 'MV榜',
        picUrl: 'sss'
      }
    ]
  },
  {
    type: '地区榜',
    topList: [
      {
        topListId: 1,
        listeners: 123.3,
        name: '飙升榜',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 2,
        listeners: 123.3,
        name: '热歌榜',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 3,
        listeners: 123.3,
        name: '新歌榜',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 4,
        listeners: 123.3,
        name: '流行指数榜',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 5,
        listeners: 123.3,
        name: '腾讯音乐人原创榜',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 6,
        listeners: 123.3,
        name: 'MV榜',
        picUrl: 'asdasdasd'
      }
    ]
  },
  {
    type: '全球榜',
    topList: [
      {
        topListId: 1,
        listeners: 123.3,
        name: 'billboard',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 2,
        listeners: 123.3,
        name: '热歌榜',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 3,
        listeners: 123.3,
        name: '新歌榜',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 4,
        listeners: 123.3,
        name: '流行指数榜',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 5,
        listeners: 123.3,
        name: '腾讯音乐人原创榜',
        picUrl: 'asdasdasd'
      },
      {
        topListId: 6,
        listeners: 123.3,
        name: 'MV榜',
        picUrl: 'asdasdasd'
      }
    ]
  }
])
</script>

<style lang="scss"></style>
