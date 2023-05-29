import { defineStore } from 'pinia'
import type { UserProfile, registerUser } from '@/models/user'
import { apiLogin, apiLoginStatus, apiRegister } from '@/utils/api/auth'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isLogin: false,
    token: '',

    showRegister: false,
    showLogin: false,
    //用户信息
    profile: {} as UserProfile,
    //登录表单
    ruleForm: {} as registerUser
  }),
  getters: {},
  actions: {
    async login(phone: string, password: string) {
      const res = await apiLogin(phone, password)
      if (res.code == 200) {
        this.token = res.token
        this.checkLogin()
      }
    },
    async register(ruleForm: registerUser) {
      this.showRegister = false
      const res = await apiRegister(ruleForm)
      if (res.code == 200) {
        this.token = res.token
        this.checkLogin()
      }
    },

    //更新一下个人信息
    async checkLogin() {
      const { data } = await apiLoginStatus()
      if (data.code === 200) {
        this.isLogin = true
        this.profile = data.profile
      }
    }
  },
  //持久化
  persist: true
})
