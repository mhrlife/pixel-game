import { getInitData } from "../hooks/telegram.ts";
import { HTTPError } from "../store/types.ts";
import { BoardSerializer, UserWithToken, HypeSerializer } from "../types/serializer.ts";

async function call<T>(action: string, data: any) {
    let token = localStorage.getItem("pixel_jwt") || '';

    if (!token) {
        try {
            token = "INIT_DATA:" + getInitData()
        } catch {
            throw new Error("No token found");
        }
    } else {
        token = "JWT:" + getInitData();
    }

    const result = await fetch("/api/call", {
        method: "POST",
        body: JSON.stringify({
            action,
            ...data
        }),
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token,
        }
    })

    const response = await result.json();
    if (!response.ok) {
        throw response as HTTPError;
    }

    return response.data as T;
}

export function useApi() {
    return {
        async login() {
            return await call<UserWithToken>("users/login", {});
        },
        async getBoard() {
            return await call<BoardSerializer>("pixels/board", {});
        },
        async setPixel(id: number, color: string) {
            return await call<string>("pixels/update", { pixel_id: id, new_color: color });
        },
        async getHype() {
            return await call<HypeSerializer>("hype/count", {});
        },
        async getOnlineUsersCount() {
            return await call<number>("online_users/count", {});
        }
    }
}

