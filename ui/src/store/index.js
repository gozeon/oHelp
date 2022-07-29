import Vue from 'vue'
import Vuex from 'vuex'
import global from './module/global'
import user from './module/user'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    global,
    user,
  },
})
