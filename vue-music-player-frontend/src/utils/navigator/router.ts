/* router根据路由名称和id跳转 */
import router from '@/router' // 导入Vue Router实例
import { NavigationFailureType, isNavigationFailure } from 'vue-router'
import { AlertRouterInfo } from '../alert/AlertPop'

//带参数跳转
export const routerPushByNameId = (name: string, id: number) => {
  router.push({ name: name, query: { id: id } }).then((failure) => {
    if (isNavigationFailure(failure, NavigationFailureType.aborted)) {
      AlertRouterInfo("you don't have auth to this path:" + failure.to.path)
      //failure.to.path
      //failure.from.path
    }
  })
}
