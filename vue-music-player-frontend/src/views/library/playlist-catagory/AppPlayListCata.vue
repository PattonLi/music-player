<template>
  <div>
    <HotPlayList />
    <div class="py-5 text-xl">歌单</div>

    <div class="gap-5 grid grid-flow-row grid-cols-3 lg:grid-cols-5 2xl:grid-cols-7">
      <div
        v-for="(item, index) in playLists"
        :key="index"
        :class="{ 'item-1': index === 0 }"
        @click="routerPushByNameId(Pages.playlistDetail, item.playListId)"
      >
        <MyCover
          :name="item.playList"
          :pic-url="item.picUrl"
          :play-count="item.size"
          :show-play-count="true"
        />
        <div class="mt-2 text-xs text-main leading-5">{{ item.playList }}</div>
        <div class="mt-2 text-xs text-main truncate text-dc">{{ item.user }}</div>
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
import type { PlayList } from '@/models/playlist'
import { apiHotPlayList } from '@/utils/api/playlist'
import MyCover from '@/components/common/MyCover.vue'
import { Pages } from '@/router/pages'
import { routerPushByNameId } from '@/utils/navigator/router'
import { AlertError } from '@/utils/alert/AlertPop'

const playLists = ref<PlayList[]>([])

const pageData = reactive({
  page: 0,
  pageSize: 10,
  pageTotal: 1,

  //是否显示加载动画
  loading: false,
  noMore: false
})

//分页查询
const pageGet = async () => {
  pageData.loading = true
  const res = await apiHotPlayList(pageData.pageSize, pageData.page)
  if (res.code == 200) {
    //判断是否已经没有页数了
    if (pageData.page >= pageData.pageTotal) {
      //所有数据已经取完
      pageData.noMore = true
      return
    }
    if (pageData.page == 0) {
      //初始时设置数据
      playLists.value = res.playlist
    } else {
      //否则push
      playLists.value.push(...res.playlist)
    }
    //当前位置页数加1
    pageData.page++
    //更新pageSize
    pageData.pageTotal = res.pageTotal
  } else {
    AlertError('抱歉，没有找到歌单！')
  }
  //加载动画
  pageData.loading = false
}

onMounted(async () => {
  pageGet()
})
</script>
<style lang="scss"></style>
