import type { Album } from '@/models/album'
import myAxios from './myAxios'
import type { Song } from '@/models/song'

//获取专辑详情
export async function apiAlbumDetail(albumId: number) {
  const data = await myAxios.get<{
    code: number
    album: Album
    songs: Song[]
  }>('detail/album', { albumId: albumId })

  return data
}

export async function useArtistAlbum(id: number, limit: number = 10, offset: number = 0) {
  return await myAxios.get<{
    code: number
    hotAlbums: Album[]
  }>('detail/artist/album', { id: id, limit: limit, offset: offset })
}
