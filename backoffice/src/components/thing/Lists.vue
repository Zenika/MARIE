<template>
  <v-list>
    <v-list-group v-if="thing.actions && thing.actions.length">
      <v-list-tile slot="item">
        <v-list-tile-content>
          Actions
        </v-list-tile-content>
        <v-list-tile-action>
          <v-icon>keyboard_arrow_down</v-icon>
        </v-list-tile-action>
      </v-list-tile>
      <v-list-tile v-for="action in thing.actions" :key="action.name" @click.native="doAction(action)">
        <v-list-tile-content>
          {{action.name}}
        </v-list-tile-content>
      </v-list-tile>
    </v-list-group>

    <v-list-group v-if="thing.getters && thing.getters.length">
      <v-list-tile slot="item">
        <v-list-tile-content>
          Getters
        </v-list-tile-content>
        <v-list-tile-action>
          <v-icon>keyboard_arrow_down</v-icon>
        </v-list-tile-action>
      </v-list-tile>
      <v-list-tile v-tooltip:top="{ html: value, visible: false }" v-for="getter in thing.getters" :key="getter.name" @mouseover.native="doGetter(getter.name)">
        <v-list-tile-content>
          {{getter.type}} - {{getter.name}}
        </v-list-tile-content>
      </v-list-tile>
    </v-list-group>
  </v-list>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'marie-thing-lists',
  props: ['thing'],
  methods: {
    doAction: function (action) {
      if (this.thing.state) {
        if (action.parameters.length > 0) {
          this.parameters = action.parameters
          this.action = action.name
          this.actionDialog = true
        } else {
          this.$store.dispatch('doAction', {name: action.name, macaddress: this.thing.macaddress})
        }
      }
    },
    doGetter: function (name) {
      if (this.thing.state) {
        this.$store.dispatch('doGetter', {name, macaddress: this.thing.macaddress})
      }
    },
    doActionWithParameters: function () {
      this.actionDialog = false
      this.$store.dispatch('doAction', {name: this.action, macaddress: this.thing.macaddress, parameters: this.parameters})
    }
  },
  computed: mapGetters([
    'value'
  ])
}
</script>

<style>

</style>
