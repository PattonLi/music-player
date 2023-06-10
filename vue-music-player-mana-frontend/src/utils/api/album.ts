import axios from '@/utils/api/axios'
import type { AlbumInfo } from '@/model/AlbumInfo'

// 获取全部歌曲信息
const getAlbumInfo = async (currentPage: number, pageSize: number) => {
  try {
    const response = await axios.get('/admin/pageAllAlbum', {
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

// 获取歌手全部专辑信息
const getArtistAllAlbum = async (artistId: number) => {
  try {
    const response = await axios.get('/admin/getAlbumByArtist', {
      params: {
        artistId: artistId
      }
    })
    return response.data
  } catch (error) {
    console.log(error)
  }
}

// 获得特定专辑
const getTheAlbumInfo = async (albumname: string) => {
  try {
    const response = await axios.get('/admin/theAlbum', {
      params: {
        albumname: albumname
      }
    })
    console.log(response)
    return response.data
  } catch (error) {
    console.log(error)
  }
}

  // 删除特定专辑
  const deleteTheAlbumInfo = async (albumId: number) => {
    try {
      const response = await axios.get('/admin/deleteAlbum', {
        params: {
          albumId: albumId
        }
      })
      console.log(response)
      return response.data
    } catch (error) {
      console.log(error)
    }
  }

    // 添加特定专辑
    const addAlbumInfo = async (data: AlbumInfo) => {
      try {
        const jsonData = JSON.stringify(data); // 将对象转换为 JSON 字符串
        const response = await axios.post('/admin/addAlbum', jsonData);
        // 处理响应
        console.log(response)
        return response.data
      } catch (error) {
        // 处理错误
        console.log(error)
      }
    }

// 修改特定专辑
const modifyAlbumInfo = async (modiArtist: AlbumInfo) => {
  try {
    const jsonData = JSON.stringify(modiArtist); // 将对象转换为 JSON 字符串
    const response = await axios.post('/admin/modifyAlbum', jsonData);
    // 处理响应
    console.log(response)
    return response.data
  } catch (error) {
    // 处理错误
    console.log(error)
  }
}

export { getAlbumInfo, getArtistAllAlbum, getTheAlbumInfo, deleteTheAlbumInfo, addAlbumInfo, modifyAlbumInfo }
