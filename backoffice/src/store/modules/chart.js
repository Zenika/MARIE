import * as types from '../mutation-types'
import recordService from '../../api/record'

const state = {
  datasetLabel: '',
  chartData: []
}

const getters = {
  datasetLabel: state => state.datasetLabel,
  chartData: state => state.chartData
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
    const data = records.map(r => { const {value} = r; return value })
    state.chartData = data
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
