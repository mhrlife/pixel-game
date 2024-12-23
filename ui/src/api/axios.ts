import axios from "axios";
import {useAuth} from "@/store/useAuth.ts";

const axiosInstance = axios.create({
    baseURL: import.meta.env.BASE_URL + "/api",
})


axiosInstance.interceptors.request.use(
    (config) => {

        let token = useAuth.getState().user?.token || "undefined";

        if (!token || token === "undefined") {
            try {
                token = "INIT_DATA:" + getInitData()
            } catch {
                throw new Error("No token found");
            }
        } else {
            token = "JWT:" + token;
        }

        config.headers.setAuthorization(token)

        return config;
    },
)


function getInitData(): string {
    let initData;
    try {
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-expect-error
        initData = window.Telegram.WebApp.initData;
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
    } catch (e) { /* empty */ }

    if (initData) {
        return initData;
    }

    return "TEST_TOKEN"
}

export default axiosInstance;

