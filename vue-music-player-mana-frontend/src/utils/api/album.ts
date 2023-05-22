import axios from '@/utils/api/axios'

// 获取全部歌曲信息
const getAlbumInfo = async (currentPage: number, pageSize: number) => {
  try {
    const response = await axios.get('/adminUser/albumInfo', {
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


export { getAlbumInfo}
