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
export interface PixelSerializer {
    id: number;
    color: string;
    updated_at: string;
}
export interface PixelWithUserSerializer {
    id: number;
    color: string;
    user?: User;
    updated_at: string;
}
export interface BoardSerializer {
    pixels: PixelWithUserSerializer[];
    width: number;
    height: number;
}
export interface HypeSerializer {
    amount_remaining: number;
    max_hype: number;
    hype_per_second: number;
    time_until_next_hype: number;
    last_updated_at: string;
}