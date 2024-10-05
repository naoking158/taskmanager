<template>
  <v-app-bar app hide-on-scroll>
    <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    <v-app-bar-title>Task Manager</v-app-bar-title>
    <v-spacer></v-spacer>
    <WorkspaceSwitcher v-if="authStore.token" />
    <v-btn v-if="!authStore.token" to="/login" text>Login</v-btn>
    <v-btn v-if="!authStore.token" to="/register" text>Register</v-btn>
    <v-btn v-if="authStore.token" @click="logout" text>Logout</v-btn>
  </v-app-bar>

  <ClientOnly>
    <v-navigation-drawer v-model="drawer" app temporary>
      <v-list>
        <v-list-item to="/" link>
          <template v-slot:prepend>
            <v-icon>mdi-home</v-icon>
          </template>
          <v-list-item-title>Home</v-list-item-title>
        </v-list-item>
        <v-list-item to="/tasks" link>
          <template v-slot:prepend>
            <v-icon>mdi-format-list-bulleted</v-icon>
          </template>
          <v-list-item-title>Tasks</v-list-item-title>
        </v-list-item>
        <v-list-item v-if="authStore.token" to="/users" link>
          <template v-slot:prepend>
            <v-icon>mdi-account</v-icon>
          </template>
          <v-list-item-title>Preference</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </ClientOnly>
  
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import WorkspaceSwitcher from '@/components/WorkspaceSwitcher.vue';

const drawer = ref(false)
const authStore = useAuthStore()

const logout = () => {
  authStore.logout()
  navigateTo('/login')
}
</script>
