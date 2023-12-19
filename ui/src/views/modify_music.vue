<template>
    <div class="functions">
        <el-button type="info" plain @click="openAddMusicDialog()">添加歌曲</el-button>
        <el-button class="f-right" type="info" plain @click="listStore.modifyList()">保存修改</el-button>
    </div>

    <el-table :data="listStore.currentMusicList" max-height="60vh" stripe highlight-current-row>
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
                <el-button type="info" size="small" plain @click="openModifyMusicDialog(item.row.id)">修改</el-button>
                <el-button type="info" size="small" plain @click="openMoveMusicDialog(item.row.id)">移动</el-button>
                <el-button type="info" size="small" plain @click="deleteMusic(item.row.id)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>

    <el-dialog class="add-music-dialog" v-model="showAddMusicDialog" title="添加歌曲">
        <el-form :model="musicIns" label-width="20%">
            <el-form-item label="ID">
                <el-input v-model="musicIns.id" disabled />
            </el-form-item>

            <el-form-item label="name">
                <el-input v-model="musicIns.name" />
            </el-form-item>

            <el-form-item label="bv">
                <el-input v-model="musicIns.bv" />
            </el-form-item>

            <el-form-item label="description">
                <el-input v-model="musicIns.description" />
            </el-form-item>

            <el-form-item label="volume">
                <el-input-number v-model="musicIns.volume" :min="-100" :max="100" :step="5" />
            </el-form-item>
        </el-form>

        <template #footer>
            <el-button type="info" plain @click="showAddMusicDialog = false">取消</el-button>
            <el-button type="info" plain @click="addMusic()">确定</el-button>
        </template>
    </el-dialog>

    <el-dialog class="modify-music-dialog" v-model="showModifyMusicDialog" title="修改歌曲信息">
        <el-form :model="newMusicIns" label-width="20%">
            <el-form-item label="ID">
                <el-input v-model="newMusicIns.id" disabled />
            </el-form-item>

            <el-form-item class="mmd-change-input" label="name">
                <el-input v-model="musicIns.name" readonly />
                <div class="mmdci-icon">&#45;&#62;</div>
                <el-input v-model="newMusicIns.name" />
            </el-form-item>

            <el-form-item class="mmd-change-input" label="bv">
                <el-input v-model="musicIns.bv" readonly />
                <div class="mmdci-icon">&#45;&#62;</div>
                <el-input v-model="newMusicIns.bv" />
            </el-form-item>

            <el-form-item class="mmd-change-input" label="description">
                <el-input v-model="musicIns.description" readonly />
                <div class="mmdci-icon">&#45;&#62;</div>
                <el-input v-model="newMusicIns.description" />
            </el-form-item>

            <el-form-item class="mmd-change-input" label="volume">
                <el-input v-model="musicIns.volume" readonly />
                <div class="mmdci-icon">&#45;&#62;</div>
                <el-input-number v-model="newMusicIns.volume" :min="-100" :max="100" :step="5" />
            </el-form-item>
        </el-form>

        <template #footer>
            <el-button type="info" plain @click="showModifyMusicDialog = false">取消</el-button>
            <el-button type="info" plain @click="modifyMusic()">确定</el-button>
        </template>
    </el-dialog>

    <el-dialog class="move-music-dialog" v-model="showMoveMusicDialog" title="移动歌曲">
        <el-form :model="musicIns" label-width="20%">
            <el-form-item label="选择播放列表">
                <el-select v-model="playlistIndexMoveTo" placeholder="请选择播放列表" @change="onChangePlaylist">
                    <el-option
                        v-for="(item, index) in listStore.list.playlists"
                        :key="item.id"
                        :label="item.name"
                        :value="index"
                    />
                </el-select>
            </el-form-item>
        </el-form>

        <template #footer>
            <el-button type="info" plain @click="showMoveMusicDialog = false">取消</el-button>
            <el-button type="info" plain @click="moveMusic()">确定</el-button>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { Music } from "@/axios/list.go";
import { v4 as uuid } from "uuid";
import { useListStore } from "@/pinia/list.ts";
import { log } from "@/ts/log.ts";
import { deepCopy } from "@/ts/utils.ts";

const listStore = useListStore();

let showAddMusicDialog = ref<boolean>(false);

let showModifyMusicDialog = ref<boolean>(false);
let musicIns = ref<Music>(new Music());
let newMusicIns = ref<Music>(new Music());

let showMoveMusicDialog = ref<boolean>(false);
let playlistIndexMoveTo = ref<number>(0);

function addMusic(): void {
    listStore.currentMusicList.push(musicIns.value);
    showAddMusicDialog.value = false;
}

function modifyMusic(): void {
    let index = getMusicIndex(newMusicIns.value.id);
    listStore.currentMusicList[index] = newMusicIns.value;

    showModifyMusicDialog.value = false;
}

function moveMusic(): void {
    // add to target playlist
    listStore.list.playlists[playlistIndexMoveTo.value].music_list.push(musicIns.value);

    // del from current playlist
    deleteMusic(musicIns.value.id);

    showMoveMusicDialog.value = false;
}

function deleteMusic(musicID: string): void {
    let index = getMusicIndex(musicID);
    if (index >= 0) {
        let lastIndex = listStore.currentMusicList.length - 1;

        let swap = listStore.currentMusicList[lastIndex];
        listStore.currentMusicList[lastIndex] = listStore.currentMusicList[index];
        listStore.currentMusicList[index] = swap;

        listStore.currentMusicList.pop();
    }
}

function openAddMusicDialog(): void {
    musicIns.value = new Music();
    musicIns.value.id = uuid();
    showAddMusicDialog.value = true;
}

function openModifyMusicDialog(musicID: string): void {
    let index = getMusicIndex(musicID);
    if (index < 0) {
        log.fail("invalid music id");
        return;
    }

    musicIns.value = listStore.currentMusicList[index];
    newMusicIns.value = deepCopy(musicIns.value);

    showModifyMusicDialog.value = true;
}

function openMoveMusicDialog(musicID: string): void {
    let index = getMusicIndex(musicID);
    if (index < 0) {
        log.fail("invalid music id");
        return;
    }

    musicIns.value = listStore.currentMusicList[index];

    showMoveMusicDialog.value = true;
}

function getMusicIndex(musicID: string): number {
    let index = -1;

    for (let i = 0; i < listStore.currentMusicList.length; i++) {
        if (listStore.currentMusicList[i].id === musicID) {
            index = i;
            break;
        }
    }

    return index;
}

function onChangePlaylist(value: number): void {
    playlistIndexMoveTo.value = value;
}
</script>

<style lang="less">
// modify style in body-level dialog, cancel 'scope' tag
.functions {
  margin-top: 2vh;
  margin-bottom: 5vh;

  .f-right {
    float: right;
  }
}

.add-music-dialog {
  .el-dialog__body {
    padding: 5rem 15%;

    .el-input {
      width: 80%;
    }

    .el-input-number {
      .el-input {
        width: 100%;
      }
    }
  }
}

.modify-music-dialog {
  .el-dialog__body {
    padding: 5rem 15%;

    .mmd-change-input {
      .el-input {
        width: 40%;
      }

      .el-input-number {
        .el-input {
          width: 100%;
        }
      }

      .mmdci-icon {
        width: 20%;
        text-align: center;
      }
    }
  }
}

.move-music-dialog {
  .el-dialog__body {
    padding: 5rem 15%;
  }
}
</style>
