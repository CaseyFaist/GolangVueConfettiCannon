<script>
// import webSocket from "../mixins/Websocket"

export default { 
  name: "ClickMe",
  data() {
    return {
      count: 0
    }
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
    },
    increment () {
      this.count += 1
    }
  }
}
</script>


<template>
  <button class='counter' @click="increment(); sendMessage(this.count)">
  <!-- <button class='counter' @click="count++; clickButton(count)"> -->
      You clicked me {{ this.count }} times.
  </button>
</template>

<style>
.counter {
  color: forestgreen;
  font-weight: bold;
}
</style>