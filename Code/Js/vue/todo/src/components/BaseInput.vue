<template>
  <input type="text" v-bind:value='value' v-on="listeners">
</template>

<script>
export default {
  props: {
    value: String,
    default: ''
  },
  computed: {
    listeners () {
      return {
        // Pass all component listeners directly to input
        ...this.$listeners,
        // Override input listener to work with v-model
        input: event => this.$emit('input', event.target.value)
      }
    }
  }
}
</script>
