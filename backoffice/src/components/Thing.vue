<template>
  <v-card>
    <v-card-title primary-title>
      <h3 class="headline">{{thing.name}}</h3><br />
    </v-card-title>
    <v-card-text>
      <div class="subheading">
        Id : {{thing.id}}<br />
        MacAddress : {{thing.macaddress}}<br />
        IPAddress : {{thing.ipaddress}}<br />
        Type : {{thing.type}}<br />
        Protocol : {{thing.protocol}}<br />
        Location : {{thing.location}}
      </div>
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
          <v-list-tile v-for="action in thing.actions" :key="action.name">
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
          <v-list-tile v-for="getter in thing.getters" :key="getter.name">
            <v-list-tile-content>
              {{getter.type}} - {{getter.name}}
            </v-list-tile-content>
          </v-list-tile>
        </v-list-group>
      </v-list>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <router-link :to="{name: 'marie-thing-form-update', params: {id: thing.id}}"><v-btn flat>Update</v-btn></router-link>
      <v-layout row justify-center style="position: relative;">
        <v-dialog v-model="dialog" lazy absolute>
          <v-btn primary flat slot="activator">Delete</v-btn>
          <v-card>
            <v-card-title>
              <div class="headline">Confirmation</div>
            </v-card-title>
            <v-card-text>Are you sure you want to delete this thing?</v-card-text>
            <v-card-actions>
              <v-btn class="green--text darken-1" flat="flat" @click.native="dialog = false">No</v-btn>
              <v-btn class="green--text darken-1" flat="flat" @click.native="confirmDeletion()">Yes</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-layout>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'marie-thing',
  props: ['thing'],
  data () {
    return {
      dialog: false,
      test: true
    }
  },
  methods: {
    confirmDeletion: function () {
      this.dialog = false
      this.$emit('delete')
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: 100;
}

a {
  text-decoration: none;
}

</style>
