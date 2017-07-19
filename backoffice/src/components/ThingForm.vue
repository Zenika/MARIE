<template>
  <div class="marie-thing-form">
    <form @submit.stop.prevent="create">
      <v-text-field
        label="Name"
        value=""
        v-model="name"
        :rules="[rules.required]"
        single-line
      ></v-text-field>
      <v-text-field
        label="Type"
        value=""
        v-model="type"
        :rules="[rules.required]"
        single-line
      ></v-text-field>    
      <v-text-field
        label="Location"
        value=""
        v-model="location"
        single-line
      ></v-text-field>   
      <v-select
        v-bind:items="protocols"
        v-model="protocol"
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
      <v-card v-for="action in actions" :key="action.id">
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
      <v-card v-for="getter in getters" :key="getter.id">
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
      <v-btn v-if="!id" type="submit">Create</v-btn>
      <v-btn v-if="id" type="submit">Update</v-btn>
    </form>
  </div>
</template>

<script>
import router from '@/router/index'
export default {
  name: 'marie-thing-form',
  data: () => {
    return {
      id: '',
      name: '',
      type: '',
      protocol: '',
      location: '',
      actions: [],
      newAction: '',
      newParamName: '',
      newParamType: '',
      newGetterName: '',
      newGetterType: '',
      getters: [],
      protocols: [ 'MQTT' ],
      rules: {
        required: (value) => !!value || 'Required.'
      }
    }
  },
  mounted () {
    if (this.$route.params.id) {
      this.getThing(this.$route.params.id)
    }
  },
  methods: {
    create () {
      if (!this.name || !this.type || !this.protocol) {
        return
      }
      for (const g in this.getters) {
        if (!this.getters[g].name || !this.getters[g].type) {
          return
        }
      }
      for (const a in this.actions) {
        if (!this.actions[a].name) {
          return
        }
        const action = this.actions[a]
        for (const p in action.parameters) {
          if (!action.parameters[p].name || !action.parameters[p].type) {
            return
          }
        }
      }
      const thing = {
        id: this.id,
        name: this.name,
        type: this.type,
        protocol: this.protocol,
        actions: this.actions,
        getters: this.getters,
        location: this.location
      }
      if (this.id) {
        this.$http.put(process.env.API_URL + '/things', thing)
          .then(res => router.push('/'))
      } else {
        this.$http.post(process.env.API_URL + '/things', thing)
          .then(res => router.push('/'))
      }
    },
    getThing (id) {
      this.$http.get(process.env.API_URL + '/things/' + id)
        .then(res => res.data)
        .then(res => {
          this.id = res.id
          this.name = res.name
          this.type = res.type
          this.location = res.location
          this.protocol = res.protocol
          this.actions = res.actions
          this.getters = res.getters
        })
    },
    addAction () {
      this.actions.push({ name: this.newAction, parameters: [] })
      this.newAction = ''
    },
    addGetter () {
      if (this.newGetterName !== '' && this.newGetterType !== '') {
        this.getters.push({ name: this.newGetterName, type: this.newGetterType })
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
  height: 40px;
}
</style>
