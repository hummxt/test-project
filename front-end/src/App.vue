<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface User {
  id: number
  name: string
  email: string
}

const users = ref<User[]>([])
const error = ref('')

async function fetchData() {
  try {
    const response = await fetch('http://localhost:8080/users')
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`)
    }
    const json = await response.json()
    users.value = json 
  } catch (err: any) {
    console.error('Error:', err)
    error.value = 'Could not fetch users.'
  }
}

onMounted(() => {
  fetchData()
})
</script>

<template>
  <div>
    <h2>Backend data test:</h2>

    <div v-if="error" style="color: red">{{ error }}</div>

    <ul v-if="users.length > 0">
      <li v-for="user in users" :key="user.id">
        <strong>{{ user.name }}</strong> — {{ user.email }}
      </li>
    </ul>

    <div v-else-if="!error">Loading...</div>
  </div>
</template>

<style scoped>
ul {
  list-style-type: none;
  padding: 0;
}

li {
  padding: 6px 0;
}
</style>
