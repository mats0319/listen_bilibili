// Generate File, Should not Edit.
// Author: mario. https://github.com/mats9693
// Version: goc_ts v0.1.0

import { axiosWrapper } from "./config"
import { GetListRes, GetOriginURLRes, GetOriginURLReq } from "./list.go"
import { objectToFormData } from "./utils"

class ListAxios {
    public getList(): Promise<GetListRes> {
        return axiosWrapper.post("/list/get")
    }

    public getOriginURL(music_id: string): Promise<GetOriginURLRes> {
        let req: GetOriginURLReq = {
            music_id: music_id,
        }

        return axiosWrapper.post("/origin-url/get", objectToFormData(req))
    }
}

export const listAxios: ListAxios = new ListAxios()
