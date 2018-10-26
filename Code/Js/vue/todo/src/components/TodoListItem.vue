<template>
  <div>
    <li>{{ item.text }}
      <button @click="$emit('remove', item.id)">
        X
      </button>
    </li>

  </div>
</template>

<script>
export default {
  props: {
    item: {
      type: Object
    }
  }
}
</script>

<style scoped>
</style>
