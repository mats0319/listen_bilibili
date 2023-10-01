import { axiosWrapper } from "@/axios/config";
import { GetListRes, GetOriginURLReq, GetOriginURLRes } from "@/axios/api.pb";
import { objectToFormData } from "@/axios/utils.ts";

class AxiosInstance {
    public getList(): Promise<GetListRes> {
        return axiosWrapper.post("/list/get")
    }

    public getOriginURL(musicID: string): Promise<GetOriginURLRes> {
        let req: GetOriginURLReq = {
            music_id: musicID
        }

        return axiosWrapper.post("/origin-url/get", objectToFormData(req))
    }
}

const axiosInstance = new AxiosInstance()
export { axiosInstance }
