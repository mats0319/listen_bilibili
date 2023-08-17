<template>
  <button @click="setVolume()">set volume</button>

  <video id="listenBilibili" controls autoplay>
    <source type="video/mp4" :src=videoURL>
  </video>
</template>

<script setup lang="ts">
import {ref, onMounted} from "vue"

let videoURL = ref("")

onMounted(() => {
  videoURL.value = "https://upos-sz-mirrorali.bilivideo.com/upgcxcode/11/44/1190394411/1190394411-1-16.mp4?e=ig8euxZM2rNcNbRVhwdVhwdlhWdVhwdVhoNvNC8BqJIzNbfq9rVEuxTEnE8L5F6Vn\n" +
      "EsSTx0vkX8fqJeYTj_lta53NCM=&uipk=5&nbs=1&deadline=1692291894&gen=playurlv2&os=alibv&oi=3738481376&trid=f70c82a6ef89487bb6d8ec4d263bf639h&mid=0&platform=html5&upsig=0471829024d44d1430\n" +
      "3b12f3e05ab67d&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,mid,platform&bvc=vod&nettype=0&f=h_0_0&bw=53263&logo=80000000"
})

function setVolume() {
  document.getElementById("listenBilibili")["volume"] = 0.35;
}
</script>

<style scoped></style>
