import type { PlayList } from '@/models/playlist'
import myAxios from './myAxios'
import type { Song } from '@/models/song'

export async function apiPlayListDetail(playListId: number) {
  const data = await myAxios.get<{
    code: number
    playlist: PlayList
    songs: Song[]
  }>('/playlist/detail', { playListId: playListId })
  return data
}
