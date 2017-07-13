<template>
  <v-navigation-drawer class="pb-0" persistent absolute height="100%" clipped enable-resize-watcher v-model="drawer">
    <v-list dense>
      <router-link to="/">
        <v-list-tile class="mt-3">
          <v-list-tile-content>
              <v-list-tile-title>
                Things
              </v-list-tile-title>
          </v-list-tile-content>
          <v-list-tile-action>
            <router-link to="/thing-form">
              <v-icon>add_circle_outline</v-icon>
            </router-link>
          </v-list-tile-action>
        </v-list-tile>
      </router-link>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
  name: 'marie-nav',
  props: [ 'drawer' ]
}
</script>

<style scoped>
a {
  text-decoration: none
}
</style>
