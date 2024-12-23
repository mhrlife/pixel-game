import * as React from "react";
import classNames from 'classnames';


interface PlaceholderProps extends React.HTMLAttributes<HTMLDivElement> {
    children: React.ReactNode
}

export function Placeholder({
                                children, ...rest
                            }: PlaceholderProps) {
    return (
        <div {...rest} className={classNames(rest.className, 'w-full  py-2 px-4')}>
            {children}
        </div>
    )
}