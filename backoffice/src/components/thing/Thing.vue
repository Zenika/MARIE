<template>
  <div>
    <v-card>
      <v-card-title primary-title class="lighten-1" :class="{'red': !thing.state, 'green': thing.state}">
        <h3 class="headline">{{thing.name}}</h3><br />
      </v-card-title>
      <v-card-text>
        <marie-thing-description :thing="thing"></marie-thing-description>
        <marie-thing-lists :thing="thing"></marie-thing-lists>
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
    <v-dialog v-model="actionDialog" persistent>
      <v-card>
        <v-card-title>
          <span class="headline">Action {{action}}</span>
        </v-card-title>
        <v-card-text>
          <v-text-field v-model="parameter.value" v-for="parameter of parameters" :key="parameter.name" :label="parameter.name"></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn class="blue--text darken-1" flat @click.native="actionDialog = false">Cancel</v-btn>
          <v-btn class="blue--text darken-1" flat @click.native="doActionWithParameters()">Send</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import Description from '@/components/thing/Description'
import Lists from '@/components/thing/Lists'

export default {
  name: 'marie-thing',
  props: ['thing'],
  components: {
    'marie-thing-description': Description,
    'marie-thing-lists': Lists
  },
  data () {
    return {
      action: '',
      actionDialog: false,
      dialog: false,
      test: true,
      state: false,
      parameters: []
    }
  },
  methods: {
    confirmDeletion: function () {
      this.dialog = false
      this.$store.dispatch('deleteThing', this.thing.id)
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

.list__tile {
  padding-left: 0px;
}

.list--group .list__tile {
  padding-left: 0px;
}

</style>
