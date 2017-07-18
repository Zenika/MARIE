<template>
  <div class="marie-things">
    <v-layout>
      <v-flex xs3>
        <marie-thing :thing="thing" v-for="thing in things" :key="thing.name" @delete="deleteThing(thing)"></marie-thing>
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
  data: () => {
    return {
      things: []
    }
  },
  mounted () {
    this.getThings()
  },
  methods: {
    getThings () {
      this.$http.get(process.env.API_URL + '/things')
        .then(res => { this.things = res.data })
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
</style>
