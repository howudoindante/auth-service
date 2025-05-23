import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import {ChakraProvider, createSystem, defaultConfig, defineConfig} from "@chakra-ui/react"
import {Toaster} from "@components/ui/toaster.tsx"
import './index.css'
import App from './App.tsx'
import {ColorModeProvider} from "@components/ui/color-mode.tsx";
import {ErrorBoundary} from "@shared/components/ErrorBoundary";


const lightConfig = defineConfig({
    theme: {
        tokens: {
            colors: {},
        },
    },
});


const defaultSystem = createSystem(defaultConfig, lightConfig);


const Root = () => {

    return (
        <StrictMode>
            <ChakraProvider value={defaultSystem}>
                <ErrorBoundary>
                    <ColorModeProvider forcedTheme="light">
                        <Toaster/>
                        <App/>
                    </ColorModeProvider>
                </ErrorBoundary>
            </ChakraProvider>
        </StrictMode>
    )
}

createRoot(document.getElementById('root')!).render(
    <Root/>,
)
