<template>
  <div>
    <el-menu :default-openeds="['1', '3']">
      <el-sub-menu index="1">
        <el-menu-item index="1-1">Option 1</el-menu-item>
        <el-menu-item index="1-2">Option 2</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="2">
        <el-menu-item index="2-1">Option 1</el-menu-item>
        <el-menu-item index="2-2">Option 2</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="3">
        <el-menu-item index="3-1">Option 1</el-menu-item>
        <el-menu-item index="3-2">Option 2</el-menu-item>
      </el-sub-menu>
    </el-menu>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped></style>
