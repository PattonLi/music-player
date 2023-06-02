<template>
  <MyTitle title="推荐歌手" />
  <div
    class="mt-2 mb-16 grid grid-flow-row gap-x-5 cursor-pointer grid-cols-2 2xl:grid-cols-4 gap-y-9"
  >
    <!-- 循环 -->
    <div
      v-for="(item, index) in _.sampleSize(personalizedArtists, 16)"
      :key="index"
      class="transition-all flex items-center"
      @click="router.push({ name: 'artist', query: { id: item.artistId } })"
    >
      <!-- 第一列图片 -->
      <img
        :src="item.picUrl"
        alt="歌曲图片"
        class="w-44 h-44 object-cover rounded-full flex-shrink-0 cover-play-image"
      />
      <!-- 第二列文字信息 -->
      <div class="px-3 truncate flex flex-col">
        <div class="text-xl truncate">
          {{ item.artist }}
        </div>
        <div class="mt-1 text-xl text-dc truncate">
          {{ item.location }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import MyTitle from '@/components/common/MyTitle.vue'
import { useMusicStore } from '@/stores/music'
import _ from 'lodash'

const { personalizedArtists } = toRefs(useMusicStore())
const { UpdatePersonalize } = useMusicStore()
const router = useRouter()

onMounted(async () => {
  await UpdatePersonalize(3)
})
</script>

<style lang="scss" scoped>
.cover-play-image {
  @apply transition-all;
  //位置移动
  @apply hover:-translate-y-3;
}
</style>
