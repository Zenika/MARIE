<template>
  <v-toolbar>
    <v-toolbar-side-icon @click.native.stop="toggleDrawer"></v-toolbar-side-icon>
    <v-toolbar-title>
      MARIE
    </v-toolbar-title>
  </v-toolbar>
</template>

<script>
export default {
  name: 'marie-toolbar',
  methods: {
    toggleDrawer: function () {
      this.$emit('toggleDrawer')
    }
  }
}
</script>
