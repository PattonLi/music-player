import myAxios from './myAxios'
import type { Swiper } from '@/models/swiper'

//获取走马灯
export async function apiSwiper() {
  const { swiper } = await myAxios.get<{ swiper: Swiper[] }>('/banner', { type: 1 })
  return swiper
}
