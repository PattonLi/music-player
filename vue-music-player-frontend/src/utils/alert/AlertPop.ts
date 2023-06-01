import { ElMessage, ElNotification } from 'element-plus'

const AlertSuccess = (msg: string) => {
  ElMessage({
    message: msg,
    type: 'success'
  })
}

const AlertError = (msg: string) => {
  ElMessage.error(msg)
}

const AlertAxiosError = (msg: string) => {
  ElNotification({
    title: 'ApiError',
    message: msg,
    type: 'error',
    customClass: 'dark:bg-zinc-800',
    offset: 50
  })
}

export { AlertError, AlertSuccess, AlertAxiosError }
