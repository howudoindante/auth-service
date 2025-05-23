import {ReactNode} from "react";
import {Navigate} from "react-router";

export const ProtectedRoute = ({children}: { children: ReactNode }) => {
    const cookies = Object.fromEntries(
        document.cookie.split('; ').map(cookie => cookie.split('='))
    );

    console.log(document.cookie)

    if (!cookies.access_token) {
        return <Navigate to="/auth/login" replace/>;
    }
    return children
}