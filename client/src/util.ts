import * as path from 'path';

export function resolveAPIUrl(scriptpath: string): string {
    return path.join(process.env('URL_APIROOT'), scriptpath);
}

export function resolvePageUrl(scriptpath: string): string {
    return path.join(process.env('URL_PAGEROOT'), scriptpath);
}

export function resolveOwnUrl(scriptpath: string): string {
    return path.join(process.env('URL_OWNROOT'), scriptpath);
}
