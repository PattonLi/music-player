<template>
  <div class="player-volume flex flex-col items-center pt-2">
    <!-- 拖动条 -->
    <div>
      <el-slider
        vertical
        height="100px"
        :show-tooltip="false"
        v-model="volume"
        :max="100"
        :min="0"
        size="small"
        :disabled="muted"
        @input="setVolume"
      />
    </div>
    <!-- 音量数字 -->
    <div class="text-sm mt-2">{{ volume }}</div>
  </div>
</template>

<script setup lang="ts">
import { usePlayerStore } from '@/stores/player'

const { volume, muted, toggleMuted, setVolume } = toRefs(usePlayerStore())
</script>

<style lang="scss"></style>
