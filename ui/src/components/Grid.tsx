import {CSSProperties} from "react";

export interface RowProps {
    align: 'center' | 'flex-start' | 'flex-end' | 'stretch' | 'baseline';
    justify: 'center' | 'flex-start' | 'flex-end' | 'space-between' | 'space-around' | 'space-evenly';
    direction?: 'row' | 'column';
    gap?: string;
    lineHeight?: string;
    style?: CSSProperties;

    children: JSX.Element | (JSX.Element | false)[];
}

export function Row({
                        align,
                        justify,
                        direction = 'row',
                        gap = "1vh",
                        children,
                        lineHeight,
                        style
                    }: RowProps) {

    style = {
        "display": "flex",
        ...style,
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
