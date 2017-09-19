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
      label="Getter"
      item-value="name"
      item-text="name"
    ></v-select>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'marie-curves',
  created () {
    this.$store.dispatch('getAllThings')
  },
  data: function () {
    return {
      thingId: '',
      thing: {},
      getter: ''
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
    }
  },
  computed: mapGetters([
    'things'
  ])
}
</script>

<style>

</style>
