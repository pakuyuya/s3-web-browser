module.exports = {
    devServer: {
        proxy: 'http://localhost:4000',
    },
    pages: {
        'browser': {
            entry: './src/pages/browser/main.ts',
            templates: 'public/index.html',
            title: 'S3 Viewer',
            chunks: [ 'chunk-vendors', 'chunk-common', 'browser' ],
        },
        'login': {
            entry: './src/pages/login/main.ts',
            templates: 'public/index.html',
            title: 'S3 Viewer',
            chunks: [ 'chunk-vendors', 'chunk-common', 'login' ],
        },
    }
};