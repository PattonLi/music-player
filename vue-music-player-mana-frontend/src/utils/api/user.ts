import axios from '@/utils/api/axios'
import type { UserInfo, CustomerInfo } from '@/model/UserInfo'

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


const getCustomerInfo = async () => {
  let customerInfoArray = ref<CustomerInfo[]>([]);
  try {
    const response = await axios.get('/User/info', {}); // 替换成你的 API 接口地址
    console.log(response.data)
    return response.data;
  } catch (error) {
    console.error('Failed to fetch customer info:', error);
  }
}


export { getUserInfo, getCustomerInfo }
