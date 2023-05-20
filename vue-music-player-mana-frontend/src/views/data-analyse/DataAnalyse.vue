<template>
  <el-card class="introduce">
    <div class="order">
      <el-card class="order-item">
        <template #header>
          <div class="card-header">
            <span>今日登录用户数量</span>
          </div>
        </template>
        <div class="item">11323</div>
      </el-card>

      <el-card class="order-item">
        <template #header>
          <div class="card-header">
            <span>今日播放歌曲总量</span>
          </div>
        </template>
        <div class="item">36271</div>
      </el-card>

      <el-card class="order-item">
        <template #header>
          <div class="card-header">
            <span>今日新注册用户数量</span>
          </div>
        </template>
        <div class="item">1242</div>
      </el-card>
    </div>
    <!-- echart 表格-->
    <div id="chart1"></div>
  </el-card>
</template>

<script setup lang="ts">
import * as echarts from 'echarts'
import { onMounted,onUnmounted, reactive, toRefs } from 'vue'
import type { DataAnalyse } from '@/model/DataAnalyse.js'
import { getDataAnalyse } from '@/utils/api/dataAnalys'
let myChart : echarts.ECharts
const state = reactive({
  totalData: {
    numOfDownloadSong:[1,1231,2,2,2,2,2],
    numOfLoginUsers:[1,2,2,2,2,2,2],
    numOfPlaySong:[1,2,2,2,2,2,2],
    numOfRegisterUsers:[1,2,2,2,2,2,2]
  } as DataAnalyse
})


//记得修改
onMounted(() => {
  // 获取数据
    getDataAnalyse().then((data) => {
    console.log('data',data);
    state.totalData.numOfDownloadSong=data.numOfDownloadSong
    console.log('data.numOfDownloadSong',data.numOfDownloadSong);
  })
  console.log('numOfDownloadSong',state.totalData.numOfDownloadSong);
  let array:number[] = state.totalData.numOfDownloadSong
  for (let index = 0; index < state.totalData.numOfDownloadSong.length; index++) {
    console.log(state.totalData.numOfDownloadSong[index]);
    
    array[index]=state.totalData.numOfDownloadSong[index]
  }
  
  
  
  // 建立表格
  let chartDom = document.getElementById('chart1')
  myChart = echarts.init(chartDom as HTMLElement)
  let option

  option = {
    title: {
      text: '大盘数据'
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {
          backgroundColor: '#6a7985'
        }
      }
    },
    legend: {
      data: ['播放歌曲数量', '下载歌曲数量', '登录用户数量', '用户注册数量']
    },
    toolbox: {
      feature: {
        saveAsImage: {}
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: [
      {
        type: 'category',
        boundaryGap: false,
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
      }
    ],
    yAxis: [
      {
        type: 'value'
      }
    ],
    series: [
      {
        name: '播放歌曲数量',
        type: 'line',
        stack: 'Total',
        areaStyle: {},
        emphasis: {
          focus: 'series'
        },
        data: array
      },
      {
        name: '登录用户数量',
        type: 'line',
        stack: 'Total',
        areaStyle: {},
        emphasis: {
          focus: 'series'
        },
        data: [1,3332,2223,4,5123,123,123]
      },
      {
        name: '用户注册数量',
        type: 'line',
        stack: 'Total',
        areaStyle: {},
        emphasis: {
          focus: 'series'
        },
        data: [1,3332,2223,4,5123,123,123]
      },
      {
        name: '下载歌曲数量',
        type: 'line',
        stack: 'Total',
        areaStyle: {},
        emphasis: {
          focus: 'series'
        },
        data: [1,3332,2223,4,5123,123,123]
      }
    ]
  }

  option && myChart.setOption(option)
})

onUnmounted(() => {
  myChart.dispose()
})
</script>

<style scoped>
.introduce .order {
  display: flex;
  margin-bottom: 50px;
}
.introduce .order .order-item {
  /* 子节点占三分之一宽度 */
  flex: 1;
  margin-right: 20px;
}
.introduce .order .order-item:last-child {
  margin-right: 0;
}
#chart1 {
  min-height: 400px;
}
</style>
