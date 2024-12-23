import {useMemo} from "react";
import {UserWithToken} from "../types/serializer.ts";
import {useAppSelector} from "../store/store.ts";


export function useCurrentUser(): UserWithToken | null {
    const auth = useAppSelector(state => state.user.auth);

    return useMemo(() => {
        if (auth.state === 'SUCCESS') {
            return auth.value as UserWithToken;
        }

        return null;
    }, [auth])
}