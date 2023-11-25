// Generate File, Should not Edit.
// Author: mario. https://github.com/mats9693
// Version: goc_ts v0.2.1

import axios, { AxiosInstance } from "axios";

export const axiosWrapper: AxiosInstance = axios.create({
    baseURL: "http://127.0.0.1:9693",
    timeout: 10000,
});
