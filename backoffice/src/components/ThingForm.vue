<template>
  <div class="marie-thing-form">
    <form @submit.stop.prevent="create">
      <v-text-field
        label="Name"
        value=""
        v-model="name"
      ></v-text-field>
      <v-text-field
        label="Type"
        value=""
        v-model="type"
      ></v-text-field>    
      <v-text-field
        label="Location"
        value=""
        v-model="location"
      ></v-text-field>   
      <v-text-field
        label="Protocol (WS, HTTP, MQTT, ...)"
        value=""
        v-model="protocol"
      ></v-text-field>   
    
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
              ></v-text-field>
              <v-text-field
                label="Parameter type"
                value=""
                v-model="param.type"
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
            ></v-text-field>
            <v-text-field
              label="Getter type"
              value=""
              v-model="getter.type"
            ></v-text-field>
          </v-layout>
        </v-card-text>
      </v-card>
      <v-btn type="submit">Create</v-btn>
    </form>
  </div>
</template>

<script>
import router from '@/router/index'
export default {
  name: 'marie-thing-form',
  data: () => {
    return {
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
      getters: []
    }
  },
  methods: {
    create () {
      console.log(this.name)
      const thing = {
        name: this.name,
        type: this.type,
        protocol: this.protocol,
        actions: this.actions,
        getters: this.getters,
        location: this.location
      }
      this.$http.post(process.env.API_URL + '/things', thing)
        .then(res => router.push('/'))
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
</style>
