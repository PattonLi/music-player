<template>
  <!-- 歌曲循环推荐 -->
  <Swiper slides-per-group-auto slides-per-view="auto" :navigation="true" :grab-cursor="true">
    <SwiperSlide v-for="(item, index) in swipers" :key="index">
      <img :src="item.pic_url" class="banner-image" @click="onClick(item)" />
    </SwiperSlide>
  </Swiper>
</template>

<script setup lang="ts">
//swiper引入
import { Swiper, SwiperSlide } from 'swiper/vue'
import 'swiper/css'
import { usePlayerStore } from '@/stores/player'
import { useSwiperStore } from '@/stores/swiper'
import type { Swiper as MySwiper } from '@/models/swiper'

const { swipers } = toRefs(useSwiperStore())
const { updateSwipers } = useSwiperStore()

onMounted(async () => {
  await updateSwipers()
})
const { play } = usePlayerStore()

const onClick = (swiper: MySwiper) => {
  //为歌曲则播放
  if (swiper.targetType === 1) {
    play(swiper.targetId)
  }
}
</script>

<style lang="scss" scoped>
.swiper {
  @apply -mx-2.5;
  .swiper-slide {
    @apply w-full lg:w-1/2 xl:w-1/3 2xl:w-1/4 px-2.5;
  }
}

.banner-image {
  @apply rounded-lg cursor-pointer transition-all object-cover;
  @apply hover:shadow hover:opacity-80;
}
</style>
