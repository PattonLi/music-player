<template>
  <div class="px-5">
    <h1 class="text-3xl font-bold pt-8 pb-4">推荐</h1>
    <MySwiper :per-page="3" />
    <PersonalizedSong />
    <PersonalizedAlbum />
    <PersonalizedArtist />
    <Mv />
  </div>
</template>

<script setup lang="ts">
import PersonalizedSong from '@/views/discover/PersonalizedSong.vue'
import PersonalizedAlbum from '@/views/discover/PersonalizedAlbum.vue'
import PersonalizedArtist from '@/views/discover/PersonalizedArtist.vue'
import MySwiper from '@/components/common/MySwiper.vue'
</script>

<style lang="scss"></style>
