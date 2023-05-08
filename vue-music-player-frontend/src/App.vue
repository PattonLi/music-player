<script setup lang="ts"></script>

<template>
  <div class="layout-container-main">
    <el-container>
      <el-aside width="200px">
        <LeftMenu />
      </el-aside>
      <el-container>
        <el-header>Header</el-header>
        <el-main> </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
.layout-container-main {
  height: 100%;
  width: 100%;
  position: relative;
  background-color: var(--el-color-primary-light-7);
  color: var(--el-text-color-primary);
}
</style>
