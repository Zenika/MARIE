<template>
  <v-app 
    standalone
  >
    <marie-nav :drawer="drawer"></marie-nav>
    <marie-toolbar v-on:toggleDrawer="toggleDrawer"></marie-toolbar>
    <main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </main>
    <!--<md-sidenav class="main-sidebar md-left md-fixed" md-swipeable ref="main-sidebar">
      <md-toolbar>
        <h1 class="md-title">MARIE</h1>
      </md-toolbar>
  
      <div class="main-sidebar-links">
        <md-list class="md-dense">
          <md-list-item>
            <router-link exact to="/">Things</router-link>
            <div class="md-list-action">
              <md-button class="md-icon-button">
                <router-link to="/thing-form">
                  <md-icon>add</md-icon>
                </router-link>
              </md-button>
            </div>
          </md-list-item>
        </md-list>
      </div>
  
    </md-sidenav>-->
  </v-app>
</template>

<script>
import Nav from '@/components/Nav'
import Toolbar from '@/components/Toolbar'

export default {
  name: 'app',
  components: {
    'marie-nav': Nav,
    'marie-toolbar': Toolbar
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
