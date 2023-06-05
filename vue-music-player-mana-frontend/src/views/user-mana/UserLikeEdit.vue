<template>
  <el-card class="main-card">
    <el-row class="el-row">
      <el-col :span="12">
        <el-card class="user-card" shadow="hover" :body-style="{ padding: '0px' }">
        <img class="image" :src="state.userData.picUrl" :fit="'cover'"/>
        <div style="padding: 14px">
          <span style="font-size: 35px; font-weight: bold">{{ state.userData.nickname }}</span>
          <span style="font-size: 20px; font-weight: bold">的收藏</span>
        </div>
      </el-card>
      </el-col>
    </el-row>
  </el-card>
  <!--分页按钮-->
  <el-pagination
    v-model:current-page="state.currentPage"
    v-model:page-size="state.pageSize"
    layout="total, prev, pager, next, jumper"
    :total="state.totals"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  />


</template>

<script setup lang="ts">
import type{ CustomerInfo } from '@/model/UserInfo'
import { getACustomerInfo } from '@/utils/api/user'
import type { size } from 'lodash';

const state = reactive({
  userData: {} as CustomerInfo, // 用户信息

  currentPage: 1,
  pageSize: 1,
  totals: 400, // 一共有多少条目
})

//加载数据
onMounted(() => {
  getACustomerInfo(state.currentPage).then((data) => {
    if(data.code === 200) {
      state.userData = data.data
      state.totals = data.totals
      console.log(state.totals)
    }
  })
})
</script>

<style scoped>
.main-card{
  display: flex;
  border-radius: 10px;
  margin-bottom: 50px;
  width: auto;
  height: 660px;
}
.el-row {
  margin-bottom: 40px;
}
.user-card{
  margin-top: 10px;
  margin-bottom: 10px;
  display: flex;
  width: 300px;
  height: 375px;
}
.image {
  width: 300px;
  height: 300px;
  display: block;
}
.test {
  display: flex;
  border-radius: 10px;
  margin-bottom: 50px;
  width: 250px;
  height: 300px;
}
.text{
    size: 200px;
}
</style>
