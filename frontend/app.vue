<template>
  <v-app>
    <v-app-bar>
      <v-app-bar-title>Task Manager</v-app-bar-title>
      <v-spacer></v-spacer>
      <v-btn to="/" text>Home</v-btn>
      <v-btn to="/login" text>Login</v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <NuxtPage />
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
// スクリプトが必要な場合はここに記述
</script>
