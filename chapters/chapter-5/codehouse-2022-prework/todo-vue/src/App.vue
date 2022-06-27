<template>
  <div id="app">
    <todo-form @add:todo="addTodo" />
    <todo-list :todos="todos" />
  </div>
</template>

<script>
import TodoForm from '@/components/TodoForm.vue'
import TodoList from '@/components/TodoList.vue'

const appData = {
  todos: []
}
  
export default {
  name: 'app',
  components: {
    TodoForm,
    TodoList
  },
  data() {
    return appData;
  },
  mounted: function() {
    this.getTodos();
  },
  methods: {
    getTodos: getTodos,
    addTodo: addTodo
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

async function addTodo(todo) {
  try {
    const response = await fetch('/api/todos', {
      method: 'POST',
      body: JSON.stringify(todo),
      headers: { "Content-type": "application/json; charset=UTF-8" }
    })
    console.log(response.status)
    await getTodos()
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