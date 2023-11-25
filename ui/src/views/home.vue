<template>
  <div class="home-top">
    <div class="ht-title" @click="link('home')">Listen Bilibili</div>

    <div class="ht-content">
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
      <el-button type="info" plain @click="listStore.getList(true)">重新加载歌单</el-button>
    </div>

    <div class="ht-code">
      <a href="https://github.com/mats9693/listenBilibili" target="_blank">Github</a>
    </div>

  </div>

  <div class="home-content">
    <div class="hc-left">
      <p class="hcl-item" @click="link('listen')">Listen</p>
      <p class="hcl-item" @click="link('modify')">Modify List</p>
    </div>

    <el-divider class="hc-divider" direction="vertical"/>

    <div class="hc-right">
      <router-view/>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, onMounted } from "vue";
import { Music } from "@/axios/list.go.ts";
import { useListStore } from "@/pinia/list.ts";

const router = useRouter()
const listStore = useListStore();

let playlistName = ref<string>("")

onMounted(() => {
    if (!(listStore.list.playlists?.length && listStore.list.playlists.length > 0)) {
        listStore.getList()
    }
})

function onChangePlaylist(value: Array<Music>): void {
    listStore.playlist = value
}

function link(name: string) {
    router.push({ name: name })
}
</script>

<style scoped lang="less">
.home-top {
  display: flex;
  height: 10rem;
  background-color: lightgray;

  .ht-title {
    width: 20vw;
    padding-left: 5vw;

    line-height: 10rem;
    font-size: 3rem;
  }

  .ht-title:hover {
    cursor: pointer;
  }

  .ht-content {
    line-height: 10rem;
    width: 55vw;
    font-size: 1.6rem;

    .el-select {
      padding-right: 3rem;
    }

    .el-button {
      margin-left: 5rem;
    }
  }

  .ht-code {
    width: 20vw;

    line-height: 10rem;
    font-size: 2.5rem;

    a {
      color: black;
      text-decoration-line: none;
    }
  }
}

.home-content {
  display: flex;
  height: calc(100vh - 10rem);

  .hc-left {
    width: 14vw;
    padding: 5vh 3vw;

    .hcl-item {
      height: 4rem;
      line-height: 4rem;
      font-size: 2rem;
    }

    .hcl-item:hover {
      cursor: pointer;
      background-color: aliceblue;
    }
  }

  .hc-divider {
    height: inherit;
  }

  .hc-right {
    width: 70vw;
    padding: 5vh 3vw;
  }
}
</style>
