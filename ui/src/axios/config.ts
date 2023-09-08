import axios, { AxiosInstance } from "axios";

export const axiosWrapper: AxiosInstance = axios.create({
  baseURL: "http://127.0.0.1:9693",
  timeout: 3000
});
