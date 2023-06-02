import { defineStore } from 'pinia'
import { apiSwiper } from '@/utils/api/swiper'
import type { Swiper } from '@/models/swiper'

export const useSwiperStore = defineStore('swiper', {
  state: () => ({
    swipers: [] as Swiper[]
  }),
  getters: {},
  actions: {
    async updateSwipers() {
      if (this.swipers.length) return
      //加载数据
      this.swipers = await apiSwiper()
    }
  }
})
