import { VuexModule, mutation, action, getter, Module } from 'vuex-class-component';
import axios from 'axios';
import * as common from '../../common'
export interface ProfileState {
    list: S3Profile[];
}
export interface S3Profile {
    profileid?: string;
    profilename?: string;
    bucket?: string;
    connjson?: string;
}

@Module({ namespacedPath: 'profile/' })
export class ProfileStore extends VuexModule {
    @getter public list: S3Profile[] = [
        {profileid: 'test', profilename: 'sample'},
    ];

    @action public async findById(profileid: string) {
        return this.list.find((item: S3Profile) => profileid === item.profileid);
    }

    @action public async insert(model: S3Profile) {
        await axios.post(common.resolveAPIUrl('profile'), model)
            .then((res) => {
                this.reload();
            });
    }
    @action public async delete(model: S3Profile) {
        await axios.delete(common.resolveAPIUrl(`profile/${model.profileid}`))
            .then((res) => {
                this.reload();
            });
    }

    @action public async reload() {
        await axios.get(common.resolveAPIUrl('profiles'))
            .then((res) => {
                this.list = res.data;
            });
    }
}

export default ProfileStore.ExtractVuexModule(ProfileStore);
