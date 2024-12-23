import {useMemo} from "react";


export const useShareUrl = (roomID: string) => {
    return useMemo(() => {
        const url = encodeURIComponent(`https://t.me/TONferenceBot/meeting?startapp=join-${roomID}`)
        const text = encodeURIComponent(`Join me in ${roomID} on TONference!`)
        return `https://t.me/share/url?url=${url}&text=${text}`
    }, [roomID])
}