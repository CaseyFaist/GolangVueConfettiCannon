<script>
export default { 
  name: "ClickMe",
  data() {
    return {
      count: 0
    }
  },
  mounted: function() {
    this.connectToWebsocket();
  },
  methods: {
    connectToWebsocket() {
        this.websocket = new WebSocket( "ws://localhost:8080/ws" );
        this.websocket.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
        this.websocket.addEventListener('message', (event) => { this.handleNewMessage(event) });
    },
    onWebsocketOpen() {
        console.log("Successfully connected to Websocket");        
    },
    handleNewMessage() {
        // let data = event.data;
        this.count += 1
        console.log(this.count)
    },
    sendMessage: function(message) {
      this.websocket.send(message);
    },
    increment () {
      this.count += 1
    }
  }
}
</script>


<template>
  <button class='counter' @click="sendMessage(this.count)">
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