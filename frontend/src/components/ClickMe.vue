<script>
export default { 
  name: "ClickMe",
  data() {
    return {
      count: 0
    }
  }
}
</script>

<template>
  <button @click="count++">
      You clicked me {{ count }} times.
  </button>
</template>

<style>
.greeting {
  color: red;
  font-weight: bold;
}
</style>