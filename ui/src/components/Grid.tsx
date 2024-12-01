import React, {CSSProperties} from "react";

export interface RowProps {
    align: 'center' | 'flex-start' | 'flex-end' | 'stretch' | 'baseline';
    justify: 'center' | 'flex-start' | 'flex-end' | 'space-between' | 'space-around' | 'space-evenly';
    direction?: 'row' | 'column';
    gap?: string;
    lineHeight?: string;

    children: JSX.Element | JSX.Element[];
}

export function Row({
                        align,
                        justify,
                        direction = 'row',
                        gap = "1vh",
                        children,
                        lineHeight,
                    }: RowProps) {

    const style: CSSProperties = {
        "display": "flex",
    };

    style["alignItems"] = align;
    style["justifyContent"] = justify;
    style["flexDirection"] = direction;
    style["gap"] = gap;

    if (lineHeight) {
        style["lineHeight"] = lineHeight;
    }

    return (
        <div style={style}>
            {children}
        </div>
    )
}

// export a row with forced center align and justify, and ability to modify other properties
export function CenterRow({children, ...props}: Omit<RowProps, 'align' | 'justify'>) {
    return <Row align={"center"} justify={"center"} {...props}>{children}</Row>
}
