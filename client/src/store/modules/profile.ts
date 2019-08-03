
export interface ProfileState {
    list: S3Profile[];
}
export interface S3Profile {
    id: string;
    name: string;
}

const state: ProfileState = {
    list: [
        {id: 'test', name: 'sample'},
    ],
};
const actions = {

};
const mutations = {

};
export default {
    namespaced: true,
    state,
    actions,
    mutations,
};
