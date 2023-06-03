import { defineStore } from 'pinia'
import { apiGetPersonalize } from '@/utils/api/personalized'
import type { PersonalizedSong, Personalizedlbum, PersonalizedArtist } from '@/models/personalized'

export const useMusicStore = defineStore('music', {
  state: () => ({
    personalizedSongs: [] as PersonalizedSong,
    personalizedAlbums: [] as Personalizedlbum,
    personalizedArtists: [] as PersonalizedArtist
  }),
  getters: {},
  actions: {
    async updatePersonalize() {}
  }
})
