import axios from '@/utils/api/axios'
import type { UserInfo } from '@/model/UserInfo'

// 获取用户信息
const getUserInfo = async () => {
  let userInfo: UserInfo = { loginUserName: '', nickName: '' }
  try {
    const response = await axios.get('/adminUser/profile', {})
    userInfo = response.data.userInfo
  } catch (error) {
    console.log(error)
  }
  return userInfo
}

export { getUserInfo }
