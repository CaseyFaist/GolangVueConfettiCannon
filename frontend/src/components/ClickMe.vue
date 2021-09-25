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
    }
  // methods: {
  //   clickButton: function(data) {
  //       // // var webSocket = new WebSocket("ws://localhost:8080")
  //       // // webSocket.send(data)
  //       // // this.$webSocketsConnect()
  //       // webSocket.send(data)
  //       // // this.$socket.send(data);
  //       // console.log(data)
  //   }
  // }
  }
}
</script>


<template>
  <button class='counter' @click="count++; sendMessage(count)">
  <!-- <button class='counter' @click="count++; clickButton(count)"> -->
      You clicked me {{ count }} times.
  </button>
</template>

<style>
.counter {
  color: forestgreen;
  font-weight: bold;
}
</style>