<template>
  <router-view />
</template>

<script setup lang="ts">
import {onMounted} from "vue";

onMounted(() => {
  console.log("> Node: time.", (new Date()).toISOString())
})
</script>

<style lang="less">
html {
  font-size: 62.5%; // 1rem = 10px
}

body {
  padding: 0;
  margin: 0;
}
</style>
