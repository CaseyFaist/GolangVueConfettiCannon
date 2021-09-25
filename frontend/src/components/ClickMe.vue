<script>
export default { 
  name: "ClickMe",
  data() {
    return {
      count: 0
    }
  },
  // methods: {
  //   clickButton: function() {
  //       // $socket is [WebSocket](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket) instance
  //       this.$socket.send('some data')
  //       console.log(this.$socket)
  //       // or with {format: 'json'} enabled
  //       // this.$socket.sendObj({awesome: 'data'})
  //       this.count++
  //   }
  // }
}
</script>

<template>
  <button @click="clickButton">
      You clicked me {{ count }} times.
  </button>
</template>

<style>
.greeting {
  color: red;
  font-weight: bold;
}
</style>