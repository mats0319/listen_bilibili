import { axiosWrapper } from "./config";
import { GetListRes, GetOriginURLReq, GetOriginURLRes } from "./list.go";
import { objectToFormData } from "./utils";

class ListAxios {
    public getList(): Promise<GetListRes> {
        return axiosWrapper.post("/list/get")
    }

    public getOriginURL(musicID: string): Promise<GetOriginURLRes> {
        let req: GetOriginURLReq = {
            music_id: musicID
        }

        return axiosWrapper.post("/originURL/get", objectToFormData(req))
    }
}

const listAxios = new ListAxios()
export { listAxios }
