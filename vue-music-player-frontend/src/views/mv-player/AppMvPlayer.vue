<template>
  <div class="p-5" v-if="mv?.url">
    <div class="flex gap-x-5">
      <div class="flex-1">
        <video class="aspect-video w-full" :src="mv?.url" autoplay controls />
      </div>
      <div class="hidden w-80 flex-shrink-0 xl:block">
        <div>大家都在看</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { apiGetMv } from '@/utils/api/mv'
import type { Mv } from '@/models/mv'
import { usePlayerStore } from '@/stores/player'
import { AlertError } from '@/utils/alert/AlertPop'

const route = useRoute()
const id = Number(route.query.id)

const { togglePlay } = usePlayerStore()

const mv = ref<Mv>()

onMounted(async () => {
  const res = await apiGetMv(id)
  if (res.code == 200) {
    mv.value = res.mv
  } else {
    AlertError('获取mv url失败')
  }
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
