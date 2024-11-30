import styles from './Board.module.css'
import classNames from "classnames";
import {useCallback, useMemo, useState} from "react";
import {useDetectClickOutside} from "react-detect-click-outside";
import {EditForm} from "../components/EditForm.tsx";
import {Color, Colors, colorToHex} from "../types/colors.ts";

function randomColor(): Color {
    return Colors[Math.floor(Math.random() * Colors.length)];
}

export default function Board() {
    const [selected, setSelected] = useState<number | null>(null);

    const ref = useDetectClickOutside({
        onTriggered: () => setSelected(null)
    })

    const ids = useMemo(() => Array.from({length: 40 * 40}, (_, i) => {
        return {
            index: i,
            color: randomColor()
        }
    }), []);


    const handleOnClick = useCallback((id: number) => {
        if (selected === id)
            setSelected(null);
        else
            setSelected(id);
    }, [selected])

    const handleOnCancel = useCallback(() => {
        setSelected(null)
    }, [])

    return <div ref={ref}>
        <div className={styles.Board}>
            {ids.map(({index, color}) =>
                <BoardItem key={index}
                           color={color}
                           selected={selected === index}
                           onClick={() => handleOnClick(index)}/>)}

        </div>

        {selected && <EditForm onCancel={handleOnCancel} selected={selected}/>}
        {!selected && <div>as</div>}
    </div>

}

export function BoardItem({color, selected = false, onClick}: {
    color: Color;
    selected?: boolean;
    onClick?: () => void;

}) {
    const style = {
        backgroundColor: colorToHex(color)
    }

    return <div className={classNames(styles.BoardItem, {
        [styles.Selected]: selected
    })}
                style={style}
                onClick={onClick}
    ></div>
}