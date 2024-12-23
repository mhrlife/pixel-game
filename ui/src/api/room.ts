import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query";
import {RoomSerializer} from "@/types/serializer.ts";
import {ApiError, post} from "@/api/post.ts";


export const useCreateRoom = () => {
    const queryClient = useQueryClient();
    return useMutation<RoomSerializer, ApiError>({
        mutationFn: async () => {
            return await post<object, RoomSerializer>("rooms/create", {});
        },
        onSuccess: () => {
            queryClient.invalidateQueries(['my-rooms']);
        },
    });
};

export const useDeleteRoom = () => {
    const queryClient = useQueryClient();
    return useMutation<string, ApiError, string>({
        mutationFn: async (id: string) => {
            return await post<object, string>(`rooms/delete`, {id: id});
        },
        onSuccess: () => {
            queryClient.invalidateQueries(['my-rooms']);
        }
    });
};

export const useMyRooms = () => {
    return useQuery<RoomSerializer[], ApiError>({
        queryKey: ['my-rooms'],
        queryFn: async () => {
            return await post<object, RoomSerializer[]>("rooms/mine", {});
        }
    });
};

export const useRoomToken = (roomID: string) => {
    return useQuery<string, ApiError>({
        queryKey: ['room-token', roomID],
        queryFn: async () => {
            return await post<object, string>("rooms/token", {id: roomID});
        },
        enabled: !!roomID
    });
};
