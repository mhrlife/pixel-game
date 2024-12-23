import {createBrowserRouter, RouterProvider} from "react-router";
import Layout from "./pages/Layout.tsx";
import {Meetings} from "@/pages/Meetings.tsx";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {Meeting} from "@/pages/Meeting.tsx";


const router = createBrowserRouter([
    {
        element: <Layout/>,
        path: "/",
        children: [
            {
                index: true,
                element: <Meetings/>
            },
            {
                path: "/meeting/:id",
                element: <Meeting/>
            }
        ]
    }
], {
    basename: import.meta.env.BASE_URL,
})

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            refetchOnWindowFocus: false,
        },
    },
});

function App() {
    return (
        <QueryClientProvider client={queryClient}>
            {/*<CentrifugeProvider url={"/pixel/events/connection/websocket"}>*/}
            <RouterProvider router={router}/>
            {/*</CentrifugeProvider>*/}
        </QueryClientProvider>
    )
}

export default App
