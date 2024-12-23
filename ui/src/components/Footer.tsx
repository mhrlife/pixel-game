import {Placeholder} from "@/components/ui/placeholder.tsx";
import {ReactNode} from "react";
import {RiBookOpenLine, RiSettings2Line, RiTelegram2Line} from "react-icons/ri";
import {Separator} from "@/components/ui/separator.tsx";
import classNames from "classnames";
import {useNavigate} from "react-router";


export function Footer() {
    const navigate = useNavigate();

    return <Placeholder className="bg-tg-bg py-6 flex items-center justify-center h-[10vh] gap-4">
        <FooterItem icon={<RiSettings2Line size={24}/>} label={"Settings"}/>
        <Separator orientation={'vertical'}/>
        <FooterItem icon={<RiTelegram2Line size={24}/>} label={"Meetings"} active={true} onClick={() => navigate("/")}/>
        <Separator orientation={'vertical'}/>
        <FooterItem icon={<RiBookOpenLine size={24}/>} label={"Help"}/>
    </Placeholder>
}


interface FooterItemProps extends React.HTMLAttributes<HTMLDivElement> {
    icon: ReactNode,
    label: string,
    active?: boolean
}

function FooterItem({icon, label, active = false, ...extras}: FooterItemProps) {
    return <div className={classNames({
        'flex flex-col items-center justify-center text-sm gap-0.5 py-1 px-2': true,
        'text-muted': !active,
    })} {...extras}>
        {icon}
        {label}
    </div>
}