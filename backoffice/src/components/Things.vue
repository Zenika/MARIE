<template>
  <div class="hello">
    <h1>Things</h1>
    <md-layout>
      <marie-thing :thing="thing" v-for="thing in things" :key="thing.name"></marie-thing>
    </md-layout>
  </div>
</template>

<script>
import Thing from '@/components/Thing'

export default {
  name: 'marie-things',
  components: {
    'marie-thing': Thing
  },
  data () {
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
