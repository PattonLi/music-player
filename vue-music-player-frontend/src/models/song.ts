export interface Song {
  id: number //歌曲ID
  albumId: number //专辑id
  name: string //歌曲名称
  pop: number //歌曲流行指数
  album: string //歌曲专辑名
  mark: number //评论数
  publishTime: number //发行时间
  type: string //歌曲类别

  picUrl?: any //歌曲封面url
  url?: any //歌曲url
  size?: number //歌曲文件大小
  duration:number//歌曲秒数时长
}
