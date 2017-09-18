// State
const state = {
  result: {
    id: '',
    action: '',
    location: '',
    type: '',
    getter: ''
  }
}

// Getters
const getters = {
  result: state => state.result
}

// Actions
const actions = {

}

// Mutations
const mutations = {
  SOCKET_ONMESSAGE: (state, message) => {
    if (message.topic === 'speech-action' || message.topic === 'speech-getter') {
      state.result.action = message.doing || ''
      state.result.getter = message.variable || ''
      state.result.id = message.id
      state.result.location = message.in
      state.result.type = message.on
    }
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
