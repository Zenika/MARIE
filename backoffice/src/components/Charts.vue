<template>
  <div>
    <v-select
      v-bind:items="things"
      v-model="thingId"
      label="Thing"
      item-value="id"
      item-text="name"
      @input="changeThing"
    ></v-select>
    <v-select
      v-if="thing.name"
      v-bind:items="thing.getters"
      v-model="getter"
      label="Getter"
      item-value="name"
      item-text="name"
      @input="changeGetter"
    ></v-select>
    <marie-chart :height="100" :labels="labels" :datasetLabel="getter"></marie-chart>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Chart from '@/components/Chart'
import * as types from '../store/mutation-types'

export default {
  name: 'marie-charts',
  created () {
    this.$store.dispatch('getAllThings')
  },
  components: {
    'marie-chart': Chart
  },
  data: function () {
    return {
      thingId: '',
      thing: {},
      getter: '',
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
      data: [40, 20, 12, 39, 10, 20, 39, 5, 40, 20, 12, 11]
    }
  },
  methods: {
    changeThing: function () {
      const thing = this.$store.getters.thing(this.thingId)
      if (thing === undefined) {
        this.$store.dispatch('changeSnackbar', 'Thing not found')
      } else {
        this.thing = thing
      }
    },
    changeGetter: function () {
      this.$store.commit(types.CHANGE_LABEL, this.getter)
      this.$store.commit(types.CHART_DATA_CHANGED, this.data)
    }
  },
  computed: mapGetters([
    'things'
  ])
}
</script>

<style>

</style>
