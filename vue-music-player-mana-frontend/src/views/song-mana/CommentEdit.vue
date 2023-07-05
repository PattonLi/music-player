<template>
    <el-card class="main-card">
      <el-row class="el-row">
        <el-col :span="4" class="el-col">
          <el-card class="user-card" shadow="hover" :body-style="{ padding: '0px' }">
            <img class="image" :src="state.songData.picUrl" :fit="'cover'" />
            <div style="margin: 14px">
              <span style="font-size: 25px; font-weight: bold">《{{ state.songData.name }}》</span>
              <!--<span style="font-size: 25px; font-weight: bold; margin-bottom: 10px">的评论</span>-->
            </div>
          </el-card>
        </el-col>
        <!--歌曲-->
        <el-col :span="10" class="el-col-3">
          <el-card class="data-card" :body-style="{ padding: '0px' }">
            <template #header>
              <div class="card-header">
                <span style="font-size: 30px; font-weight: bold">看看歌友评论</span>
              </div>
            </template>
            <el-table :data="state.commentData" style="width: 950px" :height="530">
              <el-table-column prop="picUrl" label="用户头像" width="100">
                <template #default="scope">
                  <el-avatar shape="square" :size="70" :fit="'cover'" :src="scope.row.picUrl" />
                </template>
              </el-table-column>
              <el-table-column prop="nickname" label="名称" width="100" />
              <el-table-column prop="comment" label="评论" width="320" />
              <el-table-column prop="commentTime" label="评论时间" width="170" />
              <el-table-column prop="like" label="点赞数" width="100" />
              <el-table-column prop="like" label="删除" width="120">
                <template #default="scope">
                  <el-button type="danger" :icon="Delete" @click="DeleteClick(scope.row.commentId)"
                    >删除</el-button
                  >
                </template>
              </el-table-column>
            </el-table>
          </el-card>
        </el-col>
      </el-row>
    </el-card>
    <!--删除对话框-->
    <el-dialog v-model="commentDeleteDialog.visable" title="系统提示" width="20%">
    <span>是否确认删除？</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="commentDeleteDialog.visable = false">取消</el-button>
        <el-button type="primary" @click="confirm"> 确认 </el-button>
      </span>
    </template>
  </el-dialog>
  <el-row :gutter="10">
    <el-col :span="7">
    <!--分页按钮-->
    <el-pagination
      background
      layout="prev, pager, next"
      :total="state.totals"
      :page-size="state.pageSize"
      :current-page="state.currentPage"
      @current-change="handleCurrentChange"
    />
    </el-col>
    <el-col :span="5">
      <el-input v-model="state.inputs" placeholder="请输入内容"></el-input>
    </el-col>
    <el-col :span="5">
      <el-button type="success" :icon="Search" @click="SearchClick">查询</el-button>
    </el-col>
    </el-row>
  </template>
  
  <script setup lang="ts">
  import type { CommentInfo } from '@/model/Comment'
  import type { SongInfo } from '@/model/SongInfo'
  import type { ArtistInfo } from '@/model/ArtistInfo'
  import type { AlbumInfo } from '@/model/albumInfo'
  import { getACustomerInfo } from '@/utils/api/user'
  import { getUserLikeInfo } from '@/utils/api/like'
  import { getSongInfo, getTheSongInfo } from '@/utils/api/song'
  import { getTheCommentInfo, deleteTheCommentInfo } from '@/utils/api/comment'
  import { Check, Delete, Edit, Search, Share, Upload } from '@element-plus/icons-vue'
  
  const state = reactive({
    songData: {
      songId: 0, //歌曲ID
      albumId: 0, //专辑id
      artistId: 0, //歌手id

      name: '', //歌曲名称
      album: '', //歌曲专辑名
      artist: '', //歌手名称

      pop: 0, //歌曲播放次数
      mark: 0, //评论数
      publishTime: '', //发行时间
      type: '', //歌曲类别
      duration: 0, //歌曲秒数时长

      url: '', //歌曲url
      lyricUrl: '', //歌词url
      picUrl: '', //图片url，根据专辑id,可以减少数量
    } as SongInfo, // 歌曲信息
    commentData: [] as CommentInfo[], // 专辑信息
    currentPage: 1,
    pageSize: 1,
    totals: 400, // 一共有多少条目
    inputs:''
  })
  
  //加载数据
  onMounted(async() => {
    state.currentPage = 1
    console.log('state.currentPage');
    console.log(state.currentPage);
    await getSongInfo(state.currentPage, 1).then((data) => {
      if (data.code == 200) {
        state.songData = data.data[0]
        state.totals = data.totals
        console.log('state.totals')
        console.log(state.totals)
        console.log('data')
        console.log(data.data)
        console.log('state.songData')
        console.log(state.songData)
      }
    })
    console.log('state.songData.songId' + state.songData.songId)
    await getTheCommentInfo(state.songData.songId).then((data) => {
      if (data.code == 200) {
        state.commentData = data.data
        console.log('state.commentData')
        console.log(state.commentData)
      }
    })
  })
  const handleCurrentChange = async(val: number) => {
    state.currentPage = val
    console.log(`current page: ${val}`)
    await getSongInfo(val, 1).then((data) => {
      if (data.code == 200) {
        state.songData = data.data
        state.totals = data.totals
        console.log(state.totals)
        console.log(data.data)
        console.log('data')
        console.log(data.data)
        console.log('state.songData')
        console.log(state.songData)
      }
    })
    console.log('state.userData.userId' + state.songData.songId)
    await getTheCommentInfo(state.songData.songId).then((data) => {
      if (data.code == 200) {
        state.commentData = data.data
        console.log('state.commentData')
        console.log(state.commentData)
      }
    })
  }
  const commentDeleteDialog = reactive({
    visable: false,
    id: -1
  })
  const DeleteClick = async(id: number) => {
    console.log('click')
    commentDeleteDialog.id = id
    commentDeleteDialog.visable = true
  }
  const confirm = () => {
  // 发送删除请求
  deleteTheCommentInfo(commentDeleteDialog.id)
  // 修改前端数据
  const index = state.commentData.findIndex((item) => item.commentId === commentDeleteDialog.id)
  console.log(index)
  if (index !== -1) {
    state.commentData.splice(index, 1)
  }
  console.log(state.commentData)
  commentDeleteDialog.id = -1
  commentDeleteDialog.visable = false
}
const SearchClick = async() => {
    getTheSongInfo(state.inputs).then((data) => {
    if (data.code === 300) {
      ElMessage.warning('找不到该音乐!')
    } else if (data.code === 200) {
      console.log(data.data)
      state.songData = data.data[0]
    }
  })
  console.log('state.songData.songId' + state.songData.songId)
    await getTheCommentInfo(state.songData.songId).then((data) => {
      if (data.code == 200) {
        state.commentData = data.data
        console.log('state.commentData')
        console.log(state.commentData)
      }
    })
  }
  </script>
  
  <style scoped>
  .main-card {
    display: flex;
    border-radius: 10px;
    margin-bottom: 20px;
    width: auto;
    height: 750px;
  }
  .el-row {
    margin-bottom: 40px;
  }
  
  .el-col {
    margin-left: 20px;
  }
  .el-col-3 {
    margin-left: 100px;
  }
  .user-card {
    margin-top: 10px;
    margin-bottom: 10px;
    display: flex;
    width: 300px;
    height: 375px;
  }
  .data-card {
    margin-top: 10px;
    margin-bottom: 10px;
    margin-left: 30px;
    width: 920px;
    height: 650px;
  }
  .image {
    width: 300px;
    height: 300px;
    display: block;
  }
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .infinite-list {
    height: 500px;
    padding: 0;
    margin: 0;
    list-style: none;
  }
  .data {
    font-size: 20px;
    margin-left: 20px;
  }
  .pic {
    border-radius: 50%;
    margin-left: 20px;
    margin-top: 10px;
    margin-bottom: 10px;
    margin-right: 15px;
  }
  .pic-2 {
    margin-left: 20px;
    margin-top: 10px;
    margin-bottom: 10px;
    margin-right: 15px;
  }
  .text {
    font-size: 15px;
    font-family: cursive;
    color: darkslategray;
    margin-left: 10px;
    margin-right: 10px;
  }
  </style>
  