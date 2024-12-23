import {create} from "zustand";
import {UserWithToken} from "@/types/serializer.ts";
import {ApiError, post} from "@/api/post.ts";


interface AuthState {
    user: UserWithToken | null;
    error: ApiError | null;
    login: () => void
}

export const useAuth = create<AuthState>((set) => ({
    user: null,
    error: null,
    login: async () => {
        try {
            const response = await post<any, UserWithToken>("users/login", {})
            set({user: response, error: null})
        } catch (e) {
            set({error: e as ApiError})
        }
    },
}));