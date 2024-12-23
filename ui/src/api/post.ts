import axiosInstance from "./axios.ts";
import {AxiosResponse} from "axios";


export interface ApiResponse<Response> {
    ok: true;
    data: Response;
}

export interface ApiError {
    ok: false;
    message: string;
    code: number;
}

export async function post<Request, Response>(method: string, args: Request): Promise<Response> {
    const axiosResponse: AxiosResponse<ApiResponse<Response> | ApiError> = await axiosInstance.post("/call", {
        ...args,
        "action": method,
    })

    if (!axiosResponse.data.ok) {
        throw axiosResponse.data as ApiError;
    }

    return axiosResponse.data.data as Response;
}

export async function get<Response>(url: string): Promise<Response> {
    const axiosResponse: AxiosResponse<ApiResponse<Response> | ApiError> = await axiosInstance.get(url)

    if (!axiosResponse.data.ok) {
        throw axiosResponse.data as ApiError;
    }

    return axiosResponse.data.data as Response;
}
