<template>
  <div class="listen-content">
    <el-table :data="listStore.currentMusicList" max-height="70vh" stripe highlight-current-row>
      <el-table-column type="expand">
        <template #default="item">
          <el-descriptions title="Music Details" :column="1" size="small" border>
            <el-descriptions-item label="ID">{{ item.row.id }}</el-descriptions-item>
            <el-descriptions-item label="Name">{{ item.row.name }}</el-descriptions-item>
            <el-descriptions-item label="BV">{{ item.row.bv }}</el-descriptions-item>
            <el-descriptions-item label="Description">{{ item.row.description }}</el-descriptions-item>
            <el-descriptions-item label="Volume">{{ item.row.volume / 100 }}</el-descriptions-item>
          </el-descriptions>
        </template>
      </el-table-column>

      <el-table-column label="name" prop="name" min-width="3"></el-table-column>

      <el-table-column label="options" min-width="1">
        <template #default="item">
          <el-button type="info" size="small" plain @click="playMusic(item.row.id)">Play</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-divider class="lc-divider" direction="vertical"/>

    <div class="lc-control">
      <div class="lcc-item">
        <el-input-number v-model="baseVolume" :min="0" :max="1" :step="0.05"/>
        <el-button class="lcci-marginHorizontal" type="info" plain @click="setVideoVolume">应用基础音量修改</el-button>
      </div>

      <div class="lcc-item">正在播放&#58;&nbsp;{{ currentMusicName }}</div>

      <div class="lcc-item">当前音量&#58;&nbsp;{{ volume }}</div>

      <video id="listen-bilibili" width="450" height="245" controls autoplay>
        <source type="video/mp4" src="">
      </video>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useListStore } from "@/pinia/list.ts";

const listStore = useListStore()

let baseVolume = ref<number>(0.35)
let volume = ref<number>(0.35) // final volume, base volume + volume offset of each music

let currentMusicName = ref<string>("");
let currentMusicVolumeOffset = ref<number>(0);

onMounted(() => {
    // play next music when current one is finished
    (document.getElementById("listen-bilibili") as HTMLVideoElement).onended = (_: Event): any => {
        playMusic()
    };

    // del 'onended' event
    window.addEventListener("beforeunload", () => {
        (document.getElementById("listen-bilibili") as HTMLVideoElement).onended = () => {}
    })
})

function playMusic(musicID: string = ""): void {
    listStore.playMusic(musicID, (url: string, name: string, volumeOffset: number) => {
        currentMusicName.value = name;
        currentMusicVolumeOffset.value = volumeOffset;

        // play new music
        (document.getElementById("listen-bilibili") as HTMLVideoElement).src = url;
        setVideoVolume();
        (document.getElementById("listen-bilibili") as HTMLVideoElement).play();
    });
}

function setVideoVolume(): void {
    volume.value = baseVolume.value + currentMusicVolumeOffset.value;
    if (volume.value < 0) {
        volume.value = 0
    } else if (volume.value > 1) {
        volume.value = 1
    }

    (document.getElementById("listen-bilibili") as HTMLVideoElement).volume = volume.value
}
</script>

<style scoped lang="less">
.listen-content {
  display: flex;
  width: 100%;
  height: 100%;

  .el-table {
    width: 60%;
  }

  .lc-divider {
    height: inherit;
  }

  .lc-control {
    width: 40%;

    .lcc-item {
      display: flex;
      height: 4rem;
      line-height: 4rem;
      padding-bottom: 2rem;
      font-size: 1.4rem;

      .lcci-marginHorizontal {
        margin-left: 3rem;
        margin-right: 3rem;
      }

      .el-input {
        width: fit-content;
      }

      .el-button {
        height: inherit;
      }
    }
  }
}
</style>
