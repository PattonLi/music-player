<template>
  <div aria-label="user-center" class="p-9 pt-5">
    <el-page-header @back="onBack" title="退出登录" :icon="Logout">
      <!-- 用户头像 -->
      <template #content>
        <div class="flex items-center">
          <el-avatar
            class="mr-3"
            :size="75"
            :src="
              profile.picUrl ||
              'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
            "
          />
          <!-- 用户名 -->
          <span class="text-3xl font-600 mx-3"> {{ profile.nickname }} </span>
          <span class="text-base mr-2" style="color: var(--el-text-color-regular)">
            {{ profile.email }}
          </span>
        </div>
      </template>

      <!-- btn -->
      <template #extra>
        <div class="flex items-center">
          <div class="flex items-center">
            <button class="mr-3 button-outline w-28 h-10" @click="showModify = true">
              编辑信息
            </button>
          </div>
          <div class="flex items-center">
            <button class="mr-3 button-outline w-28 h-10" @click="router.push({ name: 'vip' })">
              开通VIP
            </button>
          </div>
          <div class="flex items-center">
            <button class="button-outline w-28 h-10" @click="router.push({ name: Pages.contact })">
              联系我们
            </button>
          </div>
        </div>
        <EditProfile></EditProfile>
      </template>

      <!-- 描述 -->
      <el-descriptions :column="2" size="large" class="mt-6 w-full">
        <RouterView> </RouterView>
        <el-descriptions-item label="用户名">
          <span class="text-2xl"> {{ profile.username }} </span>
        </el-descriptions-item>
        <el-descriptions-item label="昵称">
          {{ profile.nickname }}
        </el-descriptions-item>
        <el-descriptions-item label="性别">
          {{ profile.gender }}
        </el-descriptions-item>
        <el-descriptions-item label="年龄">
          {{ profile.age }}
        </el-descriptions-item>
        <el-descriptions-item label="电话号码">
          {{ profile.phone }}
        </el-descriptions-item>
        <el-descriptions-item label="邮箱">
          {{ profile.email }}
        </el-descriptions-item>
        <el-descriptions-item label="Remarks">
          <el-tag>普通用户</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="Address"
          >No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu Province
        </el-descriptions-item>
      </el-descriptions>
      <div class="mt-6 text-xl">欢迎您来到QQ音乐网页版(仿),本网站中山大学软件工程学院学生开发!</div>
    </el-page-header>

    <div>
    <!--修改对话框-->
    <el-dialog v-model="showModify" title="信息修改" width="27%">
      <el-form :model="modifyDialog.data" label-width="100px">
        <el-form-item label="">
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
          <el-button @click="showModify = false">Cancel</el-button>
          <el-button type="primary" @click="confirmModify"> Confirm </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
  </div>
</template>

<script setup lang="ts">
import { Logout } from '@icon-park/vue-next'
import { useAuthStore } from '@/stores/auth'
import { storeToRefs } from 'pinia'
import { Pages } from '@/router/pages'
import type { UploadProps } from 'element-plus'
import { AlertError, AlertSuccess } from '@/utils/alert/AlertPop'
import { apiUploadAction } from '@/utils/api/auth'
import myAxios from '@/utils/api/myAxios'
import type { UserProfile } from '@/models/user'

const { profile, showModify,userId } = storeToRefs(useAuthStore())
const { logout, checkLogin } = useAuthStore()
const router = useRouter()

const onBack = () => {
  logout()
}
onMounted(() => {
  checkLogin()
})

interface CustomerInfo extends UserProfile {
  userId: number
}
const modifyDialog = ref({
  data: profile.value
})

const mhandleAvatarSuccess: UploadProps['onSuccess'] = (response, uploadFile) => {
  modifyDialog.value.data.picUrl = URL.createObjectURL(uploadFile.raw!)
  console.log(modifyDialog.value.data.picUrl)
}

const mbeforeAvatarUpload: UploadProps['beforeUpload'] = async (file) => {
  if (file.type !== 'image/png') {
    AlertError('头像必须是PNG文件！')
    return false
  } else if (file.size / 1024 / 1024 > 2) {
    AlertError('头像不能超过 2MB!')
    return false
  }
  const res = (await apiUploadAction(file)) as { code: number; picUrl: string }
  if (res.code == 200) {
    AlertSuccess('上传图片成功！')
    modifyDialog.value.data.picUrl = res.picUrl
  } else {
    AlertError('上传图片失败')
  }
  return false
}

const confirmModify = () => {
  // 发送修改请求
  const c = {
    age: modifyDialog.value.data.age,
    email: modifyDialog.value.data.email,
    gender: modifyDialog.value.data.gender,
    nickname: modifyDialog.value.data.gender,
    password: modifyDialog.value.data.password,
    phone: modifyDialog.value.data.phone,
    picUrl: modifyDialog.value.data.picUrl,
    userId: userId.value,
    username: modifyDialog.value.data.username
  } as CustomerInfo
  modifyTheCustomerInfo(c)
  // 修改前端数据
  showModify.value = false
}

// 修改用户信息
const modifyTheCustomerInfo = async (data: CustomerInfo) => {
  const res = await myAxios.post<{
    code: number
    picUrl: string
  }>('/User/modifyInfo', data)
  if (res.code == 200) {
    AlertSuccess('修改信息成功！')
  } else {
    AlertError('上修改信息失败')
  }
}
</script>
