import React, {createContext, useContext, useEffect, useMemo, useState} from 'react';
import {Centrifuge, ServerPublicationContext} from 'centrifuge';
import {useAppSelector} from '../store/store';

export interface CentrifugeContextValue {
    centrifuge: Centrifuge | null;
    latestUpdate: ServerEvent<any> | null;
}

const CentrifugeContext = createContext<CentrifugeContextValue | null>(null);

interface CentrifugeProviderProps {
    url: string;
    children: React.ReactNode;
}

interface ServerEvent<T> {
    data: T;
    event: string;
}

export const CentrifugeProvider: React.FC<CentrifugeProviderProps> = ({url, children}) => {
    const jwtToken = useAppSelector(state => state.user?.auth?.value?.token);

    const [latestUpdate, setLatestUpdate] = useState<ServerEvent<any> | null>(null);

    const centrifuge = useMemo(() => {
        if (!jwtToken) return null;

        return new Centrifuge(url, {
            getToken: async () => jwtToken,
        });
    }, [url, jwtToken]);

    useEffect(() => {
        if (!centrifuge) return;

        centrifuge.on('connected', () => {
            console.log("connected to websocket.");
        });

        centrifuge.on('publication', (ctx: ServerPublicationContext) => {
            setLatestUpdate(ctx.data as ServerEvent<any>);
        });

        centrifuge.connect();

        return () => {
            console.log("disconnected from websocket.")
            centrifuge.disconnect();
        };
    }, [centrifuge]);

    return (
        <CentrifugeContext.Provider value={{
            centrifuge,
            latestUpdate,
        }}>
            {children}
        </CentrifugeContext.Provider>
    );
};


export function useSubscription<T>(event: string): T | null {
    const context = useContext(CentrifugeContext);

    const [data, setData] = useState<T | null>(null);

    useEffect(() => {
        if (!context?.latestUpdate) return;

        if (context.latestUpdate.event !== event) return;

        setData(context.latestUpdate.data as T);

        return () => {
            setData(null);
        };
    }, [event, context?.latestUpdate]);

    return data;
}