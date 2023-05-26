import type { SearchHotDetail } from '@/stores/search'
import myAxios from './myAxios'
import type { SearchSuggest } from '@/models/search'

//获取搜索结果
export async function apiSearchSuggest(keywords: string) {
  const { result } = await myAxios.get<{ result: SearchSuggest }>('search/suggest', {
    keywords: keywords
  })
  return result
}

export async function apiSearchHotDetail() {
  const {data} = await myAxios.get<{ data: SearchHotDetail[] }>('search/hot/detail')
  return data
}
