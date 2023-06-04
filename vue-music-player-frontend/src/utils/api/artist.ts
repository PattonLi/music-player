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

//分页查询歌手歌曲
export async function apiArtistSongs(
  artistId: number,
  pageSize: number,
  currentPage: number,
  order: string
) {
  return await myAxios.get<{
    code: number
    pageTotal: number
    songs: Song[]
  }>('detail/artist/songs', {
    artistId: artistId,
    pageSize: pageSize,
    currentPage: currentPage,
    order: order
  })
}
