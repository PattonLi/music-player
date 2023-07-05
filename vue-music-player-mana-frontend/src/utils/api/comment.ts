import axios from '@/utils/api/axios'
import type { ArtistInfo } from '@/model/ArtistInfo'

// 获取特定歌手信息
const getTheSongInfo = async (name: string) => {
  try {
    const response = await axios.get('/admin/getArtist', {
      params: {
        name: name
      }
    })
    console.log(response)
    return response.data
  } catch (error) {
    console.log(error)
  }
}

// 获取特定歌曲评论信息
const getTheCommentInfo = async (songId: number) => {
    try {
      const response = await axios.get('/comment/getComment', {
        params: {
          songId: songId
        }
      })
      console.log(response)
      return response.data
    } catch (error) {
      console.log(error)
    }
  }
// 删除特定评论
const deleteTheCommentInfo = async (commentId: number) => {
    try {
        const response = await axios.get('/comment/deleteComment', {
          params: {
            commentId: commentId
          }
        })
        console.log(response)
        return response.data
      } catch (error) {
        console.log(error)
      }
}
export {
    getTheCommentInfo, deleteTheCommentInfo
}
