import { defineStore } from 'pinia'
import type { UserProfile } from '@/models/user'
import { apiLogin, apiLoginStatus } from '@/utils/api/auth'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isLogin: false,
    token: '',
    showLogin: false,
    profile: {} as UserProfile
  }),
  getters: {
    isLogin: (state) => {
      return state.profile?.userId > 0
    }
  },
  actions: {
    async login(phone: string, password: string) {
      const res = await apiLogin(phone, password)
      if (res.code == 200) {
        this.token = res.token
        this.checkLogin()
      }
    },
    async checkLogin() {
      const { data } = await apiLoginStatus()
      if (data.code === 200) {
        this.profile = data.profile
        this.showLogin = false
      }
    }
  },
  //持久化
  persist: true
})
