<template>
  <div>
    <!-- 评论标题 -->
    <div class="text-xl text-main ml-8 mt-0 flex items-center">
      <div>精彩评论</div>
      <IconPark :icon="MessageEmoji" size="22" :stroke-width="3" class="ml-1" />
    </div>
    <div class="text-main ml-8 mt-4 flex items-center" v-for="(item, index) in comments" :key="index">
      <CommentItem :comment="item"></CommentItem>
    </div>
  </div>
</template>

<script setup lang="ts">
import CommentItem from '@/components/common/CommentItem.vue'
import { MessageEmoji } from '@icon-park/vue-next'
import IconPark from '@/components/common/IconPark.vue'
import { apiGetSongComment } from '@/utils/api/comment'
import { AlertError } from '@/utils/alert/AlertPop'
import type { Comment } from '@/models/comment'

const props = defineProps({
  songId: Number
})
const comments = ref<Comment[]>()
onMounted(async()=>{
  const data = await apiGetSongComment(props.songId!)
  if(data.code==200){
    comments.value=data.comments
  }else{
    AlertError('获取歌曲评论失败')
  }
})
</script>
