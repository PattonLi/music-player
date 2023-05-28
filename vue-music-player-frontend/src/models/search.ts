//热搜信息
export interface SearchHotDetail {
  searchWord: string //搜索词
  type: string //歌手、专辑、歌曲
}

//搜索推荐信息
export interface SearchSuggest {
  albums: SearchSuggestAlbums[]
  artists: SearchSuggestArtists[]
  songs: SearchSuggestSongs[]
  order: string[]
}

export interface SearchSuggestAlbums {
  id: number
  name: string
  artist: string
  publishTime: number
  songSize: number
  mark: number
}
export interface SearchSuggestArtists {
  id: number
  name: string
  picUrl: string
  albumSize: number
  img1: number
}

export interface SearchSuggestSongs {
  id: number
  name: string
  artist: string
  album: string
  duration: number
  mark: number
}
