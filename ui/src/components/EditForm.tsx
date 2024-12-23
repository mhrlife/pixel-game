import styles from "./Editform.module.css";
import {Row} from "./Grid.tsx";
import {AiOutlineClose} from "react-icons/ai";
import {Color, Colors, colorToHex} from "../types/colors.ts";
import {useEffect, useState} from "react";
import {Button} from "./Typo.tsx";

export function EditForm({selected, onCancel, onSubmitted, isLoading}: {
    selected: number;
    onCancel: () => void;
    onSubmitted: (color: Color) => void;
    isLoading: boolean;
}) {

    const [selectedColor, setSelectedColor] = useState<Color | null>(null);

    useEffect(() => {
        setSelectedColor(null);
    }, [selected]);

    return <div className={styles.EditForm}>
        <Row align={'center'} justify={'space-between'}>
            <h3>تغییر رنگ به</h3>
            <div onClick={onCancel} style={{padding: "0.5vh"}}>
                <AiOutlineClose/>
            </div>
        </Row>
        <div className={styles.ColorsHolder}>
            {Colors.map(color => <div key={color}
                                      data-selected={selectedColor === color}
                                      onClick={() => setSelectedColor(color)}
                                      style={{backgroundColor: colorToHex(color)}}/>)}
        </div>
        {selectedColor &&
            <Button onClick={() => onSubmitted(selectedColor)} loading={isLoading} fullWidth={true} size={'m'}>تایید و
                رنگ زدن</Button>}
    </div>
}