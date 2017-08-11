<template>
  <div class="marie-things">
    <v-layout row wrap>
      <v-flex xs3 v-for="thing in things" :key="thing.id" >
        <marie-thing :thing="thing" @delete="deleteThing(thing)"></marie-thing>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
import Thing from '@/components/Thing'

export default {
  name: 'marie-things',
  components: {
    'marie-thing': Thing
  },
  data: function () {
    return {
      things: []
    }
  },
  mounted () {
    this.getThings()
    this.$options.sockets.onmessage = (data) => {
      const thing = JSON.parse(data.data)
      let found
      switch (thing.topic) {
        case 'register':
          this.things.push(JSON.parse(data.data))
          break
        case 'actions':
          found = this.things.find(t => t.macaddress === thing.macaddress)
          found.actions = thing.actions
          break
        case 'getters':
          found = this.things.find(t => t.macaddress === thing.macaddress)
          found.getters = thing.getters
          break
      }
    }
  },
  methods: {
    getThings () {
      this.$http.get(process.env.API_URL + '/things')
        .then(res => res.data)
        .then(res => { this.things = res || [] })
    },
    deleteThing: function (thing) {
      this.$http.delete(process.env.API_URL + '/things/' + thing.id)
                .then(() => {
                  const index = this.things.findIndex(t => t.id === thing.id)
                  this.things.splice(index, 1)
                })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: 100;
}

.flex {
  margin-bottom: 10px;
}
</style>
