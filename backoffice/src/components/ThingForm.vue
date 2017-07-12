<template>
  <div class="marie-thing-form">
    <h1>Create a thing</h1>
    <form @submit.stop.prevent="create()">
      <md-input-container>
        <label>Name</label>
        <md-input v-model="name"></md-input>
      </md-input-container>
  
      <md-input-container>
        <label>Type</label>
        <md-input v-model="type"></md-input>
      </md-input-container>

      <md-input-container>
        <label>Location</label>
        <md-input v-model="location"></md-input>
      </md-input-container>
  
      <md-input-container>
        <label>Protocol (WS, HTTP, MQTT, ...)</label>
        <md-input v-model="protocol"></md-input>
      </md-input-container>
  
      <div>
        <h2>Actions</h2>
        <md-card v-for="action in actions" :key="action.id">
          <md-card-header>
            <md-input-container>
              <label>Name</label>
              <md-input v-model="action.name"></md-input>
            </md-input-container>
          </md-card-header>
          <md-card-content>
            <h3>Parameters</h3>
            <md-table>
              <md-table-header>
                <md-table-head>Name</md-table-head>
                <md-table-head>Type</md-table-head>
              </md-table-header>
              <md-table-body>
                <md-table-row v-for="parameter in action.parameters" :key="parameter.id">
                  <md-table-cell>
                    <md-input-container md-inline>
                      <md-input v-model="parameter.name"></md-input>
                    </md-input-container>
                  </md-table-cell>
                  <md-table-cell>
                    <md-input-container md-inline>
                      <md-input v-model="parameter.type"></md-input>
                    </md-input-container>
                  </md-table-cell>
                </md-table-row>
              </md-table-body>
            </md-table>
            <md-button @click="addParameter(action)" class="md-icon-button md-raised">
              <md-icon>add</md-icon>
            </md-button>
          </md-card-content>
        </md-card>
        <md-button @click="addAction" class="md-icon-button md-raised">
          <md-icon>add</md-icon>
        </md-button>
      </div>
  
      <div>
        <h2>Getters</h2>
        <md-card v-for="getter in getters" :key="getter.id">
          <md-card-content>
            <md-input-container>
              <label>Name</label>
              <md-input v-model="getter.name"></md-input>
            </md-input-container>
            <md-input-container>
              <label>Type</label>
              <md-input v-model="getter.type"></md-input>
            </md-input-container>
          </md-card-content>
        </md-card>
        <md-button @click="addGetter" class="md-icon-button md-raised">
          <md-icon>add</md-icon>
        </md-button>
      </div>
      <input type="submit" class="md-button md-raised" value="Create">
    </form>
  </div>
</template>

<script>
import router from '@/router/index'
export default {
  name: 'marie-thing-form',
  data () {
    return {
      name: '',
      type: '',
      protocol: '',
      location: '',
      actions: [],
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
      this.actions.push({ name: '', parameters: [] })
    },
    addGetter () {
      this.getters.push({ name: '', type: '' })
    },
    addParameter (action) {
      action.parameters.push({ name: '', type: '' })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: 100;
}

.md-card {
  margin: 10px;
}

form {
  width: 75%;
  display: inline-block;
}
</style>
