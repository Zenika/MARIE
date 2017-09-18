import Vue from 'vue'
import Vuex from 'vuex'

import things from './modules/things'
import actions from './modules/actions'
import getters from './modules/getters'
import snackbar from './modules/snackbar'
import speech from './modules/speech'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    things,
    actions,
    getters,
    snackbar,
    speech
  },
  mutations: {
    SOCKET_ONOPEN (state) {
      this.dispatch('changeSnackbar', 'Connected to server with websockets')
    },
    SOCKET_ONERROR (state) {
      this.dispatch('changeSnackbar', 'An error occured')
    },
    SOCKET_ONCLOSE (state) {
      this.dispatch('changeSnackbar', 'Socket closed')
    }
  }
})
