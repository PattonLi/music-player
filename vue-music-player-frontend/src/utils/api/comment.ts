import myAxios from "./myAxios"
import type{Comment} from '@/models/comment'

//获取歌曲所有评论
export const apiGetSongComment=async(songId:number)=>{
   return await myAxios.get<{
       code:number,
       comments:Comment[]
    }>('/song/comment/publish',{songId:songId})
}
//点赞评论
export const apiLikeComment=async(commentId:number,userId:number)=>{
    return await myAxios.post<{
        code:number
    }>('/song/comment/like',{commentId:commentId,userId:userId})
}
//发表评论
export const apiPublishComment=async(comment:string,userId:number,songId:number)=>{
    return await myAxios.post<{
        code:number
    }>('/song/comment/publish',{
        comment:comment,
        userId:userId,
        songId:songId
    })
}