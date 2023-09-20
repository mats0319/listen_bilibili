<template>
    <div class="tool-bar">
        <div class="tb-volume">
            <el-input type="number" v-model="baseVolume" placeholder="base volume"/>
            <el-button type="info" plain @click="setVideoVolume">设置基础音量</el-button>
            <p>{{ volumeStr }}</p>
        </div>

        <el-button type="info" plain @click="listStore.playNextMusic">下一首</el-button>
    </div>

    <video id="listenBilibili" controls autoplay>
        <source type="video/mp4" src="">
    </video>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from "vue"
import { useListStore } from "@/pinia/list.ts";

const listStore = useListStore()

let baseVolume = ref(0.35)
let volume = ref(0.35)

onMounted(() => {
    if (!listStore.list.id) {
        listStore.getList()
    }

    (document.getElementById("listenBilibili") as HTMLVideoElement).onended = (_: Event): any => {
        listStore.playNextMusic()
    };
})

function setVideoVolume(): void {
    volume.value = baseVolume.value + listStore.volume;
    if (volume.value < 0) {
        volume.value = 0
    } else if (volume.value > 1) {
        volume.value = 1
    }

    (document.getElementById("listenBilibili") as HTMLVideoElement).volume = volume.value
}

const volumeStr = computed<string>(() => {
    return "当前视频音量：" + volume.value.toString()
})

watch(() => listStore.originURL, (newValue: string) => {
    (document.getElementById("listenBilibili") as HTMLVideoElement).src = newValue;
    setVideoVolume();
    (document.getElementById("listenBilibili") as HTMLVideoElement).play();
})
</script>

<style scoped lang="less">
.tool-bar {
    padding-bottom: 10rem;

    .tb-volume {
        display: flex;
        height: 4rem;
        padding-bottom: 2rem;

        .el-input {
            width: fit-content;
        }

        .el-button {
            height: inherit;
            margin-left: 3rem;
            margin-right: 3rem;
        }

        p {
            font-size: 1.4rem;
        }
    }
}
</style>
