import thingService from '../../api/thing'
import * as types from '../mutation-types'

// Initial state
const state = {
  things: [],
  actionControl: [],
  getterControl: [],
  snackbarText: '',
  value: 'No value'
}

// Getters
const getters = {
  things: state => state.things,
  thing: state => id => state.things.find(t => t.id === id),
  snackbarText: state => state.snackbarText,
  value: state => state.value
}

// Actions
const actions = {
  getAllThings ({commit}) {
    thingService.get().then(things => {
      commit(types.THINGS_RECEIVED, things)
    })
  },
  deleteThing ({commit}, thingId) {
    thingService.delete(thingId).then(() => {
      commit(types.THING_DELETED, thingId)
    })
  },
  updateThing ({commit}, thing) {
    return new Promise((resolve, reject) => {
      thingService.update(thing).then(() => {
        commit(types.THING_UPDATED, {thing})
        resolve()
      })
    })
  },
  createThing ({commit}, thing) {
    thingService.create(thing).then(thing => {
      commit(types.THING_CREATED, thing)
    })
  },
  doAction ({commit}, params) {
    thingService.do(params.name, params.macaddress, params.parameters).then(id => {
      commit(types.ACTION_SENT, id)
    })
  },
  doGetter ({commit}, params) {
    thingService.getter(params.name, params.macaddress).then(id => {
      commit(types.GETTER_SENT, id)
    })
  }
}

// Mutations
const mutations = {
  [types.THINGS_RECEIVED] (state, things) {
    state.things = things || []
  },
  [types.THING_DELETED] (state, thingId) {
    const index = state.things.findIndex(t => t.id === thingId)
    state.things.splice(index, 1)
  },
  [types.THING_UPDATED] (state, thing) {
    const index = state.things.findIndex(t => t.id === thing.id)
    state.things.splice(index, 1, thing)
  },
  [types.THING_CREATED] (state, thing) {
    state.things.push(thing)
  },
  [types.ACTION_SENT] (state, id) {
    const timer = setTimeout(() => { state.snackbarText = 'An error occurred' }, 5000)
    state.actionControl.push({id, timer})
  },
  [types.SNACKBAR_CLOSED] (state) {
    state.snackbarText = ''
  },
  [types.GETTER_SENT] (state, id) {
    state.getterControl.push({id})
  },
  SOCKET_ONMESSAGE (state, message) {
    let t
    let index
    switch (message.topic) {
      case 'actions':
        index = state.things.findIndex(thing => thing.macaddress === message.macaddress)
        t = state.things[index]
        t.actions = message.actions
        state.things.splice(index, 1, t)
        break
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
      case 'action-done':
        index = state.actionControl.findIndex(id => id.id === message.id)
        clearTimeout(state.actionControl[index].timer)
        state.actionControl.splice(index, 1)
        state.snackbarText = 'Action executed successfully'
        break
      case 'register':
        state.things.push(message)
        break
      case 'state-off':
        index = state.things.findIndex(t => t.macaddress === message.macaddress)
        if (index !== -1) {
          t = state.things[index]
          t.state = false
          state.things.splice(index, 1, t)
        }
        break
      case 'state-on':
        index = state.things.findIndex(t => t.macaddress === message.macaddress)
        if (index !== -1) {
          t = state.things[index]
          t.state = true
          state.things.splice(index, 1, t)
        }
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
