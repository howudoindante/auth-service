import './App.css'
import {router} from "@config/router.tsx";
import {RouterProvider} from "react-router";


function App() {

    return (
        <>
            <RouterProvider router={router}/>

        </>
    )
}

export default App
