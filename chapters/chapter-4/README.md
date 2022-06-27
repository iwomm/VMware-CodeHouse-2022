# Chapter 4 - Display the todo list and make a Vue component

## Objective
We will modify the Vue app to:
- Display the todo list
- Make a Vue component for the todo list


## Display the todo list
First delete the file `todo-vue/src/components/HellowWorld.vue` since we are going to write our own todo UI. Then modify the App.vue file to look like this:

```
<template>
  <div id="app">
    <h1>To-Do List</h1>
    <table>
      <thead>
        <tr>
          <th>Item</th>
          <th>Due Date</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="todo in todos" :key="todo.id">
          <td>{{ todo.value }}</td>
          <td>{{ todo.due_date }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
const appData = {
  todos: []
}

export default {
  name: 'app',
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
```

In the template, the `v-for` directive loops through each item in the todos collection. The `getTodos()` method makes a GET request to our todo API at the url `/api/todos` . Build the Vue app and launch the Go app:

```
cd ~/src/codehouse-2022-prework/todo-vue

# rebuild the Vue app
rm -rf dist; npm run build

#launch the Go app (only necessary if the Go app is not already running)
cd ..
go run .
```
Open a browser and navigate to http://localhost:8090. The todo list should display.

Note next time you make a change to the Vue app, you only need to re-build the Vue app while keeps the Go app running.

## The getTodos() method
Some insteresting things are worth noting about the getTodos() method:
- The method uses [Javascript async/await keywords](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Asynchronous/Async_await). This is a improvements over Promise's syntax.  
- The method uses the [Fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch) for calling the TODO API. 

## Make a Vue component for the todo list
As we add more content and interactions to our UI applicatiohn, the code in App.vue will become too big and too difficult to read at a glance. To make the code easier to manage, we will move the todo list to a Vue component and let App.vue to consume the new component.

Create a new file named `TooList.vue` in the `todo-vue/components` folder. Edit its content to be like:

```
<template>
  <div id="todo-list">
    <h1>To-Do List</h1>
    <table>
      <thead>
        <tr>
          <th>Item</th>
          <th>Due Date</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="todo in todos" :key="todo.id">
          <td>{{ todo.value }}</td>
          <td>{{ todo.due_date }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'todo-list',
  props: {
    todos: Array,
  },
  data() {
    return null;
  }
}
</script>

<style scoped>
</style>
```


This code looks similar to that in the previous App.vue file, but with two notable changes:
- The code does not pull data from API. Instead, it declared a property named todos, which of Array type
- This component returns no data

Next modify the code in App.vue to be like this:

```
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

```  

Comparing to the previous version, these are the important changes:
- The code imports the TodoList component from the TodoList.vue: `import TodoList from '@/components/TodoList.vue'`
- The template is simplied with `<todo-list :todos="todos" />`. This places the TodoList component inside the main template. It also passes the `todos` property from `appData` to the `todos` property of the component

Compile the todo-vue app and refresh the web page. You will see nothing has changed on the page. But the better organized code brings several benefits:
- Smaller code file for easier maintennance
- A predictable pattern of modules allows other people get a quick start
- A reusable TODO list component
