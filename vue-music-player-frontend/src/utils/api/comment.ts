import myAxios from './myAxios'
import type { Comment } from '@/models/comment'

//获取歌曲所有评论
export const apiGetSongComment = async (songId: number) => {
  return await myAxios.get<{
    code: number
    comments: Comment[]
  }>('/song/comment', { songId: songId })
}
//点赞评论
export const apiLikeComment = async (commentId: number, userId: number) => {
  return await myAxios.post<{
    code: number
  }>('/song/comment/like', { commentId: commentId, userId: userId })
}
//取消点赞评论
export const apiDelCommentLike = async (commentId: number, userId: number) => {
  return await myAxios.post<{
    code: number
  }>('/song/comment/unlike', { commentId: commentId, userId: userId })
}
//发表评论
export const apiPublishComment = async (comment: string, userId: number, songId: number) => {
  return await myAxios.post<{
    code: number
  }>('/song/comment/publish', {
    comment: comment,
    userId: userId,
    songId: songId
  })
}
//获取已经点赞评论
export const apiGetCommentLikes = async (userId: number) => {
  return await myAxios.post<{
    code: number
    commentLikes:number[]
  }>('/user/comment/like', {userId: userId })
}