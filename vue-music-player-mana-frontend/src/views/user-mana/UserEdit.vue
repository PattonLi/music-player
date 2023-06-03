<template>
  <div class="space">
    <el-row :gutter="20">
      <el-col :span="2">
        <span class="input-mention">用户名</span>
     </el-col>
      <el-col :span="5">
        <el-input
          v-model="input.input"
          placeholder="请输入内容"
        ></el-input>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" :icon="Search" @click="SearchClick">查询</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="primary" :icon="Upload" @click="AddClick">添加</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="info"  :icon="Check" @click="AllClick">全部</el-button>
      </el-col>
    </el-row>
  </div>

  <el-card class="main-card">
    <el-table :data="state.tableData" height="600" style="width: 100%">
      <el-table-column fixed prop="userId" label="ID" width="60" />
      <el-table-column fixed prop="" label="头像" width="120">
        <div class="demo-image">
          <el-image style="width: 100px; height: 100px" :src="'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'" />
        </div>
      </el-table-column>
      <el-table-column prop="username" label="姓名" width="85" />
      <el-table-column prop="nickname" label="昵称" width="105" />
      <el-table-column prop="gender" label="性别" width="70" />
      <el-table-column prop="age" label="年龄" width="70" />
      <el-table-column prop="email" label="邮箱" width="190" />
      <el-table-column prop="phone" label="电话" width="160" />
      <el-table-column prop="password" label="密码" width="190" />
      <el-table-column fixed="right" label="操作" width="200">
        <template #default="scope">
          <el-button type="warning" :icon="Edit" @click="EditClick(scope.row.user_id)">修改</el-button>
          <el-button type="danger" :icon="Delete" @click="DeleteClick(scope.row.user_id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!--分页按钮-->
    <el-pagination
      background
      layout="prev, pager, next"
      :total="state.totals"
      :page-size="state.pageSize"
      :current-page="state.currentPage"
      @current-change="changePage"
    />

  </el-card>

  <!--添加对话框-->
  <el-dialog
    v-model="addDialog.visable"
    title="添加"
    width="27%">

    <el-form :model="addDialog.data" label-width="100px">
      <el-form-item label="ID">
        <el-col :span="4">
          <el-input v-model="addDialog.data.userId" :readonly="true" />
        </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="3">
        <span>姓名</span>
      </el-col>
      <el-col :span="6">
        <el-input v-model="addDialog.data.username" />
      </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="3">
        <span>年龄</span>
      </el-col>
      <el-col :span="6">
        <el-input v-model="addDialog.data.age" />
      </el-col>
      
    </el-form-item>
      <el-form-item label="用户性别">
        <el-radio-group v-model="addDialog.data.gender">
          <el-radio label="男" />
          <el-radio label="女" />
        </el-radio-group>
    </el-form-item>
    <el-form-item label="用户昵称">
      <el-input v-model="addDialog.data.nickname"></el-input>
    </el-form-item>
    <el-form-item label="用户密码">
      <el-input v-model="addDialog.data.password"></el-input>
    </el-form-item>
    <el-form-item label="用户邮箱">
      <el-input v-model="addDialog.data.email"></el-input>
    </el-form-item>
    <el-form-item label="用户电话">
      <el-input v-model="addDialog.data.email"></el-input>
    </el-form-item>
    <el-form-item label="用户头像">
      <el-upload
        class="avatar-uploader"
        action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
        :show-file-list="false"
        :on-success="handleAvatarSuccess"
        :before-upload="beforeAvatarUpload"
      >
        <img v-if="imageUrl" :src="imageUrl" class="avatar" />
        <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
      </el-upload>
    </el-form-item>
  </el-form>


    <template #footer>
      <span class="dialog-footer">
        <el-button @click="addDialog.visable = false">Cancel</el-button>
        <el-button type="primary" @click="addDialog.visable = false">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>

  <!--修改对话框-->
  <el-dialog
    v-model="modifyDialog.visable"
    title="信息修改"
    width="27%">

    <el-form :model="modifyDialog.data" label-width="100px">
    <el-form-item label="ID">
      <el-col :span="4">
        <el-input v-model="modifyDialog.data.user_id" :readonly="true" />
      </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="3">
        <span>姓名</span>
      </el-col>
      <el-col :span="6">
        <el-input v-model="modifyDialog.data.username" />
      </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="3">
        <span>年龄</span>
      </el-col>
      <el-col :span="6">
        <el-input v-model="modifyDialog.data.age" />
      </el-col>
      
    </el-form-item>
    <el-form-item label="用户性别">
      <el-radio-group v-model="modifyDialog.data.gender">
        <el-radio label="男" />
        <el-radio label="女" />
      </el-radio-group>
    </el-form-item>
    <el-form-item label="更新时间">
      <el-col :span="11">
        <el-date-picker
          v-model="modifyDialog.data.updated_at"
          type="date"
          placeholder="Pick a date"
          style="width: 100%"
        />
      </el-col>
      <el-col :span="2" class="text-center">
        <span class="text-gray-500">-</span>
      </el-col>
      <el-col :span="11">
        <el-time-picker
          v-model="modifyDialog.data.updated_at"
          placeholder="Pick a time"
          style="width: 100%"
        />
      </el-col>
    </el-form-item>
    <el-form-item label="用户昵称">
      <el-input v-model="modifyDialog.data.nickname"></el-input>
    </el-form-item>
    <el-form-item label="用户密码">
      <el-input v-model="modifyDialog.data.password"></el-input>
    </el-form-item>
    <el-form-item label="用户邮箱">
      <el-input v-model="modifyDialog.data.email"></el-input>
    </el-form-item>

  </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="modifyDialog.visable = false">Cancel</el-button>
        <el-button type="primary" @click="confirmModify">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>

  <!--删除对话框-->
  <el-dialog
    v-model="deleteDialog.visable"
    title="系统提示"
    width="20%">

    <span>是否确认删除？</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="deleteDialog.visable = false">取消</el-button>
        <el-button type="primary" @click="deleteUserInfo">
          确认
        </el-button>
      </span>
    </template>
  </el-dialog>

</template>
  


<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { Check, Delete, Edit, Search, Share, Upload } from '@element-plus/icons-vue'
import { getCustomerInfo, getTheCustomerInfo, deleteTheCustomerInfo } from '@/utils/api/user'
import type{ CustomerInfo } from '@/model/UserInfo'
import type { UploadProps } from 'element-plus'
import { fill } from 'lodash';

const input = reactive({input: ""});  

const addDialog = reactive({
  visable: false,
  id: -1,
  data: {} as CustomerInfo
});

const deleteDialog = reactive({
  visable: false,
  id: -1
});

const modifyDialog = reactive({
  visable: false,
  id: -1,
  data: {} as CustomerInfo
})

const state = reactive({
  tableData: [] as CustomerInfo[], // 数据列表
  currentPage: 1,
  pageSize: 10,
  totals: 0, // 一共有多少页
})

//加载数据
onMounted(() => {
  getCustomerInfo(state.currentPage, state.pageSize).then((data) => {
    state.tableData = data.data
    state.totals = data.totals
    console.log(state.totals)
  })
})

const SearchClick = () => {
  console.log('click')
  getTheCustomerInfo(input.input).then((data) => {
    console.log(data.data)
    state.tableData = data.data
  })
}

const AddClick = () => {
  console.log('click')
  addDialog.visable = true
}

const EditClick = (id: number) => {
  console.log('click')
  modifyDialog.id = id
  const user = state.tableData.find(item => item.user_id === id);
  if(user) {
    modifyDialog.data = { ...user };
  }
  modifyDialog.visable = true
}

const DeleteClick = (id: number) => {
  console.log('click')
  deleteDialog.id = id
  deleteDialog.visable = true
}

const AllClick = () => {
  getCustomerInfo(state.currentPage, state.pageSize).then((data) => {
    state.tableData = data.data
    state.totals = data.totals
    console.log(state.totals)
  })
}

const deleteUserInfo = () => {
  deleteTheCustomerInfo(deleteDialog.id)
  const index = state.tableData.findIndex(item => item.user_id === deleteDialog.id);
  if (index !== -1) {
    state.tableData.splice(index, deleteDialog.id);
  }
  deleteDialog.id = -1
  deleteDialog.visable = false
}

const confirmModify = () => {

  modifyDialog.visable = false
}

const changePage = () => {

}

const handleAvatarSuccess: UploadProps['onSuccess'] = (
  response,
  uploadFile
) => {
  addDialog.data.picUrl = URL.createObjectURL(uploadFile.raw!)
}

const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
  if (rawFile.type !== 'image/jpeg') {
    ElMessage.error('Avatar picture must be JPG format!')
    return false
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('头像不能超过 2MB!')
    return false
  }
  return true
}

</script>

<style scoped>
.main-card{
  display: flex;
  border-radius: 10px;
  margin-bottom: 50px;
  width: auto;
  height: 660px;
}

.input-mention{
  font-size: 18px;
}

.space {
  margin-bottom: 20px;
}
</style>


