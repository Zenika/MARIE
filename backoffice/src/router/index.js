import Vue from 'vue'
import Router from 'vue-router'
import Things from '@/components/Things'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'marie-things',
      component: Things
    }
  ]
})
