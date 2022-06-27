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