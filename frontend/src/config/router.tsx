import {createBrowserRouter} from "react-router";
import LoginPage from "@pages/LoginPage";
import RegisterPage from "@pages/RegisterPage";
import ErrorPage from "@shared/components/ErrorBoundary/router.tsx";
import {ProtectedRoute} from "@shared/components/ProtectedRoute/ProtectedRoute.tsx";

const test = () => {
    return <ProtectedRoute>
        123
    </ProtectedRoute>
}

export const router = createBrowserRouter([
    {
        path: "/",
        errorElement: <ErrorPage/>,

        children: [
            {
                path: "/auth/login",
                Component: LoginPage,
            },
            {
                path: '/auth/register',
                Component: RegisterPage
            },
            {
                path: '/me',
                index: true,
                Component: test
            }
        ]
    },
])
