import { defineStore } from "pinia"
import { GetListRes, GetOriginURLRes, List, ModifyListRes, Music } from "@/axios/list.go.ts"
import { ref } from "vue"
import { listAxios } from "@/axios/list.http.ts";
import { ElMessage } from "element-plus";

export let useListStore = defineStore("list", () => {
    let list = ref<List>(new List())
    let playlist = ref<Array<Music>>([])

    // data for current playing music
    let musicID = ref("")
    let musicName = ref("")
    let originURL = ref("")
    let volume = ref(0) // already like '0.3'

    function getList(reloadList: boolean = false): void {
        listAxios.getList(reloadList)
            .then(({ data }: { data: GetListRes }) => {
                if (data.err.length > 0) {
                    throw data.err
                }

                const listIns = data.list

                list.value = listIns
                playlist.value = listIns.playlists![0].music_list as Array<Music>

                console.log("> Node: get list success.")
                ElMessage({
                    message: "get list success.",
                    type: "success",
                })
            })
            .catch((err: any) => {
                console.log("> Node: get list failed.", err)
                ElMessage({
                    message: "get list failed.",
                    type: "error",
                })
            })
    }

    function modifyList(): void {
        listAxios.modifyList(JSON.stringify(list.value))
            .then(({ data }: { data: ModifyListRes }) => {
                if (data.err.length > 0) {
                    throw data.err
                }

                console.log("> Node: modify list success.")
                ElMessage({
                    message: "modify list success.",
                    type: "success",
                })
            })
            .catch((err: any) => {
                console.log("> Node: modify list failed.", err)
                ElMessage({
                    message: "modify list failed.",
                    type: "error",
                })
            })
    }

    function playNextMusic(): void {
        let nextMusic: Music = playlist.value[0]
        for (let i = 0; i < playlist.value.length; i++) {
            if (musicID.value === playlist.value[i].id) {
                let nextIndex = (i + 1) % playlist.value.length
                nextMusic = playlist.value[nextIndex]
                break
            }
        }

        musicID.value = nextMusic.id as string

        getOriginURL()
    }

    // getOriginURL play music in current playlist, auto re-try when error
    function getOriginURL(retryTimes: number = 3): void {
        listAxios.getOriginURL(musicID.value)
            .then(({ data }: { data: GetOriginURLRes }) => {
                if (data.err.length > 0) {
                    throw data.err
                }

                originURL.value = data.url
                volume.value = data.volume / 100

                // get music name
                for (let i = 0; i < playlist.value.length; i++) {
                    if (musicID.value === playlist.value[i].id) {
                        musicName.value = playlist.value[i].name as string
                        break
                    }
                }

                console.log("> Node: get origin url success.")
            })
            .catch((err: any) => {
                console.log("> Node: get origin url failed.", err)
                ElMessage({
                    message: "get origin url failed.",
                    type: "error",
                })

                if (retryTimes > 1) {
                    setTimeout(() => {
                        getOriginURL(retryTimes-1) // auto re-try in 3s
                    }, 3000)
                }
            })
    }

    return {
        list,
        playlist,

        musicID,
        musicName,
        originURL,
        volume,

        getList,
        modifyList,
        playNextMusic,
        getOriginURL,
    }
}/*, {
    persist: true,
}*/)
