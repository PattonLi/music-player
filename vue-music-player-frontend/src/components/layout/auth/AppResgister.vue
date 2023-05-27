<template>
  <!--更新信息对话框-->
  <el-dialog
    v-model="showRegister"
    title="新用户注册"
    append-to-body
    @close="clearForm"
    destroy-on-close="tr"
    width="420px"
    class="registerDialog"
  >
    <!--提交表单-->
    <el-form :model="ruleForm" ref="ruleFormRef" :rules="rules" status-icon>
      <el-form-item label="用户名" prop="username">
        <el-input v-model="ruleForm.username" autocomplete="off" />
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model="ruleForm.phone" autocomplete="off" />
      </el-form-item>
      <el-form-item label="昵称" prop="nickname">
        <el-input v-model="ruleForm.nickname" autocomplete="off" />
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="ruleForm.email" autocomplete="off" />
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="ruleForm.password" autocomplete="off" />
      </el-form-item>
      <el-form-item label="确认密码" prop="repassword">
        <el-input v-model="repassword" autocomplete="off" />
      </el-form-item>
    </el-form>
    <!--确认按钮-->
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="clearForm">取消</el-button>
        <el-button type="primary" @click="submitForm(ruleFormRef)"> 确认 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import type { FormRules, FormInstance } from 'element-plus'
import { useAuthStore } from '@/stores/auth'
import { storeToRefs } from 'pinia'
import type { registerUser } from '@/models/user'

//表单相关
const ruleFormRef = ref<FormInstance>()
const { ruleForm, showRegister } = storeToRefs(useAuthStore())
const repassword = ref('')

const rules = reactive<FormRules>({
  username: [{ required: true, message: '不得为空', trigger: 'blur' }],
  phone: [{ required: true, message: '不得为空', trigger: 'blur' }],
  nickname: [{ required: true, message: '不得为空', trigger: 'blur' }],
  email: [{ required: true, message: '不得为空', trigger: 'change' }],
  password: [{ required: true, message: '不得为空', trigger: 'change' }],
  repassword: [{ required: true, message: '不得为空', trigger: 'change' }]
})

//提交表单
const submitForm = async (formEl: FormInstance | undefined) => {
  //如果表单不为undefined
  if (!formEl) return
  await formEl.validate((valid: any, fields: any) => {
    if (valid) {
      console.log()
    }
  })
}

//清空表单
const clearForm = () => {
  ruleForm.value = {} as registerUser
  showRegister.value = false
  repassword.value = ''
}
</script>

<style scoped lang="scss">
.registerDialog {
    
}
</style>
