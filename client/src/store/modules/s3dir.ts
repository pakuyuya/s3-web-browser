import {S3Profile} from './profile';

export interface S3DirState {
    s3profile?: S3Profile;
    currentDir: string;
    breadcrumbs: any[];
    files?: S3Item[];
}

export interface S3Item {
    type: S3ItemType;
    name: string;
    fullpath: string;
}

export enum S3ItemType {
    File = 'file',
    Directory = 'directory',
}

import { VuexModule, mutation, action, getter, Module } from 'vuex-class-component';

@Module({ namespacedPath: 's3dir/' })
export class S3dirStore extends VuexModule {

    @getter public s3profile?: S3Profile = undefined;
    @getter public currentDir: string = '';
    @getter public breadcrumbs: any[] = [];
    @getter public files: S3Item[] = [
        {
            type: S3ItemType.File,
            name: 'file.txt',
            fullpath: '/file',
        },
        {
            type: S3ItemType.Directory,
            name: 'directory',
            fullpath: '/directory',
        },
    ];


    @action public async setCurrentDir(payload: any) {
        this.updateCurrentDir({ path: payload.path });
        this.updateBreadcrumbs();
    }

    @mutation public updateProfile({profile}: any) {
        this.s3profile = profile;
    }

    @mutation public updateCurrentDir({path}: any) {
        path = path.replace(/^\/+|\/+$/gi, '').trim();
        this.currentDir = path || '';
    }
    @mutation public updateBreadcrumbs() {
        if (this.s3profile === undefined)  {
            this.breadcrumbs = [];
            return;
        }

        const fnMakeBread = (text: string, to: string) => ({
            text,
            disabled: false,
            to,
        });

        const crumbs = [];
        let hrefwork = `/s3/${this.s3profile.id}`;
        crumbs.push(fnMakeBread(this.s3profile.name, hrefwork));

        for (let path of this.currentDir.split('/')) {
            path = path.trim();
            if (path === '') {
                continue;
            }
            hrefwork += `/${path}`;

            crumbs.push(fnMakeBread(path, hrefwork));
        }

        crumbs[crumbs.length - 1].disabled = true;
        this.breadcrumbs = crumbs;
    }
}

export default S3dirStore.ExtractVuexModule(S3dirStore);
