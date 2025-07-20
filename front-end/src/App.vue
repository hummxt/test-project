<template>
  <div>
    <h2>Backend data test:</h2>
    <h3>Data:</h3>
    <pre>{{ data }}</pre>
    <div v-if="error" style="color: red">{{ error }}</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const data = ref<string | null>(null)
const error = ref('')

async function fetchData() {
  try {
    const response = await fetch('https://41aa5e9d0a17.ngrok-free.app/api/items')
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`)
    }
    const json = await response.json()
    data.value = JSON.stringify(json, null, 2)
  } catch (err: any) {
    console.error('xeta:', err)
    error.value = 'xeta'
  }
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
pre {
  background-color: #f4f4f4;
  padding: 10px;
  border-radius: 4px;
  white-space: pre-wrap;
}
</style>
