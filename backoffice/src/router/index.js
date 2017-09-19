import Vue from 'vue'
import VueRouter from 'vue-router'
import Things from '@/components/Things'
import ThingForm from '@/components/ThingForm'
import Speech from '@/components/Speech'
import Curves from '@/components/Curves'

Vue.use(VueRouter)

export default new VueRouter({
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
      path: '/thing-form/:id?',
      name: 'marie-thing-form-update',
      component: ThingForm
    },
    {
      path: '/speech-test',
      name: 'marie-speech',
      component: Speech
    },
    {
      path: '/curves',
      name: 'marie-curves',
      component: Curves
    }
  ]
})
