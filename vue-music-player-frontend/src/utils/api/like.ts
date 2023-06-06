import myAxios from './myAxios'
import type { Album } from '@/models/album'
import type { Artist } from '@/models/artist'
import type { Song } from '@/models/song'

//获取用户收藏
export async function apiUserLike(userId: number) {
  const data = await myAxios.get<{
    code: number
    albums: Album[]
    artists: Artist[]
    songs: Song[]
  }>('user/like', {
    userId: userId
  })
  return data
}
