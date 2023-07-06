<template>
  <div class="mx-8 mb-3 mt-6">
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
  <div class="flex justify-end mr-7 items-center">
    <button @click="ifEmoji = !ifEmoji" class="mr-4 hover-text">
      <IconPark :icon="SlightlySmilingFace" size="26" :stroke-width="3.5" />
    </button>
    <button class="button-outline w-16 h-8" @click="pubulish">发布</button>
  </div>
  <!-- 表情包 -->
  <div class="flex justify-end" v-if="ifEmoji">
    <Picker
      class="emoji"
      :data="emojiIndex"
      set="apple"
      :showSearch="false"
      :showSkinTones="false"
      title=""
      :emojiSize="22"
      :perLine="8"
      :showPreview="false"
      @select="showEmoji"
    />
  </div>
</template>

<script setup lang="ts">
import data from 'emoji-mart-vue-fast/data/all.json'
import 'emoji-mart-vue-fast/css/emoji-mart.css'
import { SlightlySmilingFace } from '@icon-park/vue-next'
import { Picker, EmojiIndex } from 'emoji-mart-vue-fast/src'
import { ref } from 'vue'
import { AlertError, AlertSuccess } from '@/utils/alert/AlertPop'
import { apiPublishComment } from '@/utils/api/comment'
import { useAuthStore } from '@/stores/auth'
import { storeToRefs } from 'pinia'
const { userId } = storeToRefs(useAuthStore())
// Create emoji data index.
const emojiIndex = new EmojiIndex(data)
const comment = ref('')
const ifEmoji = ref(false)

const props = defineProps<{
  songId: number
}>()
const refresh = inject('refreshComment') as any

const showEmoji = (emoji: any) => {
  comment.value = comment.value + emoji.native
}
const pubulish = async () => {
  if (comment.value == '') {
    AlertError('输入不得为空')
  } else {
    const data = await apiPublishComment(comment.value, userId.value, props.songId)
    if (data.code == 200) {
      AlertSuccess('评论成功')
      comment.value = ''
      await refresh()
    } else {
      AlertError('评论失败，请稍后再试')
    }
  }
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
.emoji {
  @apply relative top-6 mr-8;
  z-index: 1;
  margin-bottom: -27rem;
}
</style>
