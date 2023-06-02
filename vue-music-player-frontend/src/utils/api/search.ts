import type {
  SearchAlbums,
  SearchArtists,
  SearchHotDetail,
  SearchSongs,
  SearchUsers
} from '@/models/search'
import myAxios from './myAxios'

//动态获取搜索结果
export async function apiSearchSuggest(keywords: string) {
  const data = await myAxios.get<{
    code: number
    albums: SearchAlbums[]
    artists: SearchArtists[]
    songs: SearchSongs[]
    users: SearchUsers[]
  }>('search/suggest', {
    keywords: keywords
  })
  return data
}

//获取热搜结果
export async function apiSearchHotDetail() {
  const data = await myAxios.get<{
    code: number
    searchHot: SearchHotDetail[]
  }>('search/hot')
  return data
}
