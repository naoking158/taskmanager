<template>
  <v-btn color="primary" @click="dialog = !dialog">
    Create New Task
  </v-btn>

  <template v-if="tasks">
    <v-card>
      <v-card-title>
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
      </v-card-title>
      
      <v-data-table
        :headers="headers"
        :items="tasks"
        :search="search"
        :loading="tasks == null"
        :items-per-page="10"
        :fixed-header=true
      >
        <template v-slot:item.title="{ item }">
          <NuxtLink :to="{ name: 'tasks-id', params: { id: item.id } }">
            {{ item.title }}
          </NuxtLink>
        </template>
        <template v-slot:item.status="{ item }">
          <v-chip :color="getStatusColor(item.status)">
            {{ item.status }}
          </v-chip>
        </template>
        <template v-slot:item.due_date="{ item }">
          {{ new Date(item.due_date).toLocaleString() }}
        </template>
        <template v-slot:no-data>
          <v-alert v-if="tasks.length === 0" type="info" class="ma-2">
            No tasks found. Start by creating a new task!
          </v-alert>
        </template>
      </v-data-table>
    </v-card>
  </template>
</template>

<script setup lang="ts">
import { ref, computed, defineProps } from 'vue'

type Task = {
  id: string;
  title: string;
  status: string;
  due_date: Date;
}

defineProps<{
  tasks: Task[] | undefined,
}>()

const search = ref('')
const headers = [
  { title: 'Title', align: 'start', key: 'title' },
  { title: 'Status', align: 'start', key: 'status' },
  { title: 'Due Date', align: 'start', key: 'due_date' },
]

const getStatusColor = (status: string) => {
  switch (status) {
    case 'TODO': return 'blue'
    case 'In Progress': return 'orange'
    case 'DONE': return 'green'
    case 'On Hold': return 'grey'
    default: return 'black'
  }
}
const dialog = ref(false)
// const showNewTaskDialog = () => {
//   dialog.value = true
// }
</script>
