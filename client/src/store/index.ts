import Vue from 'vue';
import Vuex from 'vuex';

import profile from './modules/profile';
import s3dir from './modules/s3dir';
import user from './modules/user';


Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    profile,
    s3dir,
    user,
  },
  mutations: {

  },
  actions: {

  },
});
