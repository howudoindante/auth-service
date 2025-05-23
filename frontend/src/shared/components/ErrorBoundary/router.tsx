import {Link, useRouteError} from "react-router";
import s from "./ErrorBoundary.module.scss"
import {Heading, Image, Link as ChakraLink, Text} from "@chakra-ui/react";

export default function ErrorPage() {
    const error = useRouteError();
    console.log(error)
    return (
        <div className={s.error}>


            <Image src="./error.png" className={s.errorIcon}/>
            <Heading size="5xl">Something went wrong!</Heading>
            <Text textStyle="2xl" className={s.text}>{(error as { data: string }).data}</Text>

            <Text textStyle="2xl" className={s.text}><ChakraLink
                asChild
                variant="underline"
                colorPalette="teal">
                <Link to="/"> Return to main page?
                </Link>
            </ChakraLink></Text>

        </div>
    );
}