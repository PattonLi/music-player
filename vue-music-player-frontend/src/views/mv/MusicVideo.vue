<template>
  <el-button @click="router.push({name:'album',query:{id:'1'}})">专辑1</el-button>
</template>

<script setup lang="ts">
const router = useRouter()
</script>
