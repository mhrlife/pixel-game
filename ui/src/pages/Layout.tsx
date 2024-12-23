import {Outlet} from "react-router";
import {Footer} from "../components/Footer.tsx";
import {useAppDispatch} from "../store/store.ts";
import {loginUser} from "../store/user.ts";
import {useEffect} from "react";
import {useCurrentUser} from "../hooks/user.ts";
import {Grid} from "react-loader-spinner";
import styles from './Layout.module.css'
import {HeaderUserInfo} from "../components/HeaderUserInfo.tsx";

export default function Layout() {
    const dispatch = useAppDispatch();
    const currentUser = useCurrentUser();

    useEffect(() => {
        dispatch(loginUser());
    }, [dispatch]);

    if (!currentUser) {
        return <div className={styles.LoadingPage}>
            <Grid color={"#3B3030"} height={75} width={100}/>
        </div>
    }


    return (
        <div>
            <HeaderUserInfo/>
            <Outlet/>
            <Footer/>
        </div>
    )
}