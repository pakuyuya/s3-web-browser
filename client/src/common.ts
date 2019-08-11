import * as path from 'path';

export function resolveAPIUrl(scriptpath: string): string {
    return path.join(process.env.VUE_APP_URL_APIROOT, scriptpath);
}

export function resolvePageUrl(scriptpath: string): string {
    return path.join(process.env.VUE_APP_URL_PAGEROOT, scriptpath);
}

export function resolveOwnUrl(scriptpath: string): string {
    return path.join(process.env.VUE_APP_URL_OWNROOT, scriptpath);
}
