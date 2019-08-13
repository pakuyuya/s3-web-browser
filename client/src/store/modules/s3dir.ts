import {S3Profile} from './profile';
import axios from 'axios';
import * as common from '../../common';

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
    size: string;
    lastmodified: string;
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
    @getter public error: string = '';
    @getter public files: S3Item[] = [
        {
            type: S3ItemType.File,
            name: 'file.txt',
            fullpath: '/file',
            size: '10 Bytes',
            lastmodified: '2019/01/01 10:00:00',
        },
        {
            type: S3ItemType.Directory,
            name: 'directory',
            fullpath: '/directory',
            size: '10 Bytes',
            lastmodified: '2019/01/01 10:00:00',
        },
    ];


    @action public async setCurrentDir(payload: any) {
        if (!this.s3profile) {
            return;
        }

        const params = {
            profileid: this.s3profile.profileid,
            path: payload.path,
        };
        this.setError('');
        const url = common.resolveAPIUrl(`s3dir/${this.s3profile.profileid}/${payload.path}`);
        axios.get(url)
            .then((res) => {
                this.files = res.data.map((item: any) => ({
                   type: item.type,
                   name: item.name,
                   fullpath: item.fullpath,
                   size: item.size,
                   lastmodified: item.lastmodified,
                }));
                this.updateCurrentDir({ path: payload.path });
                this.updateBreadcrumbs();
            })
            .catch((error) => {
                const msg = error.response.data.message || 'S3への接続に失敗しました。'
                this.setError(msg);
            });
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
        let hrefwork = `/s3/${this.s3profile.profileid}`;
        crumbs.push(fnMakeBread(this.s3profile.profilename || '', hrefwork));

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
    
    @mutation public setError(error: string) {
        this.error = error;
    }

}

export default S3dirStore.ExtractVuexModule(S3dirStore);
