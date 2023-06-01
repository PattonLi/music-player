<template>
  <el-button @click="alert">asdasdas</el-button>
</template>

<script setup lang="ts">
import { AlertSuccess } from '@/utils/alert/AlertPop'
const alert = () => {
  AlertSuccess('success')
}
const router = useRouter()
</script>
