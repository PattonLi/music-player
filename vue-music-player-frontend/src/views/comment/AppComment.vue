<template>
  <SongInfo :song-id="songId"></SongInfo>
  <TextArea :song-id="songId"></TextArea>
  <CommentList :song-id="songId"></CommentList>
</template>

<script setup lang="ts">
import CommentList from './CommentList.vue'
import SongInfo from './SongInfo.vue'
import TextArea from './TextArea.vue'

const route = useRoute()
//song id
const songId = Number(route.query.id)
</script>
