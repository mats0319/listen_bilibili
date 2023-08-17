<template>
  <router-view />
</template>

<script setup lang="ts">
import {onMounted} from "vue";

onMounted(() => {
  console.log("> Node: time.", (new Date()).toISOString())
})
</script>

<style scoped>
</style>
