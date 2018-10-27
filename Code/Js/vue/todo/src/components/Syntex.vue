<template>
  <div>
    <p v-bind:text="msg" b-on:click='printl'>{{ msg }}</p>
  </div>
</template>

<script>
export default {
  name: 'Syntex',
  data: function () {
    return {
      msg: 'test'
    }
  },
  methods: {
    printl: function () {
      console.log('printl')
    }
  }
}
</script>

<style scoped>
* {
  background-color: aquamarine;
}
</style>
