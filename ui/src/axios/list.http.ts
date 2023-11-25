// Generate File, Should not Edit.
// Author: mario. https://github.com/mats9693
// Version: goc_ts v0.2.1

import { axiosWrapper } from "./config"
import { AxiosResponse } from "axios"
import { GetListRes, GetListReq, GetOriginURLRes, GetOriginURLReq, ModifyListRes, ModifyListReq } from "./list.go"
import { objectToFormData } from "./utils"

class ListAxios {
    public getList(reload_list: boolean): Promise<AxiosResponse<GetListRes>> {
        let req: GetListReq = {
            reload_list: reload_list,
        }

        return axiosWrapper.post("/list/get", objectToFormData(req))
    }

    public getOriginURL(music_id: string): Promise<AxiosResponse<GetOriginURLRes>> {
        let req: GetOriginURLReq = {
            music_id: music_id,
        }

        return axiosWrapper.post("/origin-url/get", objectToFormData(req))
    }

    public modifyList(list: string): Promise<AxiosResponse<ModifyListRes>> {
        let req: ModifyListReq = {
            list: list,
        }

        return axiosWrapper.post("/list/modify", objectToFormData(req))
    }
}

export const listAxios: ListAxios = new ListAxios()
