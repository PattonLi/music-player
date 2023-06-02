//热搜信息
export interface SearchHotDetail {
  searchWord: string //搜索词
  type: string //歌手、专辑、歌曲
}

export interface SearchResults{
  albums: SearchAlbums[]
  artists: SearchArtists[]
  songs: SearchSongs[]
  users:SearchUsers[]
}

//搜索推荐信息
export interface SearchAlbums {
  albumId: number
  album: string
  picUrl:string
  artist: string
  publishTime: number
  Size: number
}
export interface SearchArtists {
  artistId: number
  artist: string
  picUrl: string
  songSize:number
  albumSize:number
}

export interface SearchSongs {
  songId: number
  name: string
  artist: string
  album: string
  duration: number
}

export interface SearchUsers{
  nickname:string
  userId:number
  picUrl:string
}