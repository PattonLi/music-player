<script setup lang="ts">
import MyHeader from '@/views/layout/MyHeader.vue'
import MyFotter from '@/views/layout/MyFotter.vue'
import MyMenu from '@/views/layout/MyMenu.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { localGet ,pathMap} from '@/utils/index'

//登录相关逻辑
const state = ref({
  showMenu: true // 是否需要显示菜单
})
const noMenu = ['/login']
const router = useRouter()
// 监听路由的变化
//若未登录则跳转登录
router.beforeEach((to, from, next) => {
  if (to.path == '/login') {
    // 如果路径是 /login 则正常执行
    next()
  } else {
    //判断是否有token
    if (localGet('token') !== null) {
      // 如果没有，则跳至登录页面
      console.log('token',localGet('token'));
      next()
    } else {
      console.log('token',localGet('token'));
      next({ path: '/login' })
    }
  }
  //如果未
  state.value.showMenu = !noMenu.includes(to.path)
  document.title = pathMap[to.name as keyof typeof pathMap]
})
</script>

<template>
  <div id="app">
    <div class="layout">
      <el-container v-if="state.showMenu" class="container">
        <MyMenu></MyMenu>
        <!--main container-->
        <el-container class="content">
          <el-header>
            <MyHeader></MyHeader>
          </el-header>
          <el-main class="main">
            <RouterView></RouterView>
          </el-main>
          <el-footer>
            <MyFotter></MyFotter>
          </el-footer>
        </el-container>
      </el-container>
      <el-container v-else>
        <RouterView></RouterView>
      </el-container>
    </div>
  </div>
</template>

<style scoped>
.content {
  display: flex;
  flex-direction: column;
  max-height: 100vh;
  overflow: hidden;
}

.main {
  height: 100vh;
  overflow: auto;
  padding: 10;
}
</style>
