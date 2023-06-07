import { defineStore } from 'pinia'
import type { Song } from '@/models/song'

export const useRecentPlayStore = defineStore('recentPlay', {
  state: () => ({
    songs: [] as Song[]
  }),
  getters: {},
  actions: {}
})
