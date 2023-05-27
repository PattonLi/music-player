import type { SearchHotDetail } from '@/models/search'
import myAxios from './myAxios'
import type { SearchSuggest } from '@/models/search'

//获取搜索结果
export async function apiSearchSuggest(keywords: string) {
  const data = await myAxios.get<SearchSuggest>('search/suggest', {
    keywords: keywords
  })
  return data
}

//热搜结果
export async function apiSearchHotDetail() {
  const data = await myAxios.get<SearchHotDetail[]>('search/hot')
  return data
}
