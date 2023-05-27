import myAxios from './myAxios'
import type { UserProfile, registerUser } from '@/models/user'

export async function apiLogin(phone: string, password: string) {
  return await myAxios.get<{
    code: number
    token: string
  }>('login/cellphone', { phone: phone, password: password })
}

export async function apiLoginStatus() {
  return await myAxios.get<{
    data: {
      code: number
      profile: UserProfile
    }
  }>('login/status')
}

export async function apiRegister(ruleForm: registerUser) {
  return await myAxios.get<{
    code: number
    token: string
  }>('login/cellphone', ruleForm)
}
