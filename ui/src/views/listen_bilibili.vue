<template>
  <video id="listenBilibili" controls autoplay>
    <source type="video/mp4" src="">
  </video>
</template>

<script setup lang="ts">
import {ref, watch, onMounted} from "vue"
import {List, Music} from "@/axios/api.pb";
import {axiosInstance} from "@/axios/request";
import {ElMessage} from "element-plus";

let ready = ref(false) // if 'playlist' ready
let list = ref(List)
let playlist = ref(Array<Music>)
let playIndex = ref(0)

onMounted(() => {
  if (!ready.value) {
    getList()
  }

  (document.getElementById("listenBilibili") as HTMLVideoElement).onended = (_: Event): any => {
    playNextMusic()
  };
})

watch(ready, async (newValue: boolean, _: boolean) => {
  if (newValue) {
    playNextMusic()
  }
})

function getList() {
  axiosInstance.getList()
      .then((response: any) => {
        if (response.data.err.length > 0) {
          throw response.data.err
        }

        const listIns = JSON.parse(response.data.list)

        list.value = listIns
        playlist.value = listIns.playlists[0].musicList
        playIndex.value = 0
        ready.value = true

        console.log("> Node: get list success.")
      })
      .catch((err: any) => {
        console.log("> Node: get list failed.", err)
        ElMessage({
          message: "get list failed.",
          type: "error",
        })
      })
}

function playNextMusic() {
  let index = playIndex.value
  playIndex.value = (playIndex.value + 1) % playlist.value.length

  axiosInstance.getOriginURL(playlist.value[index].id)
      .then((response: any) => {
        if (response.data.err.length > 0) {
          throw response.data.err
        }

        (document.getElementById("listenBilibili") as HTMLVideoElement).src = response.data.url;
        (document.getElementById("listenBilibili") as HTMLVideoElement).play();
        (document.getElementById("listenBilibili") as HTMLVideoElement).volume = 0.35;

        console.log("> Node: get origin url success.")
      })
      .catch((err: any) => {
        console.log("> Node: get origin url failed.", err)
        ElMessage({
          message: "get origin url failed.",
          type: "error",
        })
      })
}
</script>

<style scoped lang="less"></style>
