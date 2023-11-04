<template>
  <div class="content">
    <div class="list">
      <div class="select-playlist">
        选择播放列表&#58;&nbsp;
        <el-select v-model="playlistName" placeholder="请选择播放列表" @change="onChangePlaylist">
          <el-option
            v-for="item in listStore.list.playlists"
            :key="item.id"
            :label="item.name"
            :value="item.music_list"
          />
        </el-select>
        当前列表包含歌曲数量&#58;&nbsp;{{ listStore.playlist.length }}
      </div>

      <el-table :data="listStore.playlist" max-height="60vh" stripe highlight-current-row>
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
    </div>

    <el-divider class="divider" direction="vertical"/>

    <div class="control">
      <div class="c-item">
        <el-button type="info" plain @click="listStore.getList()">重新加载歌单</el-button>
      </div>

      <div class="c-item">
        <el-input-number v-model="baseVolume" :min="0" :max="1" :step="0.05"/>
        <el-button class="marginHorizontal" type="info" plain @click="setVideoVolume">设置基础音量</el-button>
      </div>

      <div class="c-item">正在播放&#58;&nbsp;{{ listStore.musicName }}</div>

      <div class="c-item">{{ volumeStr }}</div>

      <video id="listenBilibili" width="450" height="245" controls autoplay>
        <source type="video/mp4" src="">
      </video>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from "vue"
import { useListStore } from "@/pinia/list.ts";
import { Music } from "@/axios/list.go.ts";

const listStore = useListStore()

let playlistName = ref<string>("")

let baseVolume = ref<number>(0.35)
let volume = ref<number>(0.35)

onMounted(() => {
    if (!(listStore.list.playlists?.length && listStore.list.playlists.length > 0)) {
        listStore.getList()
    }

    // play next music when current one is finished
    (document.getElementById("listenBilibili") as HTMLVideoElement).onended = (_: Event): any => {
        listStore.playNextMusic()
    };
})

function playMusic(id: string): void {
    listStore.musicID = id
    listStore.playMusic()
}

function setVideoVolume(): void {
    volume.value = baseVolume.value + listStore.volume;
    if (volume.value < 0) {
        volume.value = 0
    } else if (volume.value > 1) {
        volume.value = 1
    }

    (document.getElementById("listenBilibili") as HTMLVideoElement).volume = volume.value
}

function onChangePlaylist(value: Array<Music>): void {
    listStore.playlist = value
}

const volumeStr = computed<string>(() => {
    return "当前音量：" + volume.value.toString()
})

watch(() => listStore.originURL, (newValue: string) => {
    (document.getElementById("listenBilibili") as HTMLVideoElement).src = newValue;
    setVideoVolume();
    (document.getElementById("listenBilibili") as HTMLVideoElement).play();
})
</script>

<style scoped lang="less">
.content {
  display: flex;
  width: 100%;
  height: 100%;

  .list {
    width: 60%;

    .select-playlist {
      height: 10vh;
      font-size: 1.4rem;
    }
  }

  .divider {
    height: inherit;
  }

  .control {
    width: 40%;

    .c-item {
      display: flex;
      height: 4rem;
      line-height: 4rem;
      padding-bottom: 2rem;
      font-size: 1.4rem;

      .marginHorizontal {
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
