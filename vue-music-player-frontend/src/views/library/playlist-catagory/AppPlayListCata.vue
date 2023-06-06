<template>
  <div>
    <HotPlayList />
    <div class="py-5 text-xl">歌单</div>

    <div class="gap-5 grid grid-flow-row grid-cols-3 lg:grid-cols-5 2xl:grid-cols-7">
      <div
        v-for="(item, index) in playLists"
        :key="index"
        @click="routerPushByNameId(Pages.playlistDetail, item.playListId)"
      >
        <MyCover
          :name="item.playList"
          :pic-url="item.picUrl"
          :play-count="item.size"
          :show-play-count="true"
        />
        <div class="mt-1 ml-1 text-sm text-main leading-5">{{ item.playList }}</div>
        <div class="mt-0 ml-1 text-sm text-main truncate text-dc">{{ item.user }}</div>
      </div>
    </div>
    <!-- 加载按钮 -->
    <div class="flex justify-center py-5" v-if="playLists.length > 0 && !pageData.noMore">
      <el-button
        :loading="pageData.loading"
        link
        size="large"
        class="text-center"
        @click="pageGet()"
        >加载更多
      </el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import HotPlayList from '@/views/library/playlist-catagory/HotPlayList.vue'
import MyCover from '@/components/common/MyCover.vue'
import { Pages } from '@/router/pages'
import { routerPushByNameId } from '@/utils/navigator/router'
import { usePlayListStore } from '@/stores/playlist'
import { storeToRefs } from 'pinia'

const { pageData, playLists } = storeToRefs(usePlayListStore())
const { pageGet } = usePlayListStore()
onMounted(async () => {
  if (pageData.value.page == 0) {
    pageGet()
  }
})
</script>
<style lang="scss"></style>
