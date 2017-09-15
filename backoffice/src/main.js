// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueNativeSock from 'vue-native-websocket'
import Vuetify from 'vuetify'
import store from './store'
import 'vuetify/dist/vuetify.min.js'
import 'vuetify/dist/vuetify.min.css'

Vue.use(Vuetify)
Vue.use(VueNativeSock, process.env.WS_URL, { store, format: 'json' })
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: { App }
})
