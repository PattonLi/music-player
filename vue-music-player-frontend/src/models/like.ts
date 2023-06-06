export interface Like {
  likeId: number
  userId: number
  songId: number
  albumId: number
  artistId: number
  playlistId: number
  type: number //喜欢类型(1：歌曲 2：歌手 3：专辑 4:歌单)
}

export interface LikeForm {
  userId: number
  songId: number
  albumId: number
  artistId: number
  playlistId: number
  type: number //喜欢类型(1：歌曲 2：歌手 3：专辑 4:歌单)
}
