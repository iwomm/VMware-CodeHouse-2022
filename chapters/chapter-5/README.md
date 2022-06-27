# Chapter 5 - Post a new todo item

## Objective

We will add a Vue component for submitting a todo item to the TODO API.
- Create a Vue component with the HTML form
- Use the form component in the App component


## Create the Vue component with HTML form

Create a new file named TodoForm.vue in the **components** folder with this contentL:

```
<template>
  <div id="todo-form">
    <form @submit.prevent="handleSubmit">

      <label>Todo</label>
      <input
        ref="first"
        type="text"
        :class="{ 'has-error': submitting && invalidValue }"
        v-model="todo.value"
        @focus="clearStatus"
        @keypress="clearStatus"
      >
      <label>Due Date</label>
      <input
        type="text"
        :class="{ 'has-error': submitting && invalidDueDate }"
        v-model="todo.due_date"
        @focus="clearStatus"
      >
      <p
        v-if="error && submitting"
        class="error-message"
      >❗Please fill out all required fields</p>
      <p
        v-if="success"
        class="success-message"
      >✅ Todo item successfully added</p>
      <button>Add Todo</button>
    </form>

  </div>
</template>

<script>
export default {
  name: 'todo-form',
  data() {
    return {
      error: false,
      submitting: false,
      success: false,
      todo: {
        value: '',
        due_date: '',
      }
    }
  },
  computed: {
    invalidValue() {
      return this.todo.value === ''
    },
    invalidDueDate() {
      return this.todo.due_date === ''
    },
  },
  methods: {
    handleSubmit() {
      this.clearStatus()
      this.submitting = true
      if (this.invalidValue || this.invalidDueDate) {
        this.error = true
        return
      }
      this.$emit('add:todo', this.todo)
      this.$refs.first.focus()
      this.todo = {
        value: '',
        due_date: '',
      }
      this.success = true
      this.error = false
      this.submitting = false
    },
    clearStatus() {
      this.success = false
      this.error = false
    }
  }}
</script>

<style scoped>
form {
  margin-bottom: 2rem;
}
[class*="-message"] {
  font-weight: 500;
}
.error-message {
  color: #d33c40;
}
.success-message {
  color: #32a95d;
}
</style>
```

The handleSubmit() method is called when the "Add Todo" button is clicked. After validating the inputs, the handleSubmit() method emits an event named `add:todo` with the current todo object as the event data. This event will be caught and handled by the parent component.

## Use the form component in the App component
Modify App.vue to be like this:
```
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
```

The notable changes are:
- The new component is imported with ```import TodoForm from '@/components/TodoForm.vue'```
- The form is added to the template with ```    <todo-form @add:todo="addTodo" />```. This code also says when the add:todo event is raised, call the addTodo fucntion.
- The addTodo() function makes a POST call to the TODO API and it refreshes the list with the latest list from the server. 

Now compile the Vue app, refresh the web page and play around with the new UI. Use the debugger to set some breakpoints if you want to see the running state of the app.