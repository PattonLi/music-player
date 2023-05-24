<template>
  <!--选择方式-->
  <el-card class="select">
    <el-select v-model="state.dataType" placeholder="选择时间范围" size="large">
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
    <!--选择日期-->
    <el-date-picker
      v-model="state.data"
      v-if="state.dataType == 'custom'"
      type="date"
      placeholder="Pick a day"
      size="large"
    />
  </el-card>

  <!--时间线-->
  <el-timeline class="infinite-list" style="overflow: auto">
    <el-timeline-item
      v-for="i in count"
      :key="i"
      timestamp="2018/4/12"
      placement="top"
      class="infinite-list-item"
    >
      <el-card>
        <h4>Update Github template</h4>
        <p>Tom committed 2018/4/12 20:46</p>
      </el-card>
    </el-timeline-item>
  </el-timeline>
</template>

<script lang="ts" setup>
const count = ref(2)
const state = reactive({
  dataType: 'today',
  //选择具体日期才要
  data: ''
})
//选项
const options = [
  {
    value: 'today',
    label: '今日信息'
  },
  {
    value: 'week',
    label: '本周信息'
  },
  {
    value: 'month',
    label: '本月信息'
  },
  {
    value: 'custom',
    label: '自定义日期'
  }
]
</script>

<style>
.select {
  /* width: 100vw;   */
  margin-top: 20px;
  margin-right: 50px;
  margin-left: 50px;
  display: flex;
}

.infinite-list {
  padding: 20px;
}

.infinite-list .infinite-list-item {
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
  margin: 20px;
}
</style>
