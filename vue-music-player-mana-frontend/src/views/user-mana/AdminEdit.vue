<template>
  <div class="space">
    <el-row :gutter="20">
      <el-col :span="2">
        <span class="input-mention">管理员名称</span>
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
      <el-table-column fixed prop="adminId" label="ID" width="100" />
      <el-table-column prop="adminName" label="名称" width="100" />
      <el-table-column prop="password" label="密码" width="300" />
      <el-table-column fixed="right" label="操作" width="200">
        <template #default="scope">
          <el-button type="warning" :icon="Edit" @click="EditClick(scope.row.adminId)">修改</el-button>
          <el-button type="danger" :icon="Delete" @click="DeleteClick(scope.row.adminId)">删除</el-button>
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
          <el-input v-model="addDialog.data.adminId" :readonly="true" />
        </el-col>
      </el-form-item>
      <el-form-item label="管理员名称">
        <el-input v-model="addDialog.data.adminname"></el-input>
      </el-form-item>
      <el-form-item label="管理员密码">
        <el-input v-model="addDialog.data.password"></el-input>
      </el-form-item>
    </el-form>


    <template #footer>
      <span class="dialog-footer">
        <el-button @click="addDialog.visable = false">Cancel</el-button>
        <el-button type="primary" @click="confirmAdd">
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
        <el-input v-model="modifyDialog.data.adminId" :readonly="true" />
      </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="3">
        <span>姓名</span>
      </el-col>
      <el-col :span="6">
        <el-input v-model="modifyDialog.data.adminname" />
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
    <el-form-item label="用户昵称">
      <el-input v-model="modifyDialog.data.nickname"></el-input>
    </el-form-item>
    <el-form-item label="用户密码">
      <el-input v-model="modifyDialog.data.password"></el-input>
    </el-form-item>
    <el-form-item label="用户邮箱">
      <el-input v-model="modifyDialog.data.email"></el-input>
    </el-form-item>
    <el-form-item label="用户电话">
      <el-input v-model="modifyDialog.data.phone"></el-input>
    </el-form-item>
    <el-form-item label="用户头像">
      <el-upload
        class="avatar-uploader"
        action="#"
        :show-file-list="false"
        list-type="picture-card"
        :limit="1"
        :on-success="mhandleAvatarSuccess"
        :before-upload="mbeforeAvatarUpload"
      >
        <img style="width: 147px; height: 147px" v-if="modifyDialog.data.picUrl" :src="modifyDialog.data.picUrl" class="avatar" />
        <el-icon v-else><Plus /></el-icon>
      </el-upload>
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
        <el-button type="primary" @click="deleteAdminInfo">
          确认
        </el-button>
      </span>
    </template>
  </el-dialog>

</template>
  


<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { Check, Delete, Edit, Search, Share, Upload } from '@element-plus/icons-vue'
import { getAdminInfo, getTheAdminInfo, deleteTheAdminInfo, modifyTheAdminInfo, addTheAdminInfo } from '@/utils/api/admin'
import type{ AdminInfo } from '@/model/UserInfo'
import type { UploadProps } from 'element-plus'
import { add, fill } from 'lodash';
import axios from 'axios'

const input = reactive({input: ""});  

const addDialog = reactive({
  visable: false,
  id: -1,
  data: {} as AdminInfo,
});

const deleteDialog = reactive({
  visable: false,
  id: -1
});

const modifyDialog = reactive({
  visable: false,
  id: -1,
  data: {} as AdminInfo
})

const state = reactive({
  tableData: [] as AdminInfo[], // 数据列表
  currentPage: 1,
  pageSize: 10,
  totals: 0, // 一共有多少页
})

//加载数据
onMounted(() => {
  getAdminInfo(state.currentPage, state.pageSize).then((data) => {
    state.tableData = data.data
    state.totals = data.totals
    console.log(state.totals)
  })
})

const SearchClick = () => {
  console.log('click')
  getTheAdminInfo(input.input).then((data) => {
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
  const admin = state.tableData.find(item => item.adminId === id);
  if(admin) {
    modifyDialog.data = { ...admin };
  }
  modifyDialog.visable = true
  console.log(modifyDialog.id)
}

const DeleteClick = (id: number) => {
  console.log('click')
  deleteDialog.id = id
  deleteDialog.visable = true
}

const AllClick = () => {
  getAdminInfo(state.currentPage, state.pageSize).then((data) => {
    state.tableData = data.data
    state.totals = data.totals
    console.log(state.totals)
  })
}

const deleteAdminInfo = () => {
  // 发送删除请求
  deleteTheAdminInfo(deleteDialog.id)
  // 修改前端数据
  const index = state.tableData.findIndex(item => item.adminId === deleteDialog.id);
  console.log(index)
  if (index !== -1) {
    state.tableData.splice(index, 1);
  }
  console.log(state.tableData)
  deleteDialog.id = -1
  deleteDialog.visable = false
}

const confirmModify = () => {
  // 发送修改请求
  modifyTheAdminInfo(modifyDialog.data)
  // 修改前端数据
  // 遍历 tableData 数组
  for (let i = 0; i < state.tableData.length; i++) {
    const item = state.tableData[i];
    if (item.adminId === modifyDialog.id) {
      // 找到目标 adminId 对应的元素，进行修改
      state.tableData[i] = { ...item, ...modifyDialog.data };
      break; // 找到后可以提前结束循环
    }
  }
  modifyDialog.visable = false
}

const confirmAdd = () => {
  // 发送添加请求
  addTheAdminInfo(addDialog.data).then((data) => {
    state.currentPage = data.currentPage
    state.totals = data.totals
    state.tableData = data.data
    console.log(state.tableData)
  })
  addDialog.data = {} as AdminInfo
  addDialog.visable = false
}

const changePage = (newPage: number) => {
  state.currentPage = newPage
  getAdminInfo(state.currentPage, state.pageSize).then((data) => {
    state.tableData = data.data
    state.totals = data.totals
    console.log(state.totals)
  })
}

const mhandleAvatarSuccess: UploadProps['onSuccess'] = (
  response,
  uploadFile
) => {
  modifyDialog.data.picUrl = URL.createObjectURL(uploadFile.raw!)
  console.log(modifyDialog.data.picUrl)
}

const mbeforeAvatarUpload: UploadProps['beforeUpload'] = (file) => {
  if (file.type !== 'image/png') {
    ElMessage.error('头像必须是PNG文件！')
    return false
  } else if (file.size / 1024 / 1024 > 2) {
    ElMessage.error('头像不能超过 2MB!')
    return false
  }
  muploadAction(file)
  return false
}

// 照片上传请求
const muploadAction = (file: File) => {
  let formData = new FormData();
  formData.append('file', file)
  const url = 'https://mock.apifox.cn/m1/2794549-0-default/admin/modifyUploadPic';
  try {
    axios.post('https://mock.apifox.cn/m1/2794549-0-default/admin/modifyUploadPic', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    .then((response) => {
      console.log(response)
      modifyDialog.data.picUrl = URL.createObjectURL(file)
      console.log(modifyDialog.data.picUrl)
    })
  } catch (error) {
    console.log(error)
  }
}

const handleAvatarSuccess: UploadProps['onSuccess'] = (
  response,
  uploadFile
) => {
  addDialog.data.picUrl = URL.createObjectURL(uploadFile.raw!)
  console.log(addDialog.data.picUrl)
}

const beforeAvatarUpload: UploadProps['beforeUpload'] = (file) => {
  if (file.type !== 'image/png') {
    ElMessage.error('头像必须是PNG文件！')
    return false
  } else if (file.size / 1024 / 1024 > 2) {
    ElMessage.error('头像不能超过 2MB!')
    return false
  }
  uploadAction(file)
  return false
}

// 照片上传请求
const uploadAction = (file: File) => {
  let formData = new FormData();
  formData.append('file', file)
  const url = 'https://mock.apifox.cn/m1/2794549-0-default/admin/addUploadPic';
  try {
    axios.post('https://mock.apifox.cn/m1/2794549-0-default/admin/addUploadPic', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    .then((response) => {
      console.log(response)
      addDialog.data.picUrl = URL.createObjectURL(file)
      console.log(addDialog.data.picUrl)
    })
  } catch (error) {
    console.log(error)
  }
}

</script>

<style scoped>
.main-card{
  display: flex;
  border-radius: 10px;
  margin-bottom: 50px;
  width: 760px;
  height: 660px;
}

.input-mention{
  font-size: 18px;
}

.space {
  margin-bottom: 20px;
}
</style>


