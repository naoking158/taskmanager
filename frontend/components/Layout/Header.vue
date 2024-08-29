<template>
  <v-app-bar app>
    <v-app-bar-title>Task Manager</v-app-bar-title>
    <v-spacer></v-spacer>
    <v-btn to="/">Home</v-btn>
    <v-btn v-if="!authStore.token" to="/login text">Login</v-btn>
    <v-btn v-if="!authStore.token" to="/register" text>Register</v-btn>
    <v-btn v-if="authStore.token" to="/logout" text>Logout</v-btn>
  </v-app-bar>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const logout = () => {
  authStore.logout()
  navigateTo('/login')
}
</script>
