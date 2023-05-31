<template>
  <div>
    <audio ref="audioElement" controls></audio>
    <img :src="coverImageUrl" alt="Cover Image">
  </div>
</template>

<script setup lang="ts">

const audioElement = ref(null);
const coverImageUrl = ref('');

onMounted(() => {
  audioElement.value.addEventListener('loadedmetadata', extractCoverImage);
});

const extractCoverImage = () => {
  if (audioElement.value && audioElement.value.src) {
    const covers = audioElement.value.getAttribute('poster');
    if (covers) {
      coverImageUrl.value = covers;
    } else {
      coverImageUrl.value = '';
    }
  }
};
</script>
