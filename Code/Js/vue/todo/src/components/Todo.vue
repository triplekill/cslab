<template>

  <div>
    <BaseInput placeholder="New todo" @keydown.enter="addTodo" v-model="newTodoText" />
    <todo-list v-bind:items="items" @remove="removeTodo"/>
    <br>
    <p>{{ msg }}</p>
    </div>
</template>

<script>
import BaseInput from './BaseInput'
import TodoList from './TodoList'

let nextTodoId = 0

export default {
  name: 'Todo',
  data: function () {
    return {
      newTodoText: '',
      msg: 'Hello Todo',
      items: [
        {id: ++nextTodoId, text: 'todoa'},
        {id: ++nextTodoId, text: 'todob'}
      ]
    }
  },
  components: {
    BaseInput, TodoList
  },
  methods: {
    addTodo () {
      const trimmedText = this.newTodoText.trim()
      if (trimmedText) {
        this.items.push({
          id: ++nextTodoId,
          text: trimmedText
        })
        this.newTodoText = ''
      }
    },
    removeTodo (idToRemove) {
      this.items = this.items.filter(item => {
        console.log('remove')
        return item.id !== idToRemove
      })
    }
  }
}
</script>

<style scoped>
* {
  background-color: aquamarine;
}
</style>
