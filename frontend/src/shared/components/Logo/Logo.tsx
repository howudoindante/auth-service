import {Image} from "@chakra-ui/react";
import s from "./Logo.module.scss"
import classNames from "@shared/utils/classNames.ts";

export const Logo = ({className, mode = "light"}: { className?: string, mode?: "dark" | "light" }) => {

    const files = {
        dark: "./logo-light.png",
        light: "./logo-black.png",
    }
    return (
        <Image src={files[mode]} className={classNames(s.logo, {
            [className ?? ""]: !!className
        })}/>
    )
}