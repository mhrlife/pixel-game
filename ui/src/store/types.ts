export type State = 'IDLE' | 'LOADING' | 'SUCCESS' | 'ERROR';

export type HTTPError = {
    message: string;
    error_code: number;
};

export type HTTPAction<T> = {
    state: State;
    error?: HTTPError;
    value?: T;
}

export function EmptyHTTPAction<T>(defaultState: State = 'IDLE'): HTTPAction<T> {
    return {
        state: defaultState
    }
}