import axios from 'axios'
import myAxios from './myAxios'
import type { UserProfile, registerUser } from '@/models/user'

export async function apiLogin(phone: string, password: string) {
  return await myAxios.get<{
    code: number
    token: string
    userId: number
  }>('login/cellphone', { phone: phone, password: password })
}

export async function apiLoginStatus(id: number) {
  return await myAxios.get<{
    code: number
    profile: UserProfile
  }>('user/profile', { userId: id })
}

//去除repassword
interface ApiRegisterForm {
  username: string
  nickname: string
  phone: string
  email: string
  password: string
}

export async function apiRegister(ruleForm: registerUser) {
  return await myAxios.post<{
    code: number
    token: string
    userId: number
  }>('register', <ApiRegisterForm>{
    username: ruleForm.username,
    nickname: ruleForm.nickname,
    phone: ruleForm.phone,
    email: ruleForm.email,
    password: ruleForm.password
  })
}


export const apiUploadAction = async (file: File) => {
  return await upload(file)
}

const upload=(file: File)=>{
  const formData = new FormData()
  formData.append('file', file)
  return new Promise((resolve, reject) => {
    axios
      .post('user/profile/edit', JSON.stringify(formData))
      .then((res) => {
        resolve(res.data)
      })
      .catch((err) => {
        reject(err.data)
      })
  })
}