<template>
  <router-view/>
</template>

<script setup lang="ts">
import { onMounted } from "vue"
import config from "../package.json"

onMounted(() => {
    console.log("> version: v" + config.version + "-" + config.description);
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
