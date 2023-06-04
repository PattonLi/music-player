import { defineStore } from 'pinia'
import { apiSearchSuggest, apiSearchHotDetail } from '@/utils/api/search'
import type { SearchHotDetail, SearchResults } from '@/models/search'
import { AlertError } from '@/utils/alert/AlertPop'

//搜索结果状态
export const useSearchStore = defineStore('search', {
  state: () => {
    return {
      showSearchView: false,
      //搜索关键词
      searchKeyword: '',
      //搜索信息
      suggestData: {} as SearchResults,
      //热搜结果
      searchHotData: [] as SearchHotDetail[]
    }
  },
  getters: {
    showHot: (state) => {
      return state.searchKeyword == ''
    }
  },
  actions: {
    async updateSuggest() {
      const res = await apiSearchSuggest(this.searchKeyword)
      if (res.code == 200) {
        this.suggestData = {
          albums: res.albums,
          artists: res.artists,
          songs: res.songs
        }
      } else {
        AlertError('获取热搜失败')
      }
    },
    async updateSearchHot() {
      const res = await apiSearchHotDetail()
      if (res.code == 200) {
        this.searchHotData = res.searchHot
      } else {
        AlertError('获取搜索失败')
      }
    }
  }
})
