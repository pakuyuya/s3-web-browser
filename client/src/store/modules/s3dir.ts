import {S3Profile} from './profile';

export interface S3DirState {
    s3profile?: S3Profile;
    currentDir?: string;
    dirBreads?: string[];
    files?: S3Item[];
}

export interface S3Item {
    type: S3ItemType;
    name: string;
}

export enum S3ItemType {
    File = 'file',
    Directory = 'directory',
}


const state: S3DirState = {
    s3profile: undefined,
    currentDir: '',
    dirBreads: [],
    files: [],
};
const actions = {

};

const mutations = {
    changeProfile(s: S3DirState, payload: any) {
        s.s3profile = payload.profile;
    },
};
export default {
    namespaced: true,
    state,
    actions,
    mutations,
};
