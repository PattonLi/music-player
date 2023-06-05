<template>
  <el-card class="main-card">
    <el-row class="el-row">
      <el-col :span="4" class="el-col">
        <el-card class="user-card" shadow="hover" :body-style="{ padding: '0px' }">
          <img class="image" :src="state.userData.picUrl" :fit="'cover'"/>
          <div style="padding: 14px">
            <span style="font-size: 35px; font-weight: bold">{{ state.userData.nickname }}</span>
            <span style="font-size: 20px; font-weight: bold">的收藏</span>
          </div>
        </el-card>
      </el-col>
      <!--歌曲-->
      <el-col :span="4" class="el-col-3">
        <el-card class="data-card" :body-style="{ padding: '0px' }">
          <template #header>
            <span style="font-size: 30px; font-weight: bold">收藏的歌</span>
          </template>

        </el-card>
      </el-col>
      <!--歌手-->
      <el-col :span="4" class="el-col-3">
        <el-card class="data-card" :body-style="{ padding: '0px' }">
          <template #header>
            <span style="font-size: 30px; font-weight: bold">收藏的歌手</span>
          </template>
          
        </el-card>
      </el-col>
      <!--专辑-->
      <el-col :span="4" class="el-col-3">
        <el-card class="data-card" :body-style="{ padding: '0px' }">
          <template #header>
            <span style="font-size: 30px; font-weight: bold">收藏的专辑</span>
          </template>
          
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

const test = reactive({
  q: [
    {t: 1},
    {t: 1},
    {t: 1},
    {t: 1},
    {t: 1},
    {t: 1},
    {t: 1},
    {t: 1},
    {t: 1},
    {t: 1}
  ]
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

.el-col {
  margin-left: 20px;
}
.el-col-3 {
  margin-left: 100px;
}
.user-card{
  margin-top: 10px;
  margin-bottom: 10px;
  display: flex;
  width: 300px;
  height: 375px;
}
.data-card{
  margin-top: 10px;
  margin-bottom: 10px;
  margin-left: 10px;
  display: flex;
  width: 290px;
  height: 600px;
}
.image {
  width: 300px;
  height: 300px;
  display: block;
}

</style>
