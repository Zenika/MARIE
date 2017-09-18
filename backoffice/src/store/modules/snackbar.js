import * as types from '../mutation-types'

// State
const state = {
  snackbarText: ''
}

const getters = {
  snackbarText: state => state.snackbarText
}

const actions = {
  changeSnackbar ({commit}, text) {
    commit(types.CHANGE_SNACKBAR, text)
  }
}

const mutations = {
  [types.CHANGE_SNACKBAR] (state, text) {
    state.snackbarText = text
  },
  [types.SNACKBAR_CLOSED] (state) {
    state.snackbarText = ''
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
