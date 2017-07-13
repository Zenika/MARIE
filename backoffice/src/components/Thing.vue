<template>
  <v-card>
    <v-card-title primary-title>
      <h3 class="headline">{{thing.name}}</h3><br />
    </v-card-title>
    <v-card-text>
      <div class="subheading">
        Type : {{thing.type}}<br />
        Protocol : {{thing.protocol}}<br />
        Location : {{thing.location}}
      </div>
      <v-list>
        <v-list-group v-if="thing.actions.length">
          <v-list-tile slot="item">
            <v-list-tile-content>
             Actions
            </v-list-tile-content>
            <v-list-tile-action>
              <v-icon>keyboard_arrow_down</v-icon>
            </v-list-tile-action>
          </v-list-tile>
          <v-list-tile v-for="action in thing.actions" :key="action.name">
            <v-list-tile-content>
              {{action.name}}
            </v-list-tile-content>
          </v-list-tile>
        </v-list-group>

        <v-list-group v-if="thing.getters.length">
          <v-list-tile slot="item">
            <v-list-tile-content>
             Getters
            </v-list-tile-content>
            <v-list-tile-action>
              <v-icon>keyboard_arrow_down</v-icon>
            </v-list-tile-action>
          </v-list-tile>
          <v-list-tile v-for="getter in thing.getters" :key="getter.name">
            <v-list-tile-content>
              {{getter.type}} - {{getter.name}}
            </v-list-tile-content>
          </v-list-tile>
        </v-list-group>
      </v-list>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: 'marie-thing',
  props: ['thing']
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: 100;
}

</style>
