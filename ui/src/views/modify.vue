<template>
  <div class="functions">
    <el-button type="info" plain @click="openAddMusicDialog()">添加歌曲</el-button>
    <el-button type="info" plain @click="modifyList()">修改歌单</el-button>
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

    <el-table-column label="name" prop="name" min-width="2"></el-table-column>

    <el-table-column label="options" min-width="1">
      <template #default="item">
        <el-button type="info" size="small" plain @click="openMoveMusicDialog(item.row.id)">移动</el-button>
        <el-button type="info" size="small" plain @click="deleteMusic(item.row.id)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>

  <el-dialog class="add-music-dialog" v-model="showAddMusicDialog" title="添加歌曲">
    <el-form :model="musicIns" label-width="20%">
      <el-form-item label="ID">
        <el-input v-model="musicIns.id" disabled/>
      </el-form-item>

      <el-form-item label="name">
        <el-input v-model="musicIns.name"/>
      </el-form-item>

      <el-form-item label="bv">
        <el-input v-model="musicIns.bv"/>
      </el-form-item>

      <el-form-item label="description">
        <el-input v-model="musicIns.description"/>
      </el-form-item>

      <el-form-item label="volume">
        <el-input-number v-model="musicIns.volume" :min="-100" :max="100" :step="5"/>
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button type="info" plain @click="showAddMusicDialog = false">取消</el-button>
      <el-button type="info" plain @click="addMusic()">确定</el-button>
    </template>
  </el-dialog>

  <el-dialog class="move-music-dialog" v-model="showMoveMusicDialog" title="移动歌曲">
    <div class="mmd-item">
      <el-input v-model="musicIns.name" disabled/>
    </div>

    <div class="mmd-item">
      选择播放列表&#58;&nbsp;
      <el-select v-model="playlistIndexMoveTo" placeholder="请选择播放列表" @change="onChangePlaylist">
        <el-option
          v-for="(item, index) in listStore.list.playlists"
          :key="item.id"
          :label="item.name"
          :value="index"
        />
      </el-select>
    </div>

    <template #footer>
      <el-button type="info" plain @click="showMoveMusicDialog = false">取消</el-button>
      <el-button type="info" plain @click="moveMusic()">确定</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { Music } from "@/axios/list.go"
import { v4 as uuid } from "uuid";
import { useListStore } from "@/pinia/list.ts";

const listStore = useListStore();

let showAddMusicDialog = ref<boolean>(false);
let showMoveMusicDialog = ref<boolean>(false);

let musicIns = ref<Music>(new Music())
let playlistIndexMoveTo = ref<number>(0)

function addMusic(): void {
    listStore.playlist.push(musicIns.value);
    showAddMusicDialog.value = false;
}

function modifyList(): void {
    listStore.modifyList();
}

function moveMusic(): void {
    // add to target playlist
    listStore.list.playlists[playlistIndexMoveTo.value].music_list.push(musicIns.value);

    // del from current playlist
    deleteMusic(musicIns.value.id);

    showMoveMusicDialog.value = false;
}

function deleteMusic(musicID: string): void {
    for (let i = 0; i < listStore.playlist.length; i++) {
        if (listStore.playlist[i].id === musicID) {
            let lastIndex = listStore.playlist.length - 1;
            let musicItem = listStore.playlist[lastIndex];
            listStore.playlist[lastIndex] = listStore.playlist[i];
            listStore.playlist[i] = musicItem;
            listStore.playlist.pop();
            break;
        }
    }
}

function openAddMusicDialog(): void {
    musicIns.value = new Music();
    musicIns.value.id = uuid();
    showAddMusicDialog.value = true;
}

function openMoveMusicDialog(musicID: string): void {
    musicIns.value = new Music();

    for (let i = 0; i < listStore.playlist.length; i++) {
        if (listStore.playlist[i].id === musicID) {
            musicIns.value = listStore.playlist[i];
            break;
        }
    }

    showMoveMusicDialog.value = true;
}

function onChangePlaylist(value: number): void {
    playlistIndexMoveTo.value = value;
}
</script>

<style lang="less">
// modify style in body-level dialog, cancel 'scope' tag
.functions {
  margin: 2vh 0;
}

.add-music-dialog {
  .el-dialog__body {
    padding: 5rem 15%;

    .el-input {
      width: 80%;
    }
  }
}

.move-music-dialog {
  .el-dialog__body {
    padding: 5rem 25%;
    text-align: center;

    .mmd-item {
      margin: 3rem 0;
    }
  }
}
</style>
