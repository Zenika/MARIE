<template>
  <div>
    <h3>Test speech commands with text</h3>
    <v-text-field
      label="Speech command"
      v-model="command"
      @keypress.native.enter="send()"
    ></v-text-field>
    <v-btn @click.native="send()">Send</v-btn>
    <div v-if="result.action !== ''">
      <h4>Results</h4>
      Id : {{result.id}}<br />
      Action : {{result.action}}<br />
      Location : {{result.location}}<br />
      Type : {{result.type}}
    </div>
    <div v-if="result.getter !== ''">
      <h4>Results</h4>
      Id : {{result.id}}<br />
      getter : {{result.getter}}<br />
      Location : {{result.location}}
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'marie-speech',
  data: function () {
    return {
      command: ''
    }
  },
  methods: {
    send: function () {
      this.result.action = ''
      this.result.getter = ''
      const req = {
        type: 'speech',
        message: this.command
      }
      this.$socket.sendObj(req)
      this.command = ''
    }
  },
  computed: mapGetters([
    'result'
  ])
}
</script>

<style>

</style>
