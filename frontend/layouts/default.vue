<template>
  <v-app>
    <Header />
    
    <v-main>
      <v-container>
        <slot />
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import Header from '@/components/Layout/Header.vue'
</script>
