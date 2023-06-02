<template>
  <div>
    <MyTitle title="专辑推荐" />
    <div class="mb-4 grid grid-flow-row gap-x-20 grid-cols-3 2xl:grid-cols-5 gap-y-9">
      <!-- 循环 -->
      <div
        v-for="(item, index) in _.sampleSize(personalizedAlbums,15)"
        :key="index"
        @click="router.push({ name: 'library/album', query: { id: item.albumId } })"
      >
        <!-- 封面组件 -->
        <MyCover
          :name="item.album"
          :pic-url="item.picUrl"
          :showPlayCount="true"
          :play-count="item.size"
        />
        <!-- 专辑名 -->
        <div class="mx-2 mt-2 text-lg text-main truncate">
          {{ item.album }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import MyTitle from '@/components/common/MyTitle.vue'
import { useMusicStore } from '@/stores/music'
import MyCover from '@/components/common/MyCover.vue'
import _ from 'lodash'
const router = useRouter()

const { personalizedAlbums } = toRefs(useMusicStore())
const { UpdatePersonalize } = useMusicStore()

onMounted(async () => {
  await UpdatePersonalize(2)
})
</script>
