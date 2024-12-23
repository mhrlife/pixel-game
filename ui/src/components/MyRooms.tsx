import {useMyRooms} from "@/api/room.ts";
import {Skeleton} from "@/components/ui/skeleton.tsx";
import {RoomPreview} from "@/components/RoomPreview.tsx";

export function MyRooms() {
    const {data: myRooms, isLoading, isSuccess, isRefetching} = useMyRooms();

    if (isLoading || isRefetching) {
        return <div className={'flex flex-col gap-1'}>
            <SkeletonRoom/>
            <SkeletonRoom/>
            <SkeletonRoom/>
        </div>
    }

    if (!isSuccess) {
        return <p>Something went wrong</p>
    }

    return (
        <div className={'flex flex-col gap-1'}>
            {myRooms.map((room) => <RoomPreview room={room} key={room.id}/>)}
        </div>
    )
}


function SkeletonRoom() {
    return (
        <Skeleton className="h-[8.5vh] w-full"/>
    )
}