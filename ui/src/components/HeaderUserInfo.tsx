import styles from "./HeaderUserInfo.module.css"
import {useCurrentUser} from "../hooks/user.ts";
import {Row} from "./Grid.tsx";
import {Button, Paragraph} from "./Typo.tsx";
import {HiBars2} from "react-icons/hi2";
import {FaBars, FaUsers} from "react-icons/fa";

export function HeaderUserInfo() {
    const currentUser = useCurrentUser();

    return <div className={styles.HeaderUserInfo}>
        <Row align={'center'} justify={'space-between'}>
            <Row align={'center'} justify={'center'}>
                <img src="/pixel-logo.jpg?v=1" alt="Logo"/>
                <Row align={'center'} justify={'center'} direction={'column'} gap={'0'}>
                    <h3>تصدانه</h3>
                    <Paragraph size={'s'} caption={true }>
                        Pixel Game
                    </Paragraph>
                </Row>
            </Row>
        </Row>

        <div className={styles.OnlineInfo}>
            <Row align={'center'} justify={'center'}>
                <FaUsers />
                <Paragraph size={'s'} caption={true}>
                    ۱۲ نفر آنلاین
                </Paragraph>
            </Row>

        </div>
    </div>
}