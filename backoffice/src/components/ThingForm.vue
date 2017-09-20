<template>
  <div class="marie-thing-form">
    <form @submit.stop.prevent="create">
      <v-text-field
        label="Name"
        value=""
        v-model="thing.name"
        :rules="[rules.required]"
        single-line
      ></v-text-field>
      <v-text-field
        label="Mac Address"
        value=""
        v-model="thing.macaddress"
        single-line
      ></v-text-field>
      <v-text-field
        label="IP Address"
        value=""
        v-model="thing.ipaddress"
        single-line
      ></v-text-field>   
      <v-text-field
        label="Type"
        value=""
        v-model="thing.type"
        :rules="[rules.required]"
        single-line
      ></v-text-field>    
      <v-text-field
        label="Location"
        value=""
        v-model="thing.location"
        single-line
      ></v-text-field>   
      <v-select
        v-bind:items="protocols"
        v-model="thing.protocol"
        label="Protocol"
        single-line
        bottom
        :rules="[rules.required]"
      ></v-select> 
    
      <h5>Actions</h5>
      <v-text-field
        label="New action name"
        value=""
        v-model="newAction"
        @keydown.native.enter.prevent="addAction"
      ></v-text-field>
      <v-card v-for="action in thing.actions" :key="action.id">
        <v-card-title primary-title>
          <v-text-field
            label="Action name"
            value=""
            v-model="action.name"
            :rules="[rules.required]"
          ></v-text-field>

        </v-card-title>
        <v-card-text>
          <p class="subheading">
            Parameters
          </p>
          <v-layout row wrap>
            <v-text-field
              label="New param name"
              value=""
              v-model="newParamName"
              @keydown.native.enter.prevent="addParameter(action)"
            ></v-text-field>
            <v-text-field
              label="New param type"
              value=""
              v-model="newParamType"
              @keydown.native.enter.prevent="addParameter(action)"
            ></v-text-field>
          </v-layout>
            <v-layout row wrap v-for="param in action.parameters" :key="param.id">
              <v-text-field
                label="Parameter name"
                value=""
                v-model="param.name"
                :rules="[rules.required]"
              ></v-text-field>
              <v-text-field
                label="Parameter type"
                value=""
                v-model="param.type"
                :rules="[rules.required]"
              ></v-text-field>
            </v-layout>
        </v-card-text>
      </v-card>
  
      <h5>Getters</h5>
      <v-layout>
        <v-text-field
          label="New getter name"
          value=""
          v-model="newGetterName"
          @keydown.native.enter.prevent="addGetter"
        ></v-text-field>
        <v-text-field
          label="New getter type"
          value=""
          v-model="newGetterType"
          @keydown.native.enter.prevent="addGetter"
        ></v-text-field>
      </v-layout>
      <v-card v-for="getter in thing.getters" :key="getter.id">
        <v-card-text>
          <v-layout>
            <v-text-field
              label="Getter name"
              value=""
              v-model="getter.name"
              :rules="[rules.required]"
            ></v-text-field>
            <v-text-field
              label="Getter type"
              value=""
              v-model="getter.type"
              :rules="[rules.required]"
            ></v-text-field>
          </v-layout>
        </v-card-text>
      </v-card>
      <v-btn v-if="!thing.id" type="submit">Create</v-btn>
      <v-btn v-if="thing.id" type="submit">Update</v-btn>
    </form>
  </div>
</template>

<script>
import ThingModel from '../models/thing'
export default {
  name: 'marie-thing-form',
  data: () => {
    return {
      thing: ThingModel,
      newAction: '',
      newParamName: '',
      newParamType: '',
      newGetterName: '',
      newGetterType: '',
      macaddress: '',
      protocols: [ 'MQTT', 'HTTP' ],
      rules: {
        required: (value) => !!value || 'Required.'
      }
    }
  },
  mounted () {
    if (this.$route.params.id) {
      if (this.$store.getters.things.length === 0) {
        this.$store.dispatch('getAllThings').then(() => {
          this.getThing(this.$route.params.id)
        })
      } else {
        this.getThing(this.$route.params.id)
      }
    }
  },
  watch: {
    '$route.params.id': function (id) {
      if (this.$route.params.id === undefined) {
        this.thing = new ThingModel()
        this.newParamName = ''
        this.newParamType = ''
        this.newGetterName = ''
        this.newGetterType = ''
      }
    }
  },
  methods: {
    create () {
      if (!this.thing.name || !this.thing.type || !this.thing.protocol) {
        return
      }
      for (const g in this.getters) {
        if (!this.thing.getters[g].name || !this.thing.getters[g].type) {
          return
        }
      }
      const actions = this.thing.actions
      for (const a in actions) {
        if (!actions[a].name) {
          return
        }
        const action = actions[a]
        for (const p in action.parameters) {
          if (!action.parameters[p].name || !action.parameters[p].type) {
            return
          }
        }
      }
      if (this.thing.id) {
        this.$store.dispatch('updateThing', this.thing).then(() => {
          this.$router.push('/')
        })
      } else {
        this.$store.dispatch('createThing', this.thing).then(() => {
          this.$router.push('/')
        })
      }
    },
    addAction () {
      this.thing.actions.push({ name: this.newAction, parameters: [] })
      this.newAction = ''
    },
    addGetter () {
      if (this.newGetterName !== '' && this.newGetterType !== '') {
        this.thing.getters.push({ name: this.newGetterName, type: this.newGetterType })
        this.newGetterName = ''
        this.newGetterType = ''
      }
    },
    addParameter (action) {
      if (this.newParamName !== '' && this.newParamType !== '') {
        action.parameters.push({ name: this.newParamName, type: this.newParamType })
        this.newParamName = ''
        this.newParamType = ''
      }
    },
    getThing (id) {
      const thing = this.$store.getters.thing(id)
      if (thing === undefined) {
        this.$store.dispatch('changeSnackbar', 'Thing not found')
      } else {
        this.thing = thing
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.card {
  margin-bottom: 10px;
}

.input-group--text-field {
  height: 60px;
}
</style>
