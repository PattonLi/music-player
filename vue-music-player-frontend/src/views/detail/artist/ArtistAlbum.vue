<template>
  <div class="mt-5">
    <!--  -->
    <div class="grid grid-flow-row grid-cols-3 lg:grid-cols-5 gap-5 2xl:grid-cols-8">
      <!-- 循环 -->
      <div
        v-for="(item, index) in list"
        :key="index"
        :class="{ 'item-1': index === 0 }"
        @click="routerPushByNameId(Pages.albumDetail,item.albumId)"
      >
        <!-- 封面 -->
        <MyCover :name="item.album" :pic-url="item.picUrl" :play-count="item.size" />
        <!-- 信息 -->
        <div class="mt-2 text-xs truncate">{{ item.album }}</div>
        <div class="mt-1 text-xs text-slate-400 truncate">{{ toDate(item.publishTime) }}</div>
      </div>
    </div>
    <div class="flex justify-center py-5" v-if="list.length > 0 && !pageData.noMore">
      <el-button
        :loading="pageData.loading"
        type="text"
        class="text-center w-full"
        @click="loadMore"
        >加载更多</el-button
      >
    </div>
  </div>
</template>

<script setup lang="ts">
import { useArtistAlbum } from '@/utils/api/album'
import type { Album } from '@/models/album'
import { toDate } from '@/utils/number/number'
import MyCover from '@/components/common/MyCover.vue'
import { routerPushByNameId } from '@/utils/navigator/router';
import { Pages } from '@/router/pages';

const router = useRouter()

const props = defineProps<{ id: number }>()
const list = ref<Album[]>([])

const pageData = reactive({
  limit: 40,
  page: 1,
  loading: false,
  noMore: false
})

const offset = computed(() => {
  if (pageData.page === 1) return 0
  return list.value.length + pageData.limit
})

const getData = async () => {
  try {
    pageData.loading = true
    const { hotAlbums } = await useArtistAlbum(props.id, pageData.limit, offset.value)
    if (pageData.page === 1) {
      list.value = hotAlbums
    } else {
      list.value.push(...hotAlbums)
    }
    if (hotAlbums.length < pageData.limit) {
      pageData.noMore = true
    }
  } catch (e) {
    pageData.page--
  }
  pageData.loading = false
}

const loadMore = () => {
  pageData.page++
  getData()
}

onMounted(async () => {
  await getData()
})
</script>
<style lang="scss"></style>
