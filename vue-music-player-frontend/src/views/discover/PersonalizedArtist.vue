<template>
  <MyTitle title="推荐歌手" />
  <div
    class="mt-4 mx-6 mb-16 grid grid-flow-row gap-x-20 cursor-pointer grid-cols-3 2xl:grid-cols-5 gap-y-20"
  >
    <!-- 循环 -->
    <div
      v-for="(item, index) in _.sampleSize(personalizedArtists, 15)"
      :key="index"
      class="transition-all flex flex-col items-center justify-center"
      @click="router.push({ name: 'artist', query: { id: item.artistId } })"
    >
      <!-- 第一列图片 -->
      <div class="w-50 h-50 flex-1">
        <img
          :src="item.picUrl"
          alt="歌曲图片"
          class="w-full h-full object-cover rounded-full"
          style="object-fit: cover; aspect-ratio: 1/1"
        />
      </div>

      <!-- 第二列文字信息 -->
      <div class="px-3 pt-1 truncate flex flex-col">
        <div class="text-xl truncate flex justify-center">
          {{ item.artist }}
        </div>
        <div class="text-xl text-dc truncate flex justify-center">
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
const { updatePersonalize: UpdatePersonalize } = useMusicStore()
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
