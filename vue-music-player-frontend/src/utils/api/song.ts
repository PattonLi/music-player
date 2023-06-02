import myAxios from './myAxios'
import type { Song } from '@/models/song'

export async function apiGetSong(id: number) {
  const data = await myAxios.get<{
    code: number
    song: Song
  }>('/song/url', { id: id })
  return data
}
