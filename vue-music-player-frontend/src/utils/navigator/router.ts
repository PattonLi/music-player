/* router根据路由名称和id跳转 */
import router from '@/router' // 导入Vue Router实例
import { Pages } from '@/router/pages'
import { NavigationFailureType, isNavigationFailure } from 'vue-router'

//带参数跳转
export const routerPushByNameId = (name: string, id: number) => {
  // if(name==Pages.albumDetail||name==Pages.artistDetail||name==Pages.playlistDetail){
  // router.push({ name: name, query: { id: id } ,replace:true})
  // }else{
  router.push({ name: name, query: { id: id } })
  // }
}

//跳转搜索
export const routerPushByNameKeyWord = (name: string, keyWord: string) => {
  router.push({ name: name, query: { keyWord: keyWord }, replace: true }).then((failure) => {
    if (isNavigationFailure(failure, NavigationFailureType.aborted)) {
      //failure.to.path
      //failure.from.path
    }
  })
}

export const routerPush = (name: string) => {
  router.push({ name: name }).then((failure) => {
    if (isNavigationFailure(failure, NavigationFailureType.aborted)) {
      //failure.to.path
      //failure.from.path
    }
  })
}
