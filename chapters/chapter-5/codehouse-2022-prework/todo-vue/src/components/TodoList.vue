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