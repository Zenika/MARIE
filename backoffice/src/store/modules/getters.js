import thingService from '../../api/thing'
import * as types from '../mutation-types'

// State
const state = {
  getterControl: [],
  value: 'No value'
}

// Getters
const getters = {
  value: state => state.value
}

// Actions
const actions = {
  doGetter ({commit}, params) {
    thingService.getter(params.name, params.macaddress).then(id => {
      commit(types.GETTER_SENT, id)
    })
  }
}

const mutations = {
  [types.GETTER_SENT] (state, id) {
    state.getterControl.push({id})
  },
  SOCKET_ONMESSAGE (state, message) {
    let t, index
    switch (message.topic) {
      case 'getters':
        index = state.things.findIndex(thing => thing.macaddress === message.macaddress)
        t = state.things[index]
        t.getters = message.getters
        state.things.splice(index, 1, t)
        break
      case 'getter-value':
        index = state.getterControl.findIndex(id => id.id === message.id)
        state.getterControl.splice(index, 1)
        state.value = message.value
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
