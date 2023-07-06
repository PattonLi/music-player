<template>
  <div class="flex">
    <!-- img -->
    <div class="w-12 h-12 flex-shrink-0 mt-1">
      <img :src="props.comment.picUrl" alt="img" class="rounded-full" />
    </div>
    <div class="flex flex-col ml-4 justify-between mr-10">
      <!-- text -->
      <div>
        <span class="text-sky-600 text-main text-sm cursor-pointer"
          >{{ props.comment.nickname }}：</span
        >
        <span class="text-main">{{ props.comment.comment }}</span>
      </div>
      <!-- info -->
      <div class="flex items-center mt-1.5">
        <!-- date -->
        <div>
          <span class="text-dc text-sm">{{ props.comment.commentTime }}</span>
        </div>
        <!-- like -->
        <div class="flex">
          <!-- 点赞按钮 -->
          <div class="items-center ml-4 flex flex-1 flex-shrink-0">
            <IconPark
              v-if="!isSongLike"
              :icon="ThumbsUp"
              size="20"
              :stroke-width="3"
              class="text-gray-400 ml-3 mr-2 cursor-pointer hover:text-red-400"
              @click="apiLikeComment(comment.commentId, userId);refresh()"
            />
            <IconPark
              v-else
              :icon="ThumbsUp"
              size="20"
              :stroke-width="3"
              theme="filled"
              fill="#d0021b"
              @click="apiDelCommentLike(comment.commentId, userId);refresh()"
              class="text-gray-400 ml-3 mr-2 cursor-pointer"
            />
          </div>
          <span>{{ comment.like }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Comment } from '@/models/comment'
import { ThumbsUp } from '@icon-park/vue-next'
import { useAuthStore } from '@/stores/auth'
import { AlertError } from '@/utils/alert/AlertPop'
import { apiGetCommentLikes, apiLikeComment, apiDelCommentLike } from '@/utils/api/comment'
import { storeToRefs } from 'pinia'
const { userId } = storeToRefs(useAuthStore())
const props = defineProps<{
  comment: Comment
}>()
const likes = ref<number[]>()

const isSongLike = computed(() => {
  if (props.comment.commentId in likes) return true
  else return false
})

const refresh = async () => {
  const data = await apiGetCommentLikes(userId.value)
  if (data.code == 200) {
    likes.value = data.commentLikes
  } else {
    AlertError('获取点赞信息失败')
  }
}
onMounted(()=>{
  refresh()
})
</script>
