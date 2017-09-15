import Vue from 'vue'
import Vuex from 'vuex'

import things from './modules/things'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    things
  },
  mutations: {
    SOCKET_ONOPEN (state) {

    }
  }
})
