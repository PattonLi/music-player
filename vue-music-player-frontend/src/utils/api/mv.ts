import type { Mv } from '@/models/mv'
import myAxios from './myAxios'

export async function apiGetMv(movieId: number) {
  const data = await myAxios.get<{
    code: number
    mv: Mv
  }>('detail/mv', { movieId: movieId })
  return data
}
