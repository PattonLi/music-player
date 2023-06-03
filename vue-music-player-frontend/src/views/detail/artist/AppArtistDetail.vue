<template>
  <div class="artist-detail p-5" v-if="artistDetail">
    <ArtistInfo :artist-detail="artistDetail" />

    <el-tabs class="mt-3" v-model="tab">
      <el-tab-pane lazy label="精选" name="choice"> </el-tab-pane>
      <el-tab-pane lazy :label="`歌曲 ${artistDetail.artist.songSize}`" name="music">
        <ArtistSong :id="id" />
      </el-tab-pane>
      <el-tab-pane lazy :label="`专辑 ${artistDetail.artist.albumSize}`" name="album">
        <ArtistAlbum :id="id" />
      </el-tab-pane>
      <el-tab-pane lazy :label="`视频 ${22}`" name="mv">
        <ArtistMv :id="id" />
      </el-tab-pane>
      <el-tab-pane lazy label="详情" name="desc">
        <ArtistProfile :id="id" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import { apiArtistDetail } from '@/utils/api/artist'
import type { ArtistDetail } from '@/models/artist'
import ArtistInfo from './ArtistInfo.vue'
import ArtistProfile from './ArtistProfile.vue'
import ArtistSong from './ArtistSong.vue'
import ArtistAlbum from './ArtistAlbum.vue'
import ArtistMv from './ArtistMV.vue'
import { AlertError } from '@/utils/alert/AlertPop'

const route = useRoute()
const id = Number(route.query.id)
const artistDetail = ref<ArtistDetail>()
const tab = ref('music')

onMounted(async () => {
  const res = await apiArtistDetail(id)
  if (res.code == 200) {
    artistDetail.value = res.artistDetail
  } else {
    AlertError('获取歌手详情失败')
  }
})
</script>
<style lang="scss">
.artist-detail {
  .el-tabs__nav-wrap::after {
    height: 0;
  }

  .el-tabs__header {
    @apply m-0;
  }
}
</style>
