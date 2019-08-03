import Vue from 'vue';
import Vuex from 'vuex';

import profile from './modules/profile';
import s3directory from './modules/s3directory';
import user from './modules/user';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    profile,
    s3directory,
    user,
  },
  mutations: {

  },
  actions: {

  },
});
