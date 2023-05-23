import Mock from 'mockjs'

Mock.mock('/adminUser/login', 'post', {
  code: 200,
  token: 'aaaaaaaaaaaaaaaaaaaaaaaa' //token随便乱写
})

Mock.mock('/adminUser/profile', 'get', {
  code: 200,
  userInfo: {
    nickName: 'lpt',
    loginUserName: 'admin'
  }
})
