import s from "./RegisterPage.module.scss"
import RegisterForm from "@features/register/ui/RegisterForm.tsx";

const RegisterPage = () => {
    return (
        <div className={s.registerForm}>
            <RegisterForm/>
        </div>
    )
}

export default RegisterPage