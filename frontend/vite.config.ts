import type {ConfigEnv, UserConfig} from 'vite'
import react from "@vitejs/plugin-react";
import federation from "@originjs/vite-plugin-federation";
import {resolve} from 'path';

function pathResolve(dir: string) {
    return resolve(__dirname, '.', dir)
}


export default ({}: ConfigEnv): UserConfig => {
    return {
        plugins: [react(), federation({
            name: 'host',
            remotes: {
                auth: 'http://localhost:3001/assets/remoteEntry.js',
            },
            shared: ['react', 'react-dom'],
        }),],
        resolve: {
            alias: {
                '@': pathResolve('src'),
                '@components': pathResolve('src/components'),
                '@entities': pathResolve('src/entities'),
                '@shared': pathResolve('src/shared'),
                '@config': pathResolve('src/config'),
                '@pages': pathResolve('src/pages'),
                '@features': pathResolve('src/features'),
            }
        },

    }
}
