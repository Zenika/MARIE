import * as types from '../mutation-types'
import recordService from '../../api/record'

const state = {
  datasetLabel: '',
  chartData: [],
  chartLabels: []
}

const getters = {
  datasetLabel: state => state.datasetLabel,
  chartData: state => state.chartData,
  chartLabels: state => state.chartLabels
}

const actions = {
  getRecords ({commit}, params) {
    const { thingId, getter } = params
    recordService.getRecords(thingId, getter)
          .then(records => {
            commit(types.CHART_DATA_CHANGED, records)
          })
  }
}

const mutations = {
  [types.CHANGE_LABEL] (state, label) {
    state.datasetLabel = label
  },
  [types.CHART_DATA_CHANGED] (state, records) {
    const data = records.map(r => r.value)
    state.chartData = data
    const labels = records.map(r => {
      return new Date(r.date).getHours() + ':' + new Date(r.date).getMinutes() + ':' + new Date(r.date).getSeconds()
    })
    state.chartLabels = labels
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
