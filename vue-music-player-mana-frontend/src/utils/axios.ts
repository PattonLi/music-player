import axios from 'axios'
import router from '@/router'

// todo 设置baseurl
axios.defaults.baseURL = 'http://22.22.22.22:7777'
// 请求头，headers 信息
// axios.defaults.headers['X-Requested-With'] = 'XMLHttpRequest'

//token设置
axios.defaults.headers['token'] = localStorage.getItem('token') || ''
// 默认 post 请求，使用 application/json 形式
// axios.defaults.headers.post['Content-Type'] = 'application/json'

//设置interceptor
axios.interceptors.request.use((config)=>{
	config.headers.Authorization = 'Bearer '+localStorage.getItem('token');
	return config;
})

// 响应拦截器
axios.interceptors.response.use((res) => {
  if (typeof res.data !== 'object') {
    alert('服务端异常！')
    return Promise.reject(res)
  }
  if (res.data.resultCode != 200) {
    //认证超时
    if (res.data.resultCode == 419) {
      router.push({ path: '/login' })
    }
    return Promise.reject(res.data)
  }
  //no error
  return res.data.data
})

export default axios
