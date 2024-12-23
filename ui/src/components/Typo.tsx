import React, {CSSProperties} from "react";
import {LineWave, ThreeDots} from "react-loader-spinner";

export function Paragraph({
                              size = 'm',
                              caption = false,
                              children
                          }: {
    size?: 's' | 'm' | 'l';
    caption?: boolean;
    children: React.ReactNode;
}) {

    const style: CSSProperties = {};

    if (size === 's') {
        style.fontSize = '0.9rem';
    } else if (size === 'm') {
        style.fontSize = '1rem';
    } else if (size === 'l') {
        style.fontSize = '1.2rem';
    }

    if (caption) {
        style.opacity = 0.8;
    }

    return <p style={style}>{children}</p>
}


export function Button({
                           children,
                           onClick = () => {
                           },
                           type = 'button',
                           fullWidth = false,
                           visual = 'primary',
                           size = 'm',
                           loading = false,
                       }: {
    children: React.ReactNode;
    onClick?: () => void;
    type?: 'button' | 'submit' | 'reset';
    visual?: 'primary' | 'secondary' | 'danger' | 'border';
    size?: 's' | 'm' | 'l';
    loading?: boolean;
    fullWidth?: boolean;
}) {

    const style: CSSProperties = {
        padding: '0.5rem 1rem',
        border: 'none',
    }

    if (fullWidth) {
        style.width = '100%';
    }

    if (visual === 'primary') {
        style.backgroundColor = '#86ffd7';
        style.color = 'black';
    } else if (visual === 'secondary') {
        style.backgroundColor = '#FFF0D1';
        style.color = 'black';
    } else if (visual === 'danger') {
        style.backgroundColor = '#5a1111';
        style.color = 'white';
    } else if (visual === 'border') {
        style.backgroundColor = 'transparent';
        style.border = '1px solid #FFF0D1';
        style.color = '#FFF0D1';
    }

    style.borderRadius = '0.5rem';
    style.fontWeight = 'bold';

    if (size === 's') {
        style.fontSize = '0.9rem';
    } else if (size === 'm') {
        style.fontSize = '1rem';
    } else if (size === 'l') {
        style.fontSize = '1.2rem';
    }

    style.display = 'flex';
    style.alignItems = 'center';
    style.justifyContent = 'center';

    return <button
        style={style}
        onClick={() => {
            if (!loading)
                onClick();
        }}
        type={type}
    >
        {!loading && children}
        {loading && <ThreeDots color={'black'} width={30} height={30}/>}
    </button>
}