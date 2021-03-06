// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import VueNativeSock from 'vue-native-websocket'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.js'
import 'vuetify/dist/vuetify.min.css'

Vue.use(Vuetify)
Vue.use(VueNativeSock, process.env.WS_URL, { format: 'json' })
Vue.prototype.$http = axios
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
