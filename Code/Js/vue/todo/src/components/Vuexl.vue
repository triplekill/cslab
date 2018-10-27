<template>
  <div>
    <p v-on:click='count++'> {{ count }} </p>
    <button v-on:click='count++'>++</button>
  </div>
</template>

<script>
import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    count: 0
  },
  mutations: {// 不可以异步
    increment (state, payload) {
      state.count += payload.amount
    }
  },
  actions: {// 可以异步
    increment (context, payload) {
      context.commit('increment', payload)
    }
  }
})

store.commit('increment', {amount: 2})
store.commit({type: 'increment', amount: 3})
store.dispatch({type: 'increment', amount: 6})

console.log(store.state.count) // -> 1

export default {
  name: 'Vuexl',
  data: function () {
    return {
      count: 0
    }
  }
}
</script>

<style scoped>
</style>
