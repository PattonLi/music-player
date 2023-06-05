<template>
  <div>
    <!-- 选项 -->
    <div class="pb-5">
      <!-- 每一行 -->
      <div v-for="option in options" :key="option.key" class="flex text-xs mb-5">
        <div class="ml-3">
          <el-space wrap :size="10" :spacer="spacer">
            <!-- 每一列 -->
            <div
              type="text"
              class="hover-text px-1 py-0.5"
              :class="{
                active:
                  (item.key === pageData.type && option.key === 'type') ||
                  (item.key === pageData.location && option.key === 'area') ||
                  (item.key === pageData.firstLetter && option.key === 'fistLetter')
              }"
              v-for="(item, index) in option.list"
              :key="index"
              @click="optionChange(option.key, item.key)"
            >
              {{ item.name }}
            </div>
          </el-space>
        </div>
      </div>
    </div>
    <!-- 歌手展示 -->
    <div class="grid grid-flow-row grid-cols-6 xl:grid-cols-8 2xl:grid-cols-10 gap-5">
      <div
        v-for="(artist, index) in artists"
        :key="index"
        class="flex items-center flex-col"
        @click="routerPushByNameId(Pages.artistDetail, artist.artistId)"
      >
        <img
          :src="artist.picUrl"
          alt="歌手头像"
          class="rounded-full cursor-pointer w-full aspect-square object-cover bg-dc"
        />
        <div class="mt-2 text-sm">{{ artist.artist }}</div>
      </div>
    </div>
    <!-- 加载按钮 -->
    <div class="flex justify-center py-5" v-if="artists.length > 0 && !pageData.noMore">
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
import type { Artist } from '@/models/artist'
import { apiArtistList } from '@/utils/api/artist'
import { routerPushByNameId } from '@/utils/navigator/router'
import { Pages } from '@/router/pages'
import { ElDivider } from 'element-plus'
import { AlertError } from '@/utils/alert/AlertPop'

const spacer = h(ElDivider, { direction: 'vertical' })
const artists = ref<Artist[]>([])
const pageData = reactive({
  page: 0,
  pageSize: 10,
  pageTotal: 1,

  //是否显示加载动画
  loading: false,
  noMore: false,
  //查询信息
  firstLetter: '0',
  type: 0,
  location: 0
})

//分页查询
const pageGet = async () => {
  pageData.loading = true
  const res = await apiArtistList(
    pageData.pageSize,
    pageData.page,
    pageData.firstLetter,
    pageData.type,
    pageData.location
  )
  if (res.code == 200) {
    //判断是否已经没有页数了
    if (pageData.page >= pageData.pageTotal) {
      //所有数据已经取完
      pageData.noMore = true
      return
    }
    if (pageData.page == 0) {
      //初始时设置数据
      artists.value = res.artists
    } else {
      //否则push
      artists.value.push(...res.artists)
    }
    //当前位置页数加1
    pageData.page++
    //更新pageSize
    pageData.pageTotal = res.pageTotal
  } else {
    AlertError('抱歉，没有符合筛选条件的歌手！')
  }
  //加载动画
  pageData.loading = false
}

onMounted(async () => {
  pageGet()
})

const optionChange = (keyName: string, keyValue: number | string) => {
  pageData.page = 0
  pageData.pageTotal = 0
  if (keyName === 'location') pageData.location = keyValue as number
  if (keyName === 'type') pageData.type = keyValue as number
  if (keyName === 'firstLetter') pageData.firstLetter = keyValue as string
  pageGet()
}

//筛选项
interface Option {
  key: string
  list: {
    key: string | number
    name: string
  }[]
}

const options: Option[] = [
  {
    key: 'location',
    list: [
      { key: 0, name: '全部' },
      { key: 1, name: '华语' },
      { key: 2, name: '欧美' },
      { key: 3, name: '日本' },
      { key: 4, name: '韩国' }
    ]
  },
  {
    key: 'type',
    list: [
      { key: 0, name: '全部' },
      { key: 1, name: '男' },
      { key: 2, name: '女' },
      { key: 3, name: '组合' }
    ]
  },
  {
    key: 'firstLetter',
    list: [
      { key: '0', name: '全部' },
      { key: 'a', name: 'A' },
      { key: 'b', name: 'B' },
      { key: 'c', name: 'C' },
      { key: 'd', name: 'D' },
      { key: 'e', name: 'E' },
      { key: 'f', name: 'F' },
      { key: 'g', name: 'G' },
      { key: 'h', name: 'H' },
      { key: 'i', name: 'I' },
      { key: 'j', name: 'J' },
      { key: 'k', name: 'K' },
      { key: 'l', name: 'L' },
      { key: 'm', name: 'M' },
      { key: 'n', name: 'N' },
      { key: 'o', name: 'O' },
      { key: 'p', name: 'P' },
      { key: 'q', name: 'Q' },
      { key: 'r', name: 'R' },
      { key: 's', name: 'S' },
      { key: 't', name: 'T' },
      { key: 'u', name: 'U' },
      { key: 'v', name: 'V' },
      { key: 'w', name: 'W' },
      { key: 'x', name: 'X' },
      { key: 'y', name: 'Y' },
      { key: 'z', name: 'Z' },
      { key: '#', name: '#' }
    ]
  }
]
</script>

<style lang="scss" scoped>
.active {
  @apply bg-emerald-400 text-white rounded;
}
</style>
