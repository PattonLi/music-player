<template>
  <div class="header">
    <div class="left">
      <span style="font-size: 30px">{{ state.name }}</span>
    </div>
    <div class="right">右</div>
  </div>
</template>

<script setup lang="ts">
import { pathMap } from '@/utils'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const state = ref({
  name: '首页'
})

// watch router change
router.afterEach((to) => {
  console.log('to', to)
  //to.name is the router name
  state.value.name = pathMap[to.name as keyof typeof pathMap]
})
</script>
<style scoped>
.header {
  height: 50px;
  border-bottom: 1px solid #e9e9e9;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
}
</style>
