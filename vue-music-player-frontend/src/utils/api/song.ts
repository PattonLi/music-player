import type { SongUrl } from '@/models/song_url'
import myAxios from './myAxios'
import type { Song } from '@/models/song'

export async function useSongUrl(id: number) {
  const { data } = await myAxios.get<{ data: SongUrl[] }>('/song/url', { id: id })
  return data.first()
}

export async function useDetail(id: number) {
  const { songs } = await myAxios.get<{ songs: Song[] }>('/song/detail', { ids: id })
  return songs.first()
}
