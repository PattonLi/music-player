<template>
  <div class="grid grid-flow-row grid-cols-8 2xl:grid-cols-12 gap-2.5">
    <div class="button-dc">全部</div>
    <div v-for="(tag, index) in tags" :key="index" class="button-dc">
      <button>{{ tag.tagName }}</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { PlayListTag } from '@/models/playlist'

const tags = ref<PlayListTag[]>([])

onMounted(() => {
  tags.value = [
    {
      tagName: 'ACG'
    },
    {
      tagName: '古风'
    },
    {
      tagName: 'ACG'
    },
    {
      tagName: 'ACG'
    },
    {
      tagName: 'ACG'
    },
    {
      tagName: 'ACG'
    },
    {
      tagName: 'ACG'
    }
  ]
})
</script>

<style lang="scss"></style>
