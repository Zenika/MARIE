import axios from 'axios'

export default {
  getRecords (thingId, getter) {
    return new Promise((resolve, reject) => {
      axios.get(`${process.env.API_URL}/records/${thingId}/${getter}`)
        .then(res => res.data)
        .then(records => resolve(records))
        .catch(err => reject(err))
    })
  }
}
