<template>
<canvas :data-dummyvalue="ctx &amp;&amp; redraw"/>
</template>

<script>

export default {
    name: 'a-canvas',
    props: ['width', 'height', 'draw'],
    data() { 
        return { ctx: null, }; 
        },
    mounted() { this.ctx = this.$el.getContext('2d'); },
    computed: {
        redraw() {
            let canvas = this.$el;
            canvas.width = this.width;
            canvas.height = this.height;
            this.ctx.clearRect(0, 0, this.width, this.height);
            this.ctx.save();
            this.draw(canvas, this.ctx);
            this.ctx.restore();
            return canvas
        },
    },
}
</script>

<style scoped>
body {
  background-color: black;
  font-family: "Signika Negative", sans-serif;
  color: white;
  text-align: center;
}
</style>