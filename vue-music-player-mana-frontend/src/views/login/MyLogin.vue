<template>
  <div class="login-body">
    <!--登录框div-->
    <div class="login-container">
      <!--登录框头部logo部分-->
      <div class="head">
        <img class="logo" src="https://s.weituibao.com/1582958061265/mlogo.png" />
        <div class="name">
          <div class="title">sysu 网页音乐</div>
          <div class="tips">Vue3.2 后台管理系统</div>
        </div>
      </div>
      <el-form
        label-position="top"
        :rules="rules"
        :model="state.ruleForm"
        ref="loginForm"
        class="login-form"
      >
        <el-form-item label="账号" prop="username">
          <el-input
            type="text"
            v-model.trim="state.ruleForm.username"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            type="password"
            v-model.trim="state.ruleForm.password"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <div style="color: #333">登录表示您已同意<a>《服务条款》</a></div>
          <el-button style="width: 100%" type="primary" @click="submitForm">立即登录</el-button>
          <el-checkbox v-model="state.checked" @change="!state.checked">下次自动登录</el-checkbox>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import axios from 'axios'
import { Md5 } from 'ts-md5'
import { useAuthStore } from '@/stores/auth'
import type { FormRules } from 'element-plus'

const authStore = useAuthStore()
const loginForm = ref()
const rules = reactive<FormRules>({
  username: [{ required: true, message: '账户不能为空', trigger: 'blur' }],
  password: [{ required: true, message: '密码不能为空', trigger: 'blur' }]
})
const state = reactive({
  ruleForm: {
    username: '',
    password: ''
  },
  checked: true
  // 表单验证判断。
})

// 表单提交方法,回调函数
const submitForm = async () => {
  loginForm.value.validate((valid: boolean) => {
    //表示表单是否通过了上面 rules 的规则。
    if (valid) {
      axios
        .post('/adminUser/login', {
          userName: state.ruleForm.username || '',
          passwordMd5: Md5.hashStr(state.ruleForm.password)
        })
        .then((res) => {
          // 返回一个 token
          authStore.SetToken(res.data['token'])
          // 刷新页面
          window.location.href = '/'
        })
    } else {
      console.log('登录表单不符合规则！')
      return false
    }
  })
}
</script>

<style scoped>
.login-body {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100vh;
  background-color: #fff;
}
.login-container {
  width: 420px;
  height: 500px;
  background-color: #fff;
  border-radius: 4px;
  box-shadow: 0px 21px 41px 0px rgba(0, 0, 0, 0.2);
}
.head {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px 0 20px 0;
}
.head img {
  width: 100px;
  height: 100px;
  margin-right: 20px;
}
.head .title {
  font-size: 28px;
  color: #1baeae;
  font-weight: bold;
}
.head .tips {
  font-size: 12px;
  color: #999;
}

.login-form {
  width: 70%;
  margin: 0 auto;
}
</style>
