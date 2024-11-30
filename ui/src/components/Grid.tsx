import React, {CSSProperties} from "react";

export function Row({
                        align,
                        justify,
                        direction = 'row',
                        gap = "1vh",
                        children
                    }: {
    align: 'center' | 'flex-start' | 'flex-end' | 'stretch' | 'baseline';
    justify: 'center' | 'flex-start' | 'flex-end' | 'space-between' | 'space-around' | 'space-evenly';
    direction?: 'row' | 'column';
    gap?: string;

    children: JSX.Element | JSX.Element[];
}) {

    const style: CSSProperties = {
        "display": "flex",
    };

    style["alignItems"] = align;
    style["justifyContent"] = justify;
    style["flexDirection"] = direction;
    style["gap"] = gap;

    return (
        <div style={style}>
            {children}
        </div>
    )
}