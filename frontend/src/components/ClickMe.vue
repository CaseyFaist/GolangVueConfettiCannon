<script>
// import webSocket from "../mixins/Websocket"

export default { 
  name: "ClickMe",
  data() {
    return {
      count: 0
    }
  },
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
</script>


<template>
  <button class='counter' @click="count++">
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