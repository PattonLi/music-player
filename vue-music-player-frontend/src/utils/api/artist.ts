import type { Artist, ArtistDesc } from '@/models/artist'
import myAxios from './myAxios'
import type { Song } from '@/models/song'

//获取详情
export async function apiArtistDetail(artistId: number) {
  const data = await myAxios.get<{
    code: number
    artistDetail: Artist
  }>('detail/artist', { artistId: artistId })
  return data
}

//获取歌手描述
export async function apiArtistDesc(artistId: number) {
  const data = await myAxios.get<{
    code: number
    describe: ArtistDesc
  }>('detail/artist/describe', { artistId: artistId })
  return data
}

export async function useArtistSongs(
  id: number,
  order: string = 'time',
  limit: number = 10,
  offset: number = 0
) {
  return await myAxios.get<{
    code: number
    songs: Song[]
  }>('detail/artist/songs', { artistId: id, order: order, limit: limit, offset: offset })
}
