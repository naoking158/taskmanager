<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
</script>


<template>
  <v-container>
    <h1 class="text-h2 mb-4">Welcome to Task Manager</h1>
    <p v-if="authStore.token">You are logged in.</p>
    <p v-else>Please login or register to start managing your taskts.</p>
  </v-container>
</template>
