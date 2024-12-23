/* Do not change, this code is generated from Golang structs */


export interface UserWithToken {
    id: string;
    display_name: string;
    token: string;
}
export interface User {
    id: string;
    display_name: string;
}
export interface RoomSerializer {
    id: string;
    max_participants: number;
    permission_video: string;
    permission_audio: string;
    created_at: number;
    updated_at: number;
}