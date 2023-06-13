<template>
  
  <!--修改对话框-->
  <el-dialog v-model="modifyDialog.visable" title="信息修改" width="27%">
    <el-form :model="modifyDialog.data" label-width="100px">
      <el-form-item label="ID">
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
          <img
            style="width: 147px; height: 147px"
            v-if="modifyDialog.data.picUrl"
            :src="modifyDialog.data.picUrl"
            class="avatar"
          />
          <el-icon v-else><Plus /></el-icon>
        </el-upload>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="modifyDialog.visable = false">Cancel</el-button>
        <el-button type="primary" @click="confirmModify"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { storeToRefs } from 'pinia';
import type { UploadProps } from 'element-plus'
import { AlertError, AlertSuccess } from '@/utils/alert/AlertPop';
import {apiUploadAction} from '@/utils/api/auth'
import myAxios from '@/utils/api/myAxios';
import type{UserProfile} from '@/models/user'

interface CustomerInfo extends UserProfile{
  userId:number
}

const {profile,showModify,userId} = storeToRefs(useAuthStore())

const modifyDialog = reactive({
  visable: showModify.value,
  data: profile.value
})

const mhandleAvatarSuccess: UploadProps['onSuccess'] = (response, uploadFile) => {
  modifyDialog.data.picUrl = URL.createObjectURL(uploadFile.raw!)
  console.log(modifyDialog.data.picUrl)
}

const mbeforeAvatarUpload: UploadProps['beforeUpload'] = async(file) => {
  if (file.type !== 'image/png') {
    AlertError('头像必须是PNG文件！')
    return false
  } else if (file.size / 1024 / 1024 > 2) {
    AlertError('头像不能超过 2MB!')
    return false
  }
  const res= await apiUploadAction(file) as {code:number,picUrl:string}
  if(res.code==200){
    AlertSuccess('上传图片成功！')
    modifyDialog.data.picUrl = res.picUrl
  }else{
    AlertError('上传图片失败')
  }
  return false
}

const confirmModify = () => {
  // 发送修改请求
  const c = {
    age:modifyDialog.data.age,
    email:modifyDialog.data.email,
    gender:modifyDialog.data.gender,
    nickname:modifyDialog.data.gender,
    password:modifyDialog.data.password,
    phone:modifyDialog.data.phone,
    picUrl:modifyDialog.data.picUrl,
    userId:userId.value,
    username:modifyDialog.data.username
  } as CustomerInfo
  modifyTheCustomerInfo(c)
  // 修改前端数据
  modifyDialog.visable = false
}

// 修改用户信息
const modifyTheCustomerInfo = async (data: CustomerInfo) => {
    const res = await myAxios.post<{
      code:number,
      picUrl:string
    }>('/User/modifyInfo', data)
    if(res.code==200){
      AlertSuccess('修改信息成功！')
    }else{
    AlertError('上修改信息失败')
  }
}



</script>
