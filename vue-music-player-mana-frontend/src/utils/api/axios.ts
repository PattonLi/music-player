import axios from 'axios'

import { useAuthStore } from '@/stores/auth'

// todo 设置baseurl
// axios.defaults.baseURL = 'http://baidu.com'
// 请求头，headers 信息
// axios.defaults.headers['X-Requested-With'] = 'XMLHttpRequest'

// 默认 post 请求，使用 application/json 形式
// axios.defaults.headers.post['Content-Type'] = 'application/json'

//发送interceptor
//token设置
axios.interceptors.request.use((config) => {
  const authStore = useAuthStore()
  config.headers.Authorization = authStore.getAuthStatus.token
  return config
})

// 响应interceptor
axios.interceptors.response.use(
  (res) => {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    console.log(res)
    return res
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    console.log(error)
    return Promise.reject(error)
  }
)

export default axios
