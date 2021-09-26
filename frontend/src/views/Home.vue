<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <ClickMe/>
  </div>
</template>

<script>
// @ is an alias to /src
import ClickMe from '@/components/ClickMe.vue'

export default {
  name: 'Home',
  components: {
    ClickMe
  }
}
</script>
