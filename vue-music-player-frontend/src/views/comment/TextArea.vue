<template>
  <div class="row">
    <Picker :data="emojiIndex" set="twitter" @select="showEmoji" />
  </div>

  <div class="row">
    <div>
      {{ emojisOutput }}
    </div>
  </div>
</template>

<script setup lang="ts">
import data from 'emoji-mart-vue-fast/data/all.json'
import 'emoji-mart-vue-fast/css/emoji-mart.css'
import { Picker, EmojiIndex } from 'emoji-mart-vue-fast/src'
// Create emoji data index.
const emojiIndex = new EmojiIndex(data)
const emojisOutput = ref('')
const showEmoji = (emoji) => {
  this.emojisOutput.value = this.emojisOutput + emoji.native
}
</script>
