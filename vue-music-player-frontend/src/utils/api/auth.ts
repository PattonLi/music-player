import myAxios from './myAxios'
import type { UserProfile } from '@/models/user'

export async function apiLogin(phone: string, password: string) {
  return await myAxios.get<{
    code: number
    cookie: string
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
