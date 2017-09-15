<template>
  <v-snackbar :timeout="2000" :bottom="true" :right="true" v-model="snackbar">
    {{snackbarText}}
    <v-btn flat class="pink--text" @click.native="snackbar = false">Close</v-btn>
  </v-snackbar>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'marie-snackbar',
  data: function () {
    return {
      snackbar: false
    }
  },
  watch: {
    snackbarText (value) {
      if (value !== '') {
        this.snackbar = true
      }
    },
    snackbar (value) {
      if (!value) {
        this.$store.commit('SNACKBAR_CLOSED')
      }
    }
  },
  computed: mapGetters([
    'snackbarText'
  ])
}
</script>

<style>

</style>
