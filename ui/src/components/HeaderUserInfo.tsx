import styles from "./HeaderUserInfo.module.css"
import {Row} from "./Grid.tsx";
import {Paragraph} from "./Typo.tsx";
import {FaUsers} from "react-icons/fa";
import {fetchOnlineUsersCount} from "../store/stats.ts";
import {useAppDispatch, useAppSelector} from "../store/store.ts";
import {useEffect} from "react";
import {forceFarsiNumbers} from "../utils.ts";

export function HeaderUserInfo() {
    const dispatch = useAppDispatch();
    const updateOnlineUsers = () => dispatch(fetchOnlineUsersCount());
    const onlineUsers = useAppSelector(state => state.stats.onlineUsers);


    useEffect(() => {
        updateOnlineUsers()
    }, []);

    useEffect(() => {
        const interval = setInterval(() => {
            updateOnlineUsers()
        }, 60 * 1000)

        return () => clearInterval(interval);
    }, []);

    return <div className={styles.HeaderUserInfo}>
        <Row align={'center'} justify={'space-between'}>
            <Row align={'center'} justify={'center'}>
                <img src="/pixel-logo.jpg?v=1" alt="Logo"/>
                <Row align={'center'} justify={'center'} direction={'column'} gap={'0'}>
                    <h3>تصدانه</h3>
                    <Paragraph size={'s'} caption={true}>
                        Pixel Game
                    </Paragraph>
                </Row>
            </Row>
        </Row>

        <div className={styles.OnlineInfo}>
            <Row align={'center'} justify={'center'}>
                <FaUsers/>
                {onlineUsers.state === 'SUCCESS' && (
                    <Paragraph caption={true}>کاربران
                        آنلاین: {forceFarsiNumbers(onlineUsers.value || 0)}</Paragraph>
                )}

                {onlineUsers.state === 'LOADING' && (
                    <Paragraph caption={true}>در حال دریافت تعداد کاربران آنلاین...</Paragraph>
                )}

                {onlineUsers.state === 'ERROR' && (
                    <Paragraph caption={true}>{onlineUsers.error?.message}</Paragraph>
                )}
            </Row>

        </div>
    </div>
}