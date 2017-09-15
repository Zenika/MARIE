<template>
  <v-app 
    standalone
  >
    <marie-nav :drawer="drawer"></marie-nav>
    <marie-toolbar v-on:toggleDrawer="toggleDrawer"></marie-toolbar>
    <main>
      <marie-snackbar></marie-snackbar>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </main>
  </v-app>
</template>

<script>
import Nav from '@/components/Nav'
import Toolbar from '@/components/Toolbar'
import Snackbar from '@/components/Snackbar'

export default {
  name: 'app',
  components: {
    'marie-nav': Nav,
    'marie-toolbar': Toolbar,
    'marie-snackbar': Snackbar
  },
  data () {
    return {
      drawer: true
    }
  },
  methods: {
    toggleDrawer: function () {
      this.drawer = !this.drawer
    }
  }
}
</script>

<style lang="scss">
</style>
