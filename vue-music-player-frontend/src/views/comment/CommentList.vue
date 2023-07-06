<template>
  <div>
    <!-- 评论标题 -->
    <div class="text-xl text-main ml-8 mt-0 flex items-center">
      <div>精彩评论</div>
      <IconPark :icon="MessageEmoji" size="22" :stroke-width="3" class="ml-1" />
    </div>
    <div v-if="comments != undefined">
      <div
        class="text-main ml-8 mt-4 flex items-center"
        v-for="(item, index) in comments"
        :key="index"
      >
        <CommentItem :comment="item"></CommentItem>
      </div>
    </div>
    <div v-else>
      <span class="ml-8 text-xl text-dc">抱歉，暂无评论</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import CommentItem from '@/components/common/CommentItem.vue'
import { MessageEmoji } from '@icon-park/vue-next'
import IconPark from '@/components/common/IconPark.vue'
import { apiGetSongComment } from '@/utils/api/comment'
import { AlertError } from '@/utils/alert/AlertPop'
import type { Comment } from '@/models/comment'

const props = defineProps<{ comments: Comment[] }>()
const refresh = inject('refreshComment') as any
onMounted(async () => {
  refresh()
})
</script>
