import {cn} from "@/lib/utils"

function Skeleton({
                      className,
                      ...props
                  }: React.HTMLAttributes<HTMLDivElement>) {
    return (
        <div
            className={cn("animate-pulse rounded-md bg-tg-secondaryBg", className)}
            {...props}
        />
    )
}

export {Skeleton}
