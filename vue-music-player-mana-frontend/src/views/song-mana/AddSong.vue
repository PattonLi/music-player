<template>
  <!--总体容器el-card-->
  <el-card class="swiper-container">
    <!--增加，删除按钮-->
    <div class="header">
      <el-button type="primary" size="small" @click="handleAdd">增加</el-button>
      <el-popconfirm
        title="确定删除吗？"
        confirmButtonText="确定"
        cancelButtonText="取消"
        @confirm="handleDelete"
      >
        <template #reference>
          <el-button type="danger" size="small">批量删除</el-button>
        </template>
      </el-popconfirm>
    </div>

    <!--主要显示内容区域-->
    <el-table
      ref="multipleTable"
      :data="state.tableData"
      tooltip-effect="dark"
      style="width: 100%"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55"> </el-table-column>
      <el-table-column label="图片" width="120">
        <template #default="scope">
          <img style="width: 150px; height: 150px" :src="scope.row.carouselUrl" alt="轮播图" />
        </template>
      </el-table-column>
      <el-table-column prop="name" label="歌曲名" width="120"> </el-table-column>
      <el-table-column prop="artist" label="演唱者" width="120"> </el-table-column>
      <el-table-column prop="album" label="专辑" width="120"> </el-table-column>
      <el-table-column prop="time" label="发行时间" width="120"> </el-table-column>
      <el-table-column prop="style" label="歌曲分类" width="120"> </el-table-column>
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
</template>

<script setup lang="ts">
import { onMounted, reactive } from 'vue'
import { getSongInfo } from '@/utils/api/song'
import type { SongInfo } from '@/model/SongInfo'

const state = reactive({
  tableData: [] as SongInfo[], // 数据列表
  currentPage: 1, // 当前页数
  pageSize: 10, // 每页请求数
  type: 'add', // 操作类型
  totals: 0,//所有项的总数
  multipleSelection: [] // 选中项
})

//加载数据
onMounted(() => {
  getSongInfo(1, state.pageSize).then((data) => {
    state.totals = data.totals
    state.tableData = data.songInfo
  })
})

const handleAdd = () => {}
const handleDelete = () => {}
const handleSelectionChange = () => {}

//当选择其他的界面
const changePage = (value: number) => {
  state.currentPage = value
  getSongInfo(state.currentPage, state.pageSize).then((data) => {
    state.totals = data.totals
    state.tableData = data.songInfo
  })
}
</script>
