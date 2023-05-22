import axios from '@/utils/api/axios'
import type { SongInfo } from '@/model/SongInfo'

// 获取全部歌曲信息
const getSongInfo = async (currentPage: number, pageSize: number) => {
  try {
    const response = await axios.get('/adminUser/songInfo', {
      params: {
        currentPage: currentPage,
        pageSize: pageSize
      }
    })
    return response.data
  } catch (error) {
    console.log(error)
  }
}

// 添加歌曲
const addSong = async (newSong: SongInfo) => {
  try {
    const response = await axios.post('/admin/addSong', {
      params: {
        newSong: newSong
      }
    })
    return response.data
  } catch (error) {
    console.log(error)
  }
}

// 修改歌曲
const modiSong = async (modiSong: SongInfo) => {
  try {
    const response = await axios.get('/admin/modiSong', {
      params: {
        modiSong: modiSong
      }
    })
    return response.data
  } catch (error) {
    console.log(error)
  }
}

export { getSongInfo, addSong, modiSong }
