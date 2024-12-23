import {Outlet, useNavigate} from "react-router";
import {useEffect} from "react";
import {useAuth} from "@/store/useAuth.ts";
import {ScrollArea} from "@/components/ui/scroll-area.tsx";
import {Placeholder} from "@/components/ui/placeholder.tsx";
import {Header} from "@/components/Header.tsx";
import {Footer} from "@/components/Footer.tsx";
import {Toaster} from "@/components/ui/toaster.tsx";
import {InfinitySpin} from "react-loader-spinner";

export default function Layout() {
    const {login, user, error: authError} = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        login()
    }, [login])

    useEffect(() => {
        if (!user)
            return;

        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-expect-error
        const initData: string | null = window.Telegram.WebApp?.initDataUnsafe?.start_param;
        if (!initData) {
            return;
        }

        if (initData.startsWith("join-")) {
            const roomId = initData.replace("join-", "");
            navigate(`/meeting/${roomId}`);
        }
    }, [navigate, user]);


    if (authError) {
        return <p>{authError.message}</p>
    }

    if (!user) {
        return <div className="h-[100vh] flex flex-col min-h-screen bg-tg-bg shrink-0 text-tg-text">
            <Header/>

            <Placeholder className={'h-[80vh]'}>
                <div className={'py-10 flex flex-col items-center justify-center text-tg-hint'}>
                    <InfinitySpin
                        width="200"
                        color="var(--tg-theme-text-color)"
                    />
                    <p className={'text-tg-text'}>
                        Connecting to Telegram...
                    </p>
                </div>
            </Placeholder>

            <Footer/>
            <Toaster/>
        </div>
    }


    return (
        <div className="h-[100vh] flex flex-col min-h-screen bg-tg-bg shrink-0 text-tg-text">
            <Header/>

            <Placeholder className={'h-[80vh]'}>
                <ScrollArea className="h-full">
                    <Outlet/>
                </ScrollArea>
            </Placeholder>

            <Footer/>
            <Toaster/>
        </div>
    )
}