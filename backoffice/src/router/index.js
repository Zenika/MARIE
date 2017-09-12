import Vue from 'vue'
import Router from 'vue-router'
import Things from '@/components/Things'
import ThingForm from '@/components/ThingForm'
import Speech from '@/components/Speech'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'marie-things',
      component: Things
    },
    {
      path: '/thing-form',
      name: 'marie-thing-form',
      component: ThingForm
    },
    {
      path: '/thing-form/:id',
      name: 'marie-thing-form-update',
      component: ThingForm
    },
    {
      path: '/speech-test',
      name: 'marie-speech',
      component: Speech
    }
  ]
})
