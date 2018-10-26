<template>
  <ul v-if="items.length">
    <TodoListItem v-for="item in items" v-bind:key="item.id" :item="item"/>
  </ul>
</template>

<script>
import TodoListItem from './TodoListItem'
export default {
  props: {
    items: {
      type: Array
    }
  },
  components: {
    TodoListItem
  }
}
</script>

<style scoped>
</style>
