<template>
  <div class="marie-things">
    <div v-if="!things" class="text-xs-center">
      <h2>Application currently not available</h2>
      <v-btn @click.native="getThings">Retry</v-btn>
    </div>
    <v-layout row wrap>
      <v-flex xs3 v-for="thing in things" :key="thing.id" >
        <marie-thing :thing="thing"></marie-thing>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Thing from '@/components/thing/Thing'

export default {
  name: 'marie-things',
  components: {
    'marie-thing': Thing
  },
  data: function () {
    return {
      message: ''
    }
  },
  created () {
    this.getThings()
  },
  methods: {
    getThings: function () {
      this.$store.dispatch('getAllThings')
    }
  },
  computed: mapGetters([
    'things'
  ])
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
