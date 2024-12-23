import {RoomSerializer} from "@/types/serializer.ts";
import {useDeleteRoom} from "@/api/room.ts";
import {useCallback, useEffect, useState} from "react";
import {toast} from "@/hooks/use-toast.ts";
import {RiGroupLine, RiLoginBoxLine, RiMic2Line, RiSettings4Line, RiShareForwardLine} from "react-icons/ri";
import {TimeAgo} from "@/components/time.tsx";
import {
    Drawer,
    DrawerContent,
    DrawerDescription,
    DrawerFooter,
    DrawerHeader,
    DrawerTitle
} from "@/components/ui/drawer.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Loader2} from "lucide-react";
import {useNavigate} from "react-router";
import {useShareUrl} from "@/hooks/useShareUrl.ts";

export function RoomPreview({room}: { room: RoomSerializer }) {
    const {mutate: deleteRoom, isLoading, isError, error} = useDeleteRoom();
    const [showManageDrawer, setShowManageDrawer] = useState(false);
    const [showDeleteDrawer, setShowDeleteDrawer] = useState(false);
    const navigate = useNavigate();
    const shareUrl = useShareUrl(room.id);

    useEffect(() => {
        if (!isError)
            return;

        toast({
            title: "Something went wrong",
            description: <p className={'text-sm'}>{error?.message}</p>,
            variant: 'destructive'
        })
    }, [error?.message, isError]);

    const handleDeleteRoom = useCallback(() => {
        deleteRoom(room.id)
    }, [room, deleteRoom])

    const handleJoinRoom = useCallback(() => {
        navigate(`/meeting/${room.id}`)
    }, [navigate, room.id])

    return (
        <>
            <div className={'p-4 bg-tg-secondaryBg rounded flex items-center justify-between active:opacity-80'}
                 onClick={() => setShowManageDrawer(true)}>
                <h3 className={'flex items-center  justify-start gap-4'}>
                    <RiMic2Line size={28}/>
                    <div className={'flex flex-col items-start justify-center'}>
                        <p>{room.id}</p>
                        <p className={'text-muted text-xs'}>
                            <TimeAgo unixTimestamp={room.created_at}/>
                        </p>
                    </div>
                </h3>
                <div className={'flex items-center justify-start gap-3'}>
                    <Button variant={'outline'} size={'sm'} onClick={handleJoinRoom}>
                        <RiLoginBoxLine/>
                        Join
                    </Button>

                    <Button size={'sm'}>
                        <RiSettings4Line/>
                    </Button>
                </div>
            </div>
            <Drawer onClose={() => setShowManageDrawer(false)} onOpenChange={open => setShowManageDrawer(open)}
                    open={showManageDrawer}>
                <DrawerContent className={'border-none'}>
                    <DrawerHeader>
                        <DrawerTitle className={'text-tg-text'}>Manage {room.id}</DrawerTitle>
                        <DrawerDescription><strong>Caution!</strong> Deleted rooms are unrecoverable! If you delete this
                            room, all its data will be lost!</DrawerDescription>
                    </DrawerHeader>
                    <DrawerFooter>
                        <Button className={'w-full bg-transparent text-tg-text'}
                                disabled={isLoading} onClick={handleJoinRoom}><RiGroupLine/> Join Room</Button>

                        <Button className={'w-full bg-transparent text-tg-text'}
                                disabled={isLoading} onClick={() => window.open(shareUrl)}><RiShareForwardLine/> Share
                            Room</Button>

                        <Button variant={'destructive'} onClick={() => setShowDeleteDrawer(true)}
                                disabled={isLoading}>
                            {isLoading && <Loader2 className="animate-spin"/>} Delete Room
                        </Button>
                    </DrawerFooter>
                </DrawerContent>
            </Drawer>

            <Drawer onClose={() => setShowDeleteDrawer(false)} onOpenChange={open => setShowDeleteDrawer(open)}
                    open={showDeleteDrawer}>
                <DrawerContent className={'border-none'}>
                    <DrawerHeader>
                        <DrawerTitle className={'text-tg-text'}>Delete Room!</DrawerTitle>
                        <DrawerDescription>Are you absolutly sure you want to delete this room? All its data will be
                            gone and unrecoverable.</DrawerDescription>
                    </DrawerHeader>
                    <DrawerFooter>
                        <Button variant={'destructive'} onClick={handleDeleteRoom}
                                disabled={isLoading}>
                            {isLoading && <Loader2 className="animate-spin"/>} Delete {room.id}
                        </Button>
                        <Button className={'w-full bg-transparent text-tg-text'}
                                disabled={isLoading} onClick={() => setShowDeleteDrawer(false)}>Close</Button>
                    </DrawerFooter>
                </DrawerContent>
            </Drawer>


        </>
    )
}
