import adapter from '@sveltejs/adapter-static';

const config = {
    kit: {
        adapter: adapter(),
        alias: {
            "@/*": "./path/to/lib/*",
        },
    }
};

export default config;
