<template>
  <el-button 
  @click="routerPushByNameId(Pages.mvPlayer,1)"
  >
  测试mv
</el-button>
    
</template>

<script setup lang="ts">
import { Pages } from '@/router/pages';
import { routerPushByNameId } from '@/utils/navigator/router';

</script>
