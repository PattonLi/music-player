import type { ArtistDetail } from '@/models/artist'
import myAxios from './myAxios'

export async function apiArtistDetail(artistId: number) {
  const data = await myAxios.get<{
    code: number
    artistDetail: ArtistDetail
  }>('detail/artist', { artistId: artistId })
  return data
}
