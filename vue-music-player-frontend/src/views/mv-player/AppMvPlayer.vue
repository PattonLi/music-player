<template>
  <div class="p-5" v-if="mv">
    <div class="flex gap-x-5">
      <div class="flex-1">
        <video class="aspect-video w-full" :src="mv.url" autoplay controls />
      </div>
      <div class="hidden w-80 flex-shrink-0 xl:block">
        <div>大家都在看</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useMvStore } from '@/stores/mv'
import { usePlayerStore } from '@/stores/player'
import { storeToRefs } from 'pinia'
import _ from 'lodash'

const route = useRoute()
const { mvs } = storeToRefs(useMvStore())
const id = Number(route.query.id)
const { togglePlay } = usePlayerStore()
const mv = ref()

onMounted(async () => {
  let index = _.findIndex(mvs.value, (o) => {
    return o.movieId == id
  })
  mv.value = mvs.value[index]
  //暂停音乐播放
  togglePlay()
})

onUnmounted(() => {
  //恢复音乐播放
  setTimeout(() => {
    togglePlay()
  }, 1000)
})
</script>
<style lang="scss"></style>
