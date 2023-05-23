import Mock from 'mockjs'

Mock.setup({
  timeout: 1000
})

// 当用户ajax地址里面 有/api/login 的时候 返回如下数据
Mock.mock('/adminUser/login', 'post', {
  code: 200,
  msg: '登录成功',
  token: 'aaaaaaaaaaaaaaaaaaaaaaaa' //token随便乱写
})

Mock.mock('/adminUser/profile', 'get', {
  code: 200,
  userInfo: {
    nickNmae: 'lissss',
    loginUserName: 'admin'
  }
})
