import Mock from 'mockjs'

Mock.setup({
  timeout: 2000
})

// 当用户ajax地址里面 有/api/login 的时候 返回如下数据
Mock.mock('/adminUser/login', 'post', {
  code: 200,
  msg: '登录成功',
  token: 'aaaaaaaaaaaaaaaaaaaaaaaa' //token随便乱写
})
