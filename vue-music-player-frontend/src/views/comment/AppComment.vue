<template>
  <SongInfo :song-id="songId"></SongInfo>
  <TextArea :song-id="songId"></TextArea>
  <CommentList :comments="(comments as Comment[])"></CommentList>
</template>

<script setup lang="ts">
import { apiGetSongComment } from '@/utils/api/comment'
import CommentList from './CommentList.vue'
import SongInfo from './SongInfo.vue'
import TextArea from './TextArea.vue'
import { AlertError } from '@/utils/alert/AlertPop'
import type { Comment } from '@/models/comment'

const route = useRoute()
//song id
const songId = Number(route.query.id)
const comments = ref<Comment[]>()

const refreshComment = async () => {
  const data = await apiGetSongComment(songId)
  if (data.code == 200) {
    comments.value = data.comments
  } else {
    AlertError('获取歌曲评论失败')
  }
}
provide('refreshComment', refreshComment)
</script>
