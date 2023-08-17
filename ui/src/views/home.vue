<template>
  <button @click="link('/listen-bilibili', 'ListenBilibili')">listen bilibili</button>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
const router = useRouter()

function link(path: string, name: string) {
  if (location.href.includes(path)) {
    console.log("> Node: you are trying to link to current page.")
    return
  }

  router.push({name: name})
}
</script>

<style scoped></style>
