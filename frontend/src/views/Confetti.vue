<template>
  <div class="confetti">
    <mouse-follow/>
  </div>
</template>

<script>
// @ is an alias to /src
// import ConfettiCannon from '@/components/ConfettiCannon.vue'
import MouseFollow from '../components/MouseFollow.vue'

export default {
  name: 'Confetti',
  components: {
    MouseFollow
  }
}
</script>