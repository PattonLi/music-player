<template>
  <div>
    <div v-for="(item, index) in [1,2,3]" :key="index">
      <CommentItem></CommentItem>
  </div>
  </div>
</template>

<script setup lang="ts">
import CommentItem from '@/components/common/CommentItem.vue';


</script>
