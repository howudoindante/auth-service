import {Controller, useForm} from "react-hook-form";
import {yupResolver} from "@hookform/resolvers/yup";
import {loginSchema} from "@entities/auth/model/validation.ts";
import {Button, Checkbox, Field, Heading, Image, Input, Link as ChakraLink, Separator, Text} from "@chakra-ui/react";
import {PasswordInput} from "@components/ui/password-input.tsx";
import s from "./LoginForm.module.scss"
import {observer} from "mobx-react-lite";
import authStore from "@entities/auth/model/store.ts";
import Layout from "@shared/components/Layout/Layout.tsx";
import {Link} from "react-router"

const LoginForm = observer(() => {
    const {login} = authStore;
    const {control, handleSubmit} = useForm({
        resolver: yupResolver(loginSchema),
    })
    return (
        <div className={s.container}>
            <Layout/>
            <form onSubmit={handleSubmit((v) => login(v))} className={s.form}>
                <Image src="./logo-black.png" className={s.logo}/>

                <Heading className={s.heading}>Login to your account</Heading>
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
                <div className={s.row}>
                    <Checkbox.Root>
                        <Checkbox.HiddenInput/>
                        <Checkbox.Control/>
                        <Checkbox.Label>Remember me</Checkbox.Label>
                    </Checkbox.Root>


                    <ChakraLink
                        asChild
                        variant="underline"
                        colorPalette="teal"
                    >
                        <Link to="/auth/register"> Forgot password?
                        </Link>

                    </ChakraLink>
                </div>
                <Button className={s.btn} type="submit">Sign in with email</Button>
                <div className={`${s.row} ${s.centeredContent}`}>
                    <Text className={s.link}>Don't have account?
                        <ChakraLink
                            asChild
                            variant="underline"
                            colorPalette="teal">
                            <Link to="/auth/register"> Let's create!
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

export default LoginForm