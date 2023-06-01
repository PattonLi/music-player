<template>
  <div class="space">
    <el-row :gutter="20">
      <el-col :span="2">
        <span class="input-mention">用户名</span>
     </el-col>
      <el-col :span="5">
        <el-input
          v-model="input"
          placeholder="请输入内容"
        ></el-input>
      </el-col>
      <el-col :span="2">
        <el-button type="success" :icon="Search" @click="SearchClick">查询</el-button>
      </el-col>
      <el-col :span="2">
        <el-button type="primary" :icon="Upload" @click="AddClick">添加</el-button>
      </el-col>
    </el-row>
  </div>

  <el-card class="main-card">
    <el-table :data="state.tableData" style="width: 100%">
      <el-table-column fixed prop="date" label="注册日期" width="140" />
      <el-table-column prop="name" label="姓名" width="100" />
      <el-table-column prop="nickname" label="昵称" width="120" />
      <el-table-column prop="gender" label="性别" width="100" />
      <el-table-column prop="age" label="年龄" width="100" />
      <el-table-column prop="email" label="邮箱" width="220" />
      <el-table-column prop="password" label="密码" width="220" />
      <el-table-column fixed="right" label="操作" width="220">
        <template #default="scoped">
          <el-button type="warning" :icon="Edit" @click="EditClick">修改</el-button>
          <el-button type="danger" :icon="Delete" @click="DeleteClick">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

  </el-card>
</template>
  


<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { Delete, Edit, Search, Share, Upload } from '@element-plus/icons-vue'
import { getCustomerInfo } from '@/utils/api/user'
import type{ CustomerInfo } from '@/model/UserInfo'

const input = ref('');  
// 测试数据
const state = reactive({
  tableData: [] as CustomerInfo[], // 数据列表
  currentPage: 1, // 当前页数
  pageSize: 10, // 每页请求数
  type: 'add', // 操作类型
  totals: 0, //所有项的总数
  multipleSelection: [] // 选中项
})
//加载数据

onMounted(() => {
  getCustomerInfo().then((data) => {
    state.tableData = data.data
    console.log(state.tableData)
  })
})


const SearchClick = () => {
  console.log('click')
}

const AddClick = () => {
  console.log('click')
}

const EditClick = () => {
  console.log('click')
}

const DeleteClick = () => {
  console.log('click')
}

</script>

<style scoped>
.main-card{
  display: flex;
  align-items: center;
  border-radius: 10px;
  margin-bottom: 50px;
}

.input-mention{
  font-size: 18px;
}

.space {
  margin-bottom: 20px;
}
</style>


