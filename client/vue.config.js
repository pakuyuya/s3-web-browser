module.exports = {
    devServer: {
        proxy: 'http://localhost:4000',
    },
    publicPath: process.env.NODE_ENV === 'production' ? '/browser/' : '/',
    pages: {
        'browser': {
            filename: 'browser.html',
            entry: './src/pages/browser/main.ts',
            templates: 'public/index.html',
            title: 'S3 Viewer',
            chunks: [ 'chunk-vendors', 'chunk-common', 'browser' ],
        },
        'login': {
            filename: 'login.html',
            publicPath: process.env.NODE_ENV === 'production' ? '/login/' : '/',
            entry: './src/pages/login/main.ts',
            templates: 'public/index.html',
            title: 'S3 Viewer',
            chunks: [ 'chunk-vendors', 'chunk-common', 'login' ],
        },
    }
};