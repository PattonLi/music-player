import axios from '@/utils/api/axios'
import type { DataAnalyse } from '@/model/DataAnalyse'

// 获取用户信息
const getDataAnalyse = async () => {
  let dataAnalyse: DataAnalyse = {
    numOfDownloadSong: [],
    numOfLoginUsers: [],
    numOfPlaySong: [],
    numOfRegisterUsers: []
  }
  try {
    const response = await axios.get('/adminUser/dataAnalyse', {})
    dataAnalyse = response.data.dataAnalyse
  } catch (error) {
    console.log(error)
  }
  return dataAnalyse
}

export { getDataAnalyse }
