<template>
  <div>
    <v-card>
      <v-card-title primary-title :class="{'red lighten-1': !thing.state, 'green lighten-1': thing.state}">
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
            <v-list-tile v-for="action in thing.actions" :key="action.name" @click.native="doAction(action.name)">
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
    <v-snackbar
      :timeout="2000"
      :bottom="true"
      :right="true"
      v-model="snackbar"
    >
    {{text}}
    <v-btn flat class="pink--text" @click.native="snackbar = false">Close</v-btn>
    </v-snackbar>
  </div>
</template>

<script>
export default {
  name: 'marie-thing',
  props: ['thing'],
  data () {
    return {
      dialog: false,
      test: true,
      id: '',
      state: false,
      snackbar: false,
      text: '',
      value: 'No value'
    }
  },
  mounted () {
    this.$options.sockets.onmessage = (res) => {
      res = JSON.parse(res.data)
      if (res.id === this.id) {
        if (res.value) {
          this.value = res.value
        } else {
          if (res.code === 0) {
            this.text = 'Action executed successfully'
            this.snackbar = true
          } else {
            this.text = 'An error occured'
            this.snackbar = true
          }
        }
      } else if (res['state-off'] === this.thing.macaddress) {
        this.thing.state = false
      } else if (res['state-on'] === this.thing.macaddress) {
        this.thing.state = true
      }
    }
  },
  methods: {
    confirmDeletion: function () {
      this.dialog = false
      this.$emit('delete')
    },
    doAction: function (name) {
      if (this.thing.state) {
        this.$http.post(process.env.API_URL + '/things/do', {name, macaddress: this.thing.macaddress})
                  .then(res => { this.id = res.data })
      }
    },
    doGetter: function (name) {
      if (this.thing.state) {
        this.value = 'No value'
        this.$http.post(process.env.API_URL + '/things/get', {name, macaddress: this.thing.macaddress})
                  .then(res => { this.id = res.data })
      }
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
