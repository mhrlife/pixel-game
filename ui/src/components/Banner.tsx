import React from "react";
import classNames from "classnames";
import {RiArrowDropRightLine} from "react-icons/ri";

interface BannerProps extends React.HTMLAttributes<HTMLDivElement> {
    icon: React.ReactNode
    title: string
    description: string
    button?: string,
}

export function Banner({
                           icon, title, description, button,
                           ...extra
                       }: BannerProps) {

    return <div {...extra}
                className={classNames(extra.className, 'flex items-stretch bg-tg-secondaryBg rounded-xl px-6 py-4 shadow active:opacity-70')}>
        <div className={'w-1/4 flex items-center justify-center'}>
            {icon}
        </div>
        <div className={'w-3/4 flex flex-col justify-between items-start'}>
            <h3 className={'font-bold text-xl pb-1'}>{title}</h3>
            <p className={'text-sm text-muted'}>{description}</p>
            {button && <p className={'text-sm self-end flex gap-1 items-center text-tg-button'}>
                <RiArrowDropRightLine/> {button}</p>}
        </div>
    </div>
}