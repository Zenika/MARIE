<template>
  <div class="hello">
    <h1>Things</h1>
    <ul>
      <li v-for="thing in things" :key="thing.name">
        {{thing.name}}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'hello',
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
