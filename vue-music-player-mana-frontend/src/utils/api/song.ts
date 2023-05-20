import axios from '@/utils/api/axios'

// 获取用户信息
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

export { getSongInfo }
