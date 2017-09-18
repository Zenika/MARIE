import thingService from '../../api/thing'
import * as types from '../mutation-types'

// State
const state = {
  actionControl: []
}

// Getters
const getters = {
}

// Actions
const actions = {
  doAction ({commit}, params) {
    thingService.do(params.name, params.macaddress, params.parameters).then(id => {
      commit(types.ACTION_SENT, id)
    })
  }
}

const mutations = {
  [types.ACTION_SENT] (state, id) {
    const timer = setTimeout(() => { state.snackbarText = 'An error occurred' }, 5000)
    state.actionControl.push({id, timer})
  },
  SOCKET_ONMESSAGE (state, message) {
    let t, index
    switch (message.topic) {
      case 'actions':
        index = state.things.findIndex(thing => thing.macaddress === message.macaddress)
        t = state.things[index]
        t.actions = message.actions
        state.things.splice(index, 1, t)
        break
      case 'action-done':
        index = state.actionControl.findIndex(id => id.id === message.id)
        if (state.actionControl[index]) {
          clearTimeout(state.actionControl[index].timer)
        }
        state.actionControl.splice(index, 1)
        this.dispatch('changeSnackbar', 'Action executed successfully')
        break
    }
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
