import Vue from 'vue'
import Vuex from 'vuex'

import things from './modules/things'
import actions from './modules/actions'
import getters from './modules/getters'
import snackbar from './modules/snackbar'
import speech from './modules/speech'
import chart from './modules/chart'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    actions,
    chart,
    getters,
    snackbar,
    speech,
    things
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
