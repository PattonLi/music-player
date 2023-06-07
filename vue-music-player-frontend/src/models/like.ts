export interface Like {
  likeId: number
  userId: number
  songId: number
  albumId: number
  artistId: number
  type: number //喜欢类型(1：歌曲 2：歌手 3：专辑)
}

export interface LikeForm {
  userId: number
  songId: number
  albumId: number
  artistId: number
  type: number //喜欢类型(1：歌曲 2：歌手 3：专辑)
}
