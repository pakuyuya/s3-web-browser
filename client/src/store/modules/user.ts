import { VuexModule, mutation, action, getter, Module } from 'vuex-class-component';

import axios from 'axios';
import * as common from '../../common';


@Module({ namespacedPath: 'user/' })
export class UserStore extends VuexModule {
    @getter public name: string = 'Guest';
    @getter public permissions: any = {};

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
    @action public async loadLoginInfo() {
        const url = common.resolveAPIUrl('logininfo');
        return axios
            .get(url)
            .then(((response) => {
                this.setUsername(response.data.username);
                this.setPermissions(response.data.permissions);
            }
        ));
    }

    @mutation public setUsername(username: string) {
        this.name = username;
    }
    @mutation public setPermissions(permissions: any) {
        this.permissions = permissions;
        console.log(permissions);
    }
}

export default UserStore.ExtractVuexModule(UserStore);
