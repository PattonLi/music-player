<template>
  <div class="mx-6 mb-3">
    <!-- 输入框 -->
    <el-input
      v-model="comment"
      maxlength="200"
      placeholder="期待你的神评论 !"
      type="textarea"
      :show-word-limit="true"
      clearable:true
      rows="3"
      resize="none"
      input-style="height:100px"
      class="myInput"
    />
  </div>
  <!-- 发布按钮 -->
  <div class="flex justify-end mr-7 justify-between">
    <button @click="ifEmoji=!ifEmoji">
      <IconPark :icon="SlightlySmilingFace" size="22" :stroke-width="3.5"/>
    </button>
    <button class="button-outline w-16 h-8">发布</button>
  </div>
  <!-- 表情包 -->
  <div class="row" v-if="ifEmoji">
    <Picker 
      :data="emojiIndex" 
      set="apple"
      :showSearch="false" 
      :showSkinTones="false"
      title=''
      :emojiSize="22"
      :perLine="8"
      :showPreview="false"
      @select="showEmoji" 
    />
  </div>
  <!-- 评论标题 -->
  <div class="text-xl text-main mt-0.5 mt-2 flex items-center">
    <div>精彩评论</div>
    <IconPark :icon="MessageEmoji" size="22" :stroke-width="3" class="ml-1" />
  </div>
  
  
</template>

<script setup lang="ts">
import data from 'emoji-mart-vue-fast/data/all.json'
import 'emoji-mart-vue-fast/css/emoji-mart.css'
import { Picker, EmojiIndex } from 'emoji-mart-vue-fast/src'
import { ref } from 'vue';
// Create emoji data index.
const emojiIndex = new EmojiIndex(data)
const comment = ref('')
const ifEmoji = ref(false)
import { MessageEmoji,SlightlySmilingFace } from '@icon-park/vue-next'
import IconPark from '@/components/common/IconPark.vue'

defineProps<{
  songId: number
}>()

const showEmoji = (emoji:any) => {
  comment.value = comment.value + emoji.native
}
</script>

<style lang="scss">
.myInput {
  @apply mt-0.5;
  .el-textarea__inner {
    @apply rounded-xl text-base bg-slate-50 dark:bg-zinc-900 dark:ring-gray-700 dark:ring-2;
  }
  .el-input__count {
    @apply dark:bg-zinc-900;
  }
}

.row{
  
}
</style>