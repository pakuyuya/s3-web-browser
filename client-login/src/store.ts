import Vue from 'vue';
import Vuex from 'vuex';
import config from './config';

import axiosBase from 'axios';
const axios = axiosBase.create({
  baseURL: `${config.API_BASE}/login`,
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
  },
  responseType: 'json',
});

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    error: '',
    redirectTo: '',
  },
  mutations: {
    setError: (state, payload) => {
      state.error = payload.error;
    },
    login: (state, {redirectTo}) => {
      state.redirectTo = redirectTo;
    },
  },
  actions: {
    login: ({commit}, {loginid, password}) => {
      axios
        .post(`login`, {loginid, password})
        .then((response) => {
          if (response.data.result === 'OK') {
            const redirectTo = response.data.redirectTo;
            commit('login', {redirectTo});
          }
        });
    },
  },
});
