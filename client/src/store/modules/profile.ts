import { mapActions } from 'vuex';

import { VuexModule, mutation, action, getter, Module } from 'vuex-class-component';
export interface ProfileState {
    list: S3Profile[];
}
export interface S3Profile {
    id: string;
    name: string;
}

@Module({ namespacedPath: 'profile/' })
export class ProfileStore extends VuexModule {
    @getter public list: S3Profile[] = [
        {id: 'test', name: 'sample'},
    ];

    @action public async findById(id: string) {
        return this.list.find((item: S3Profile) => id === item.id);
    }
}

export default ProfileStore.ExtractVuexModule(ProfileStore);
