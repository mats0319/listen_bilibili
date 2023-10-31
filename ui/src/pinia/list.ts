import { defineStore } from "pinia"
import { List, Music } from "@/axios/list.go.ts"
import { ref } from "vue"
import { listAxios } from "@/axios/list.http.ts";
import { ElMessage } from "element-plus";

export let useListStore = defineStore("list", () => {
    let list = ref<List>({})
    let playlist = ref<Array<Music>>([])

    // data for current playing music
    let musicID = ref("")
    let musicName = ref("")
    let originURL = ref("")
    let volume = ref(0) // already like '0.3'

    function getList(): void {
        listAxios.getList()
            .then((response: any) => {
                if (response.data.err.length > 0) {
                    throw response.data.err
                }

                const listIns = response.data.list

                list.value = listIns
                playlist.value = listIns.playlists[0].music_list

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
        let nextMusic = getNextMusic()
        musicID.value = nextMusic.id as string

        playMusic()
    }

    function playMusic(): void {
        if (musicID.value.length < 1) {
            let nextMusic = getNextMusic()
            musicID.value = nextMusic.id as string
        }

        listAxios.getOriginURL(musicID.value)
            .then((response: any) => {
                if (response.data.err.length > 0) {
                    throw response.data.err
                }

                originURL.value = response.data.url
                volume.value = response.data.volume / 100

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

                setTimeout(() => {
                    playMusic() // auto re-try in 3s
                }, 3000)
            })
    }

    function getNextMusic(): Music {
        let nextMusic: Music = playlist.value[0]
        for (let i = 0; i < playlist.value.length; i++) {
            if (musicID.value === playlist.value[i].id) {
                let index = (i + 1) % playlist.value.length
                nextMusic = playlist.value[index]
                break
            }
        }

        return nextMusic
    }

    return {
        list,
        playlist,

        musicID,
        musicName,
        originURL,
        volume,

        getList,
        playNextMusic,
        playMusic,
    }
}/*, {
    persist: true,
}*/)
