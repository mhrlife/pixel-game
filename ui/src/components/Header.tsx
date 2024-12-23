import {Avatar, AvatarFallback} from "@/components/ui/avatar.tsx";
import {Placeholder} from "@/components/ui/placeholder.tsx";
import {useAuth} from "@/store/useAuth.ts";
import {SiTon} from "react-icons/si";

export function Header() {

    return <Placeholder className="flex items-center justify-between bg-tg-bg p-4 h-[10vh]">
        <div className={'flex items-center text-lg gap-1.5'}>
            <span className={'text-2xl'}><SiTon/></span>
            <span className={'font-bold'}>
                TON <span className={'text-muted font-normal'}>ference</span>
            </span>
        </div>
        <UserAvatar/>
    </Placeholder>
}

function UserAvatar() {
    const {user} = useAuth();

    if (!user) {
        return null;
    }

    return <div className="flex items-center gap-2">
        <Avatar>
            <AvatarFallback className="bg-tg-secondaryBg text-tg-text">
                {user.display_name.charAt(0)}
            </AvatarFallback>
        </Avatar>
        <p className="text-sm">{user.display_name}</p>
    </div>
}