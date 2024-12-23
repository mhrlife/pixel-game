import {Banner} from "@/components/Banner.tsx";
import {RiChatNewFill} from "react-icons/ri";
import {useCallback, useEffect} from "react";
import {InfinitySpin} from "react-loader-spinner";
import {toast} from "@/hooks/use-toast.ts";
import {useCreateRoom} from "@/api/room.ts";
import {MyRooms} from "@/components/MyRooms.tsx";
import {useNavigate} from "react-router";

export function Meetings() {
    const {isLoading, isError, error, isSuccess, data, mutate: createRoom} = useCreateRoom();
    const navigate = useNavigate();

    const handleIsCreating = useCallback(() => {
        createRoom();

    }, [createRoom]);

    useEffect(() => {
        if (!isError)
            return;
        toast({
            title: "Something went wrong",
            description: <p className={'text-sm'}>{error?.message}</p>,
            variant: 'destructive'
        })
    }, [error?.message, isError]);

    useEffect(() => {
        if (!isSuccess)
            return;

        navigate(`/meeting/${data?.id}`)
    }, [data?.id, isSuccess, navigate]);


    if (isLoading || isSuccess) {
        return (
            <div className={'py-10 flex flex-col items-center justify-center text-tg-hint'}>
                <InfinitySpin
                    width="200"
                    color="var(--tg-theme-text-color)"
                />
                <p className={'text-tg-text'}>
                    Creating a new meeting...
                </p>
            </div>
        )
    }


    return (
        <>

            <Banner title={"Create a Meeting!"} onClick={handleIsCreating} icon={<RiChatNewFill size={48}/>}
                    description={'Create a new meetings and share it in Telegram'}/>

            <div className={'py-5 text-sm text-tg-hint text-center px-2'}>
                <p>
                    Create a new meetings and share it with anyone you want in Telegram.
                </p>
            </div>

            <div className={'mb-5'}></div>

            <MyRooms/>
        </>
    )
}