/* eslint-disable */
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    token: null,
  },
  mutations: {
    update_token(state, new_token) {
      state.token = new_token;
    },
  },
  actions: {

  },
});
