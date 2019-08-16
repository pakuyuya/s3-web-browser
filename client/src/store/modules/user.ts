import { VuexModule, mutation, action, getter, Module } from 'vuex-class-component';

import axios from 'axios';
import * as common from '../../common';


@Module({ namespacedPath: 'user/' })
export class UserStore extends VuexModule {
    @getter public name: string = 'Guest';

    @action public async login(payload: {loginid: string, password: string}) {
        const url = common.resolveAPIUrl('login');
        const params = {
            loginid: payload.loginid,
            password: payload.password,
        };
        return axios
            .post(url, params)
            .then((response) => {
            if (response.data.result === 'OK') {
                return response.data.redirectTo;
            }
            return undefined;
        });
    }
}

export default UserStore.ExtractVuexModule(UserStore);
