import LoginForm from "../../features/login/ui/LoginForm.tsx";
import s from "./LoginPage.module.scss"

const LoginPage = () => {
    return (
        <div className={s.loginForm}>
            <LoginForm/>
        </div>
    )
}

export default LoginPage