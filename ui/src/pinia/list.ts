import { defineStore } from "pinia"
import { GetListRes, GetOriginURLRes, List, ModifyListRes, Music } from "@/axios/list.go.ts"
import { ref } from "vue"
import { listAxios } from "@/axios/list.http.ts";
import { log } from "@/ts/log.ts";

export let useListStore = defineStore("list", () => {
    let list = ref<List>(new List())
    // 一个有意思的地方：如果在vue文件里用类似`store.list.playlists[store.playlistIndex].music_list`这样的代码，
    // 无论直接写在html部分，或者写成计算属性——总之就是让vue自己监听和更新——是不行的，错误提示是：不能用一个未定义的内容点'music_list'。
    // 所以其他组件想要正常使用，需要提供一个【当前播放列表】参数出去，或者修改list的结构（去掉list一级）
    // 此外，如果所有位置对currentMusicList的修改都是使用类似`store.currentMusicList = store.list.playlist[index]`的代码，
    // 那么对store.list和store.currentMusicList上的修改实际上都是互通的
    let currentMusicList = ref<Array<Music>>([])
    let currentMusicID = ref<string>("") // record current music, used when 'play next music'

    function getList(reloadList: boolean = false): void {
        listAxios.getList(reloadList)
            .then(({ data }: { data: GetListRes }) => {
                if (data.err.length > 0) {
                    throw data.err
                }

                list.value = data.list
                currentMusicList.value = list.value.playlists[0].music_list

                log.success("get list success")
            })
            .catch((err: any) => {
                log.fail("get list failed", err)
            })
    }

    function modifyList(): void {
        listAxios.modifyList(JSON.stringify(list.value))
            .then(({ data }: { data: ModifyListRes }) => {
                if (data.err.length > 0) {
                    throw data.err
                }

                log.success("modify list success")
            })
            .catch((err: any) => {
                log.fail("modify list failed", err)
            })
    }

    // playMusic play music according to 'music id', or play next music without 'music id'
    function playMusic(musicID: string = "", cb: (url: string, name: string, volumeOffset: number) => void = () => {}): void {
        if (musicID === "") {
            musicID = currentMusicList.value[0].id;
            for (let i = 0; i < currentMusicList.value.length; i++) {
                if (currentMusicID.value === currentMusicList.value[i].id) {
                    let nextIndex = (i + 1) % currentMusicList.value.length
                    musicID = currentMusicList.value[nextIndex].id
                    break
                }
            }
        }

        currentMusicID.value = musicID;

        getOriginURL(musicID, cb);
    }

    function getOriginURL(
        musicID: string,
        cb: (url: string, name: string, volumeOffset: number) => void,
        retryTimes: number = 3
    ): void {
        listAxios.getOriginURL(musicID)
            .then(({ data }: { data: GetOriginURLRes }) => {
                if (data.err.length > 0) {
                    throw data.err
                }

                cb(data.url, data.name, data.volume/100);

                console.log("> Node: get origin url success.")
            })
            .catch((err: any) => {
                log.fail("get origin url failed", err)

                if (retryTimes > 1) {
                    setTimeout(() => {
                        getOriginURL(musicID, cb, retryTimes-1) // auto re-try in 3s
                    }, 3000)
                }
            })
    }

    // 为对象字段命名，在ide里可以点击跳转到外部使用该字段的位置；而匿名字段在ide里只能跳转到本文件上方、定义它的位置
    return {
        list: list,
        currentMusicList: currentMusicList,

        getList: getList,
        modifyList: modifyList,
        playMusic: playMusic,
    }
})
