export function resolveAPIUrl(scriptpath: string): string {
    return process.env.VUE_APP_URL_APIROOT + '/' + scriptpath.replace(/^\//, '');
}

export function resolvePageUrl(scriptpath: string): string {
    return process.env.VUE_APP_URL_PAGEROOT + '/' + scriptpath.replace(/^\//, '');
}

export function resolveOwnUrl(scriptpath: string): string {
    return process.env.VUE_APP_URL_OWNROOT + '/' + scriptpath.replace(/^\//, '');
}
