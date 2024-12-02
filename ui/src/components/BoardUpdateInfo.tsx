import styles from './BoardUpdateInfo.module.css';
import {Paragraph} from "./Typo.tsx";
import {UpdatedBoardSerializer} from "../types/serializer.ts";


export type BoardUpdateInfoProps = {
    boardUpdate: UpdatedBoardSerializer | null;
}

export function BoardUpdateInfo({boardUpdate}: BoardUpdateInfoProps) {
    if (!boardUpdate) {
        return <div className={styles.BoardUpdateHolder}>
            <h3>آخرین تغییرات</h3>
            <Paragraph>
                روی تصدانه‌ای که می‌خواهید ویرایش کنید کلیک کنید.
            </Paragraph>
        </div>
    }

    return (
        <div className={styles.BoardUpdateHolder}>
            <h3>آخرین تغییرات</h3>
            <Paragraph>
                یک تصدانه توسط {boardUpdate.user.display_name} ویرایش شد.
            </Paragraph>
        </div>
    )
}