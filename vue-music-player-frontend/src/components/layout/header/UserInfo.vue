<template>
  <!-- 头像 -->
  <div class="flex items-center cursor-pointer hover-text">
    <ElAvatar size="default" round class="bg-gray-200" :src="profile?.picUrl">
      <!-- 加载失败时展示 -->
      <img src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
    </ElAvatar>
    <span class="text-sm mx-2" v-if="isLogin">{{ profile.nickname }}</span>
    <span class="text-sm mx-2" @click="showLogin = true" v-else>点击登录</span>
  </div>

  <!-- 登录框 -->
  <AppLogin></AppLogin>
  <!-- 注册框 -->
  <AppResgister></AppResgister>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { storeToRefs } from 'pinia'
import AppLogin from '../auth/AppLogin.vue'
import AppResgister from '../auth/AppResgister.vue'

const { checkLogin } = useAuthStore()
const { isLogin, profile, showLogin } = storeToRefs(useAuthStore())

onMounted(() => {
  checkLogin()
})
</script>

<style lang="scss"></style>
