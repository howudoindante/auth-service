import {Controller, useForm} from "react-hook-form";
import {yupResolver} from "@hookform/resolvers/yup";

import {Button, Field, Heading, Input, Link as ChakraLink, Separator, Text} from "@chakra-ui/react";
import {PasswordInput} from "@components/ui/password-input.tsx";
import s from "./RegisterForm.module.scss"
import {observer} from "mobx-react-lite";
import authStore from "@entities/auth/model/store.ts";
import {registerSchema} from "@entities/auth/model/validation.ts";
import Layout from "@shared/components/Layout/Layout.tsx";
import {Logo} from "@shared/components/Logo/Logo.tsx";
import {Link} from "react-router"

const RegisterForm = observer(() => {
    const {register} = authStore;
    const {control, handleSubmit} = useForm({
        resolver: yupResolver(registerSchema),
    })
    return (
        <div className={s.container}>
            <Layout/>
            <form onSubmit={handleSubmit((v) => register(v))} className={s.form}>

                <Logo className={s.logo}/>
                <Heading className={s.heading}>Register account</Heading>
                <Controller control={control} render={
                    ({field, fieldState: {error},}) =>
                        <Field.Root required invalid={!!error} className={s.username}>

                            <Input placeholder="Email" value={field.value} onChange={field.onChange}
                                   size="lg"/>
                            {error && <Field.ErrorText>{error.message}</Field.ErrorText>}
                        </Field.Root>
                } name="email"/>
                <Controller control={control} render={
                    ({field, fieldState: {error},}) =>
                        <Field.Root required invalid={!!error} className={s.username}>


                            <Input placeholder="Username" value={field.value} onChange={field.onChange}
                                   size="lg"/>
                            {error && <Field.ErrorText>{error.message}</Field.ErrorText>}
                        </Field.Root>
                } name="username"/>
                <Controller control={control} render={
                    ({field, fieldState: {error},}) =>
                        <Field.Root required invalid={!!error} className={s.password}>


                            <PasswordInput placeholder="Password" value={field.value}
                                           onChange={field.onChange}
                                           size="lg"/>
                            {error && <Field.ErrorText>{error.message}</Field.ErrorText>}
                        </Field.Root>
                } name="password"/>

                <Button className={s.btn} type="submit">Register account</Button>
                <div className={`${s.row} ${s.centeredContent}`}>
                    <Text className={s.link}>Already have account?
                        <ChakraLink
                            asChild
                            variant="underline"
                            colorPalette="teal">
                            <Link to="/auth/login"> Let's login!
                            </Link>
                        </ChakraLink>
                    </Text>
                </div>
                <div className={s.separator}>
                    <Separator flex="1"/>
                    <Text flexShrink="0">OR</Text>
                    <Separator flex="1"/>
                </div>

                <div className={s.oauthBlock}>
                    <div className={s.oauthBtn}>
                        <picture>
                            <source srcSet="./google.webp" type="image/webp"/>
                            <img src="./google.webp" alt="Моё изображение" className={s.img}/>
                        </picture>
                    </div>
                    <div className={s.oauthBtn}>
                        <picture>
                            <img src="./github.png" alt="Моё изображение" className={s.img}/>
                        </picture>
                    </div>
                </div>
            </form>


        </div>
    )
})

export default RegisterForm