import {createBrowserRouter, RouterProvider} from "react-router";
import Layout from "./pages/Layout.tsx";
import Board from "./pages/Board.tsx";
import {Provider} from "react-redux";
import {store} from "./store/store.ts";


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
])

function App() {


    return (
        <Provider store={store}>
            <RouterProvider router={router}>
            </RouterProvider>
        </Provider>
    )
}

export default App
