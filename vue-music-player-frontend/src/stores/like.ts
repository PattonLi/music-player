import { defineStore } from 'pinia'
import { AlertError } from '@/utils/alert/AlertPop'
import { apiUserLike } from '@/utils/api/like'
import type { Album } from '@/models/album'
import type { Artist } from '@/models/artist'
import type { Song } from '@/models/song'

export const useLikeStore = defineStore('like', {
  state: () => ({
    albums: [] as Album[],
    artists: [] as Artist[],
    songs: [] as Song[]
  }),
  getters: {},
  actions: {
    async updateLikes(userId: number) {
      //加载数据
      const res = await apiUserLike(userId)
      if (res.code == 200) {
        this.albums = res.albums
        this.artists = res.artists
        this.songs = res.songs
      } else {
        AlertError('获取用户收藏')
      }
    }
  }
})
