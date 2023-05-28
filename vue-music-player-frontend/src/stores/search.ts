import { defineStore } from 'pinia'
import { apiSearchSuggest, apiSearchHotDetail } from '@/utils/api/search'
import type { SearchHotDetail, SearchSuggest } from '@/models/search'

//搜索结果状态
export const useSearchStore = defineStore('search', {
  state: () => {
    return {
      showSearchView: false,
      //搜索关键词
      searchKeyword: '',
      //
      suggestData: {} as SearchSuggest,
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
      this.suggestData = await apiSearchSuggest(this.searchKeyword)
    },
    async updateSearchHot() {
      this.searchHotData = await apiSearchHotDetail()
    }
  }
})
