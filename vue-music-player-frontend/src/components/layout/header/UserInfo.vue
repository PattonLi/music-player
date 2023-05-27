<template>
  <div>
    <el-switch
    v-model="isDarkMode"
    class="ml-2"
    style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
    />
  </div>
  <div class="flex items-center cursor-pointer hover-text">
    <ElAvatar size="small" round class="bg-gray-200" :src="profile?.avatarUrl ?? ''"></ElAvatar>
    <span class="text-xs ml-1.5" v-if="isLogin">{{ profile.nickname }}</span>
    <span class="text-xs ml-1.5" @click="showLogin = true" v-else>点击登录</span>
  </div>
  <el-dialog v-model="showLogin" title="登录" width="330px" append-to-body>
    <div>
      <el-input size="large" placeholder="手机号码" :prefix-icon="Phone" v-model="phone" />
      <el-input
        size="large"
        class="mt-5"
        placeholder="登录密码"
        :prefix-icon="Lock"
        v-model="password"
      />

      <button @click="loginSubmit" class="button w-full mt-5 py-5" style="border-radius: 5px">
        登录
      </button>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { Lock, Phone } from '@icon-park/vue-next'
import { useAuthStore } from '@/stores/auth'
import { storeToRefs } from 'pinia'

//夜间模式
const isDarkMode = ref(false)
watch(isDarkMode, (newValue, oldValue) => {
  if(newValue){
    document.documentElement.classList.add('dark')
  }else{
    document.documentElement.classList.remove('dark')
  }
})

const phone = ref('')
const password = ref('')
const { login, checkLogin } = useAuthStore()
const { isLogin, profile, showLogin } = storeToRefs(useAuthStore())

const loginSubmit = () => {
  login(phone.value, password.value)
}

onMounted(() => {
  checkLogin()
})
</script>

<style lang="scss"></style>
