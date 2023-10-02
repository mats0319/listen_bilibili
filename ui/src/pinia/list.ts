import { defineStore } from "pinia"
import { List, Music } from "@/axios/api.pb.ts"
import { ref } from "vue"
import { axiosInstance } from "@/axios/request.ts";
import { ElMessage } from "element-plus";

export let useListStore = defineStore("list", () => {
    let list = ref<List>({})
    let playlist = ref<Array<Music>>([])
    let playIndex = ref(0)

    let originURL = ref("")
    let volume = ref(0) // already like '0.3'

    function getList(): void {
        axiosInstance.getList()
            .then((response: any) => {
                if (response.data.err.length > 0) {
                    throw response.data.err
                }

                const listIns = response.data.list

                list.value = listIns
                playlist.value = listIns.playlists[0].music_list
                playIndex.value = 0

                playNextMusic()

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

    function playNextMusic(): void {
        axiosInstance.getOriginURL(playlist.value[playIndex.value].id!)
            .then((response: any) => {
                if (response.data.err.length > 0) {
                    throw response.data.err
                }

                originURL.value = response.data.url
                volume.value = playlist.value[playIndex.value].volume! / 100

                playIndex.value = (playIndex.value + 1) % playlist.value.length

                console.log("> Node: get origin url success.")
            })
            .catch((err: any) => {
                console.log("> Node: get origin url failed.", err)
                ElMessage({
                    message: "get origin url failed.",
                    type: "error",
                })

                setTimeout(() => {
                    playNextMusic() // auto re-try in 3s
                }, 3000)
            })
    }

    return {
        list,
        playlist,
        playIndex,
        originURL,
        volume,
        getList,
        playNextMusic,
    }
})
