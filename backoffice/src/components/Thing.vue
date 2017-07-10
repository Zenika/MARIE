<template>
  <md-layout class="marie-thing" md-flex="25">
    <md-card style="width: 100%">
      <md-card-header>
        <div class="md-title">{{thing.name}}</div>
        <div class="md-subhead">{{thing.type}}</div>
      </md-card-header>

      <md-card-content>
        <div>
          Protocol : {{thing.protocol}}
        </div>
        <div>
          Actions : <span v-for="action in thing.actions" :key="action.name">{{action.name}}</span>
        </div>
        <div>
          Getters : <span v-for="getter in thing.getters" :key="getter.name">{{getter.name}}</span>
        </div>
      </md-card-content>
      
    </md-card>
  </md-layout>
</template>

<script>
export default {
  name: 'marie-thing',
  props: ['thing']
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: 100;
}

.marie-thing {
  width: 100%;
}
</style>
