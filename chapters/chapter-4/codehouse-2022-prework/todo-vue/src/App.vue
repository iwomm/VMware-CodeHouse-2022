<template>
  <div id="app">
    <todo-list :todos="todos" />
  </div>
</template>

<script>
import TodoList from '@/components/TodoList.vue'

const appData = {
  todos: []
}
  
export default {
  name: 'app',
  components: {
    TodoList
  },
  data() {
    return appData;
  },
  mounted: function() {
    this.getTodos();
  },
  methods: {
    getTodos: getTodos
  }
}

async function getTodos() {
    try {
      const response = await fetch('api/todos')
      const data = await response.json()
      appData.todos = data.list
    } 
    catch (error) {
      console.error(error)
    }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  margin-top: 50px;
}
</style>