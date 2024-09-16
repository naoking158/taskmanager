<template>
  <v-card v-if="task">
    <v-card-title>
      <h1 class="text-h4">
        {{ task.title }}
        <v-chip :color="getStatusColor(task.status)" class="ml-2">
          {{ task.status }}
        </v-chip>
        <v-btn color="primary" @click="showEditTaskDialog">
          Edit
        </v-btn>
      </h1>

    </v-card-title>
    <v-card-text>
      <div class="description">
        <p>{{ task.description }}</p>
      </div>
      <v-row>
        <v-col cols="6">
          <p><strong>Due Date:</strong> {{ new Date(task.due_date).toLocaleString() }}</p>
        </v-col>
        <v-col cols="6">
          <p><strong>Assigned To:</strong> {{ task.assigned_to || 'Unassigned' }}</p>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import type { Task } from '@/types'

defineProps<{
  task: Task | undefined;
}>()

const getStatusColor = (status: string) => {
  switch (status) {
    case 'TODO': return 'blue'
    case 'In Progress': return 'orange'
    case 'DONE': return 'green'
    case 'On Hold': return 'grey'
    default: return 'black'
  }
}

// TODO
const showEditTaskDialog = () => {
  alert('clicked')
}
</script>

<style scoped>
.description {
  margin: 1.5em 0.5em;
  padding: 1.5em 0.5em;
  border-width: 1px;
  border-style: outset;
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>
