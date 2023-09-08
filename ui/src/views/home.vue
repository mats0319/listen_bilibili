<template>
  <div class="home-top">
    <div class="ht-title" @click="link('/', 'home')">Listen Bilibili</div>

    <div class="ht-item">
      <a href="https://github.com/mats9693/listenBilibili" target="_blank">Github</a>
    </div>

  </div>

  <div class="home-content">
    <div class="h-left">
      <p class="hl-item" @click="link('/listen-bilibili', 'ListenBilibili')">
        Listen Bilibili
      </p>
    </div>

    <el-divider class="h-divider" direction="vertical" />

    <div class="h-right">
      <router-view />
    </div>
  </div>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
const router = useRouter()

function link(path: string, name: string) {
  if (location.href.endsWith(path)) {
    console.log("> Node: you are trying to link to current page.", path)
    return
  }

  router.push({name: name})
}
</script>

<style scoped lang="less">
.home-top {
  display: flex;
  height: 10rem;
  background-color: darkgray;

  .ht-title {
    width: 20vw;
    padding-left: 5vw;

    line-height: 10rem;
    font-size: 3rem;
  }
  .ht-title:hover {
    cursor: pointer;
  }

  .ht-item {
    width: 20vw;
    padding-left: 55vw;

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

  .h-left {
    width: 14vw;
    height: calc(100vh - 10rem);

    padding: 5vh 3vw 0;

    .hl-item {
      height: 4rem;
      line-height: 4rem;
      font-size: 2rem;
    }
    .hl-item:hover {
      cursor: pointer;
      background-color: aliceblue;
    }
  }

  .h-divider {
    height: calc(100vh - 10rem);
  }

  .h-right{
    padding: 5vh 3vw 0;
  }
}
</style>
