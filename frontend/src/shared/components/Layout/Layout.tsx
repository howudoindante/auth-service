import {Image} from "@chakra-ui/react";
import s from "./Layout.module.scss"

const Layout = () => {
    return (
        <>
            <Image src="./rocket.png" className={s.rocket}/>
            <Image src="./notebook.png" className={s.notebook}/>
            <div className={s.gradient}></div>
        </>
    )
}


export default Layout;