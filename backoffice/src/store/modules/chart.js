import * as types from '../mutation-types'

const state = {
  datasetLabel: '',
  chartData: []
}

const getters = {
  datasetLabel: state => state.datasetLabel,
  chartData: state => state.chartData
}

const actions = {

}

const mutations = {
  [types.CHANGE_LABEL] (state, label) {
    state.datasetLabel = label
  },
  [types.CHART_DATA_CHANGED] (state, data) {
    state.chartData = data
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
