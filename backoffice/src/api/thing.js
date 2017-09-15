import axios from 'axios'

export default {
  get () {
    return new Promise((resolve, reject) => {
      axios.get(process.env.API_URL + '/things')
      .then(res => res.data)
      .then(res => resolve(res))
      .catch(err => reject(err))
    })
  },
  delete (id) {
    return new Promise((resolve, reject) => {
      axios.delete(process.env.API_URL + '/things/' + id)
        .then(() => resolve(id))
        .catch(err => reject(err))
    })
  },
  update (thing) {
    return new Promise((resolve, reject) => {
      axios.put(process.env.API_URL + '/things', thing)
        .then(() => resolve(thing))
        .catch(err => reject(err))
    })
  },
  create (thing) {
    return new Promise((resolve, reject) => {
      axios.post(process.env.API_URL + '/things', thing)
        .then(res => res.data)
        .then(thing => resolve(thing))
        .catch(err => reject(err))
    })
  },
  do (name, macaddress, parameters) {
    return new Promise((resolve, reject) => {
      axios.post(process.env.API_URL + '/things/do', {name, macaddress, parameters})
        .then(res => resolve(res.data))
        .catch(err => reject(err))
    })
  },
  getter (name, macaddress) {
    return new Promise((resolve, reject) => {
      axios.post(process.env.API_URL + '/things/get', {name, macaddress})
        .then(res => { resolve(res.data) })
        .catch(err => reject(err))
    })
  }
}
