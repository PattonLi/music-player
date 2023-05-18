<template>
  <el-aside class="aside">
    <!--icon-->
    <div class="head">
      <div>
        <span>music manager</span>
        <img src="//s.weituibao.com/1582958061265/mlogo.png" alt="logo" />
      </div>
    </div>
    <!--icon menu devide-->
    <div class="line"></div>
    <el-menu backgroud-color="#222832" :router="true">
      <!--一级menu-->
      <el-sub-menu index="1">
        <template #title>
          <span>功能区1</span>
        </template>
        <!--二级menu-->
        <el-menu-item-group>
          <el-menu-item index="/">
            <el-icon> <DataLine /> </el-icon>首页
          </el-menu-item>
          <el-menu-item index="/addsong">
            <el-icon> <DataLine /> </el-icon>添加歌曲
          </el-menu-item>
          <el-icon> <InfoFilled /> </el-icon>系统介绍
        </el-menu-item-group>
      </el-sub-menu>
      <el-sub-menu index="2">
        <template #title>
          <span>功能区2</span>
        </template>
      </el-sub-menu>
    </el-menu>
  </el-aside>
</template>

<script setup lang="ts"></script>

<style scoped>
.line {
  color: black;
}

.layout {
  min-height: auto;
  background-color: #ffffff;
}

.container {
  height: 100vh;
}

.aside {
  width: 300px;
  background-color: #222832;
}

.head {
  display: flex;
  align-items: center;
}

.head > div {
  display: flex;
  align-items: center;
}

.head img {
  width: 50px;
  height: 50px;
}

.head span {
  color: antiquewhite;
  font-size: medium;
}

.line {
  border-top: 5px solid hsla(0, 0%, 100%, 0.05);
  border-bottom: 5px solid rgba(0, 0, 0, 0.2);
}
</style>
