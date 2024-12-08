import {createBrowserRouter, RouterProvider} from "react-router";
import Layout from "./pages/Layout.tsx";
import Board from "./pages/Board.tsx";
import {Provider} from "react-redux";
import {store} from "./store/store.ts";
import {CentrifugeProvider} from "./context/centrifuge.tsx";


const router = createBrowserRouter([
    {
        element: <Layout/>,
        path: "/",
        children: [
            {
                index: true,
                element: <Board/>
            }
        ]
    }
], {
    basename: import.meta.env.BASE_URL,
})

function App() {


    return (
        <Provider store={store}>
            <CentrifugeProvider url={"/pixel/events/connection/websocket"}>
                <RouterProvider router={router}/>
            </CentrifugeProvider>
        </Provider>
    )
}

export default App
