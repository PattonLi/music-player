import { defineStore } from 'pinia'
import { apiSearchSuggest } from '@/utils/api/search'
import type { SearchSuggest } from '@/models/search'

//搜索结果状态
export const useSearchStore = defineStore('search', {
  state: () => {
    return {
      showSearchView: false,
      searchKeyword: '',
      suggestData: {} as SearchSuggest
    }
  },
  getters: {
    showHot: (state) => {
      return state.searchKeyword == ''
    }
  },
  actions: {
    async suggest() {
      this.suggestData = await apiSearchSuggest(this.searchKeyword)
    }
  }
})

export interface SearchHotDetail {
	searchWord: string;
	score: number;
	content: string;
	source: number;
	iconType: number;
	iconUrl?: string;
	url: string;
	alg: string;
}