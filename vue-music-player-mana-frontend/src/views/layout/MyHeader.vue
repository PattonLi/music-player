<template>
  <div class="header">
    <div class="left">
      <span style="font-size: 30px">{{ state.name }}</span>
    </div>
    <div class="right">
      <el-popover placement="bottom" :width="300" trigger="click" popper-class="popper-user-box">
        <template #reference>
          <div class="author">
            <i class="icon el-icon-custom" />
            {{ (state.userInfo && state.userInfo.nickName) || '' }}
            <i class="el-icon-caret-bottom" />
          </div>
        </template>
        <div class="nickname">
          <p>登录名：{{ (state.userInfo && state.userInfo.loginUserName) || '' }}</p>
          <p>昵称：{{ (state.userInfo && state.userInfo.nickName) || '' }}</p>
          <el-tag size="small" effect="dark" class="logout" @click="logout">退出</el-tag>
        </div>
      </el-popover>
    </div>
  </div>
</template>

<script setup lang="ts">
import { pathMap } from '@/router/pathMap'
import { useAuthStore } from '@/stores/auth'
import axios from '@/utils/axios'
import { onMounted, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

const authStore = useAuthStore()
const router = useRouter()
const state = reactive({
  name: 'dashboard',
  userInfo: {
    nickName: null,
    loginUserName: null
  } // 用户信息变量
})

// onMounted方法
onMounted(() => {
  const pathname = window.location.hash.split('/')[1] || ''
  if (!['login'].includes(pathname)) {
    getUserInfo()
  }
})
// 获取用户信息
const getUserInfo = async () => {
  const userInfo = await (await axios.get('/adminUser/profile')).data.userInfo
  state.userInfo = userInfo
}
// 退出登录
const logout = () => {
  axios.delete('/adminUser/logout').then(() => {
    // 退出之后，将本地保存的 token  清理掉
    authStore.removeToken()
    // 回到登录页
    router.push({ path: '/login' })
  })
}
// watch router change
router.afterEach((to) => {
  console.log('to', to)
  //to.name is the router name
  state.name = pathMap[to.name as keyof typeof pathMap]
})
</script>
<style scoped>
.header {
  height: 50px;
  border-bottom: 1px solid #e9e9e9;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
}

.right > div > .icon {
  font-size: 18px;
  margin-right: 6px;
}
.author {
  margin-left: 10px;
  cursor: pointer;
}
</style>
<style>
.popper-user-box {
  background: url('https://s.yezgea02.com/lingling-h5/static/account-banner-bg.png') 50% 50%
    no-repeat !important;
  background-size: cover !important;
  border-radius: 0 !important;
}
.popper-user-box .nickname {
  position: relative;
  color: #ffffff;
}
.popper-user-box .nickname .logout {
  position: absolute;
  right: 0;
  top: 0;
  cursor: pointer;
}
</style>
