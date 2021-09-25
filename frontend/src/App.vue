<template>
  <div>
  <img alt="Vue logo" src="./assets/logo.png">
  <button v-on:click="sendMessage('hello')">Send Message</button>
  <ul>
    <li>
      <ClickMe/>
    </li>
  </ul>
  </div>
</template>

<script>
import ClickMe from './components/ClickMe.vue'

export default {
  name: 'App',
  components: {
    // HelloWorld,
    ClickMe
  },
  created: function() {
    console.log("Starting connection to WebSocket Server")
    this.connection = new WebSocket("ws://localhost:8080/ws")

    this.connection.onmessage = function(event) {
      console.log(event);
    }

    this.connection.onopen = function(event) {
      console.log(event)
      console.log("Successfully connected to the echo websocket server...")
    }
  },
  methods: {
    sendMessage: function(message) {
      console.log(message)
      console.log(this.connection);
      this.connection.send(message);
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
