import styles from "./Footer.module.css"
import {LuBookOpen, LuGift, LuPaintbrush2} from "react-icons/lu";
import {useLocation} from "react-router";

export function Footer() {
    return <div className={styles.Footer}>
        <FooterItem title={"جایزه"} icon={<LuGift/>} url={"/gift"}/>
        <FooterItem title={"نقاشی کن"} icon={<LuPaintbrush2/>} url={"/"}/>
        <FooterItem title={"آموزش"} icon={<LuBookOpen/>} url={"/learn"}/>
    </div>
}

function FooterItem({title, icon, url}: { title: string, icon: JSX.Element, url: string }) {
    const location = useLocation();

    return (
        <div className={styles.Item} data-active={location.pathname === url}>
            {icon}
            <p>{title}</p>
        </div>
    )
}