import thingService from '../../api/thing'
import * as types from '../mutation-types'

// Initial state
const state = {
  things: []
}

// Getters
const getters = {
  things: state => state.things,
  thing: state => id => state.things.find(t => t.id === id)
}

// Actions
const actions = {
  getAllThings ({commit}) {
    return new Promise((resolve, reject) => {
      thingService.get().then(things => {
        commit(types.THINGS_RECEIVED, things)
        resolve()
      }).catch((err) => reject(err))
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
  SOCKET_ONMESSAGE (state, message) {
    let t, index
    switch (message.topic) {
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
