import { sveltekit } from '@sveltejs/kit/vite';
import type { UserConfig } from 'vite';

const config: UserConfig = {
	plugins: [sveltekit()],

    css: {
        preprocessorOptions: {
            scss: {
                additionalData: '@use \'src/lib/scss/variables\' as *;',
            },
        },
    },
};

export default config;
