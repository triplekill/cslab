import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Todo from '@/components/Todo'
import Syntex from '@/components/Syntex'
import Vuexl from '@/components/Vuexl'
import Vuetifyjsl from '@/components/Vuetifyjsl'
import Login from '@/components/Login'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/todo',
      name: 'todo',
      component: Todo
    },
    {
      path: '/test',
      name: 'test',
      component: Syntex
    },
    {
      path: '/vuex',
      name: 'vuex',
      component: Vuexl
    },
    {
      path: '/vuetifyjsl',
      name: 'vuetifyjsl',
      component: Vuetifyjsl
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    }
  ]
})
