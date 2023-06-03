import type { Album } from '@/models/album'
import myAxios from './myAxios'
import type { Song } from '@/models/song'

export async function apiAlbumDetail(albumId: number) {
  const data = await myAxios.get<{
    code: number
    album: Album
    songs: Song[]
  }>('detail/album', { albumId: albumId })

  return data
}
