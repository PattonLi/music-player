<template>
  <div>
    <MyTitle title="MV推荐" />
    <div class="mb-10 grid grid-flow-row gap-x-20 grid-cols-3 2xl:grid-cols-5 gap-y-9">
      <!-- 循环 -->
      <div
        v-for="(item, index) in mvs"
        :key="index"
        @click="routerPushByNameId(Pages.mvDetail, item.movieId)"
      >
        <!-- 封面组件 -->
        <MyCover
          :name="item.movie"
          :pic-url="item.picUrl"
          :showPlayCount="true"
          :movie="true"
          :play-count="item.duration"
        />
        <!-- MV名 -->
        <div class="mx-2 mt-2 text-lg text-main truncate">
          {{ item.movie }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import MyTitle from '@/components/common/MyTitle.vue'

import MyCover from '@/components/common/MyCover.vue'
import { routerPushByNameId } from '@/utils/navigator/router'
import { Pages } from '@/router/pages'
import { useMvStore } from '@/stores/mv'
import { storeToRefs } from 'pinia'

const { updateMv } = useMvStore()
const { mvs } = storeToRefs(useMvStore())

onMounted(async () => {
  await updateMv()
})
</script>
