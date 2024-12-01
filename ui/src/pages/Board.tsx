// ui/src/pages/Board.tsx
import styles from './Board.module.css'
import classNames from "classnames";
import {useCallback, useEffect, useState} from "react";
import {useDetectClickOutside} from "react-detect-click-outside";
import {EditForm} from "../components/EditForm.tsx";
import {Color, colorToHex} from "../types/colors.ts";
import {useApi} from "../api/useApi.tsx";
import {BoardSerializer} from "../types/serializer.ts";
import {fetchUserHype} from "../store/user.ts";
import {useAppDispatch, useAppSelector} from "../store/store.ts";
import {CenterRow, Row} from "../components/Grid.tsx";
import {useCurrentUser} from "../hooks/user.ts";
import {Paragraph} from "../components/Typo.tsx";
import {forceFarsiNumbers} from "../utils.ts";
import {FaFire} from "react-icons/fa";

export default function Board() {
    const [selected, setSelected] = useState<number | null>(null);
    const api = useApi();
    const [board, setBoard] = useState<BoardSerializer | null>(null);
    const currentUser = useCurrentUser();

    const [isLoading, setIsLoading] = useState(false);
    const [countdown, setCountdown] = useState<number | null>(null);

    const ref = useDetectClickOutside({
        onTriggered: () => setSelected(null)
    })

    const dispatch = useAppDispatch();
    const hype = useAppSelector(state => state.user.hype);

    const updateBoard = () => api.getBoard().then(setBoard)
    const updateHype = () => dispatch(fetchUserHype());

    useEffect(() => {
        updateBoard()
        updateHype();
    }, [dispatch]);

    useEffect(() => {
        if (hype.state === 'SUCCESS' && hype.value) {
            if (hype.value.amount_remaining < hype.value.max_hype && hype.value.time_until_next_hype > 0) {
                setCountdown(hype.value.time_until_next_hype);
            } else {
                setCountdown(null);
            }
        }
    }, [hype]);

    useEffect(() => {
        if (countdown === null) return;
        if (countdown <= 0) {
            updateHype();
            return;
        }

        const timer = setInterval(() => {
            setCountdown(prev => {
                if (prev && prev > 0) {
                    return prev - 1;
                }
                return null;
            });
        }, 1000);
        return () => clearInterval(timer);
    }, [countdown, updateHype]);

    const handleSetPixel = useCallback((color: Color) => {
        if (!selected) return;

        setIsLoading(true)
        api.setPixel(selected, color).then(() => updateBoard()).finally(() => {
            setIsLoading(false)
            setSelected(null)
            dispatch(fetchUserHype())
        })
    }, [selected, updateBoard, dispatch])

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
        <div className={styles.HypeInfo}>

            <Row align={'center'} justify={'space-between'}>
                <>
                    <Row align={'flex-start'} justify={'flex-start'} direction={'column'} lineHeight={"0.5rem"}>
                        <h4>{currentUser?.display_name}</h4>
                        <Paragraph caption={true}>خوش آمدید</Paragraph>
                    </Row>

                    {hype.state === 'SUCCESS' && hype.value && (
                        <Row align={'center'} justify={'center'} gap={'0.5vh'}>
                            {hype.state === 'SUCCESS' && hype.value && hype.value.amount_remaining < hype.value.max_hype && countdown !== null && (
                                <CenterRow>
                                    <Paragraph caption={true} size={'s'}>(۰۰:{forceFarsiNumbers(countdown)})</Paragraph>
                                </CenterRow>
                            )}

                            <CenterRow gap={'0.3vh'}>
                                <h4>{forceFarsiNumbers(hype.value.amount_remaining)}</h4>
                                <FaFire/>
                            </CenterRow>

                            <CenterRow gap={'0.3vh'}>
                                <Paragraph caption={true}>از</Paragraph>
                                <Paragraph caption={true}>{forceFarsiNumbers(hype.value.max_hype)}</Paragraph>
                            </CenterRow>
                        </Row>
                    )}

                    {hype.state === 'LOADING' && (
                        <h4>در حال بارگذاری...</h4>
                    )}

                </>

            </Row>


        </div>
        <div className={styles.Board}>
            {board && board.pixels.map((pixel) =>
                <BoardItem key={pixel.id}
                           color={pixel.color as Color}
                           selected={selected === pixel.id}
                           onClick={() => handleOnClick(pixel.id)}/>
            )}
        </div>

        {selected && <EditForm onCancel={handleOnCancel} onSubmitted={handleSetPixel} isLoading={isLoading}
                               selected={selected}/>}
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
