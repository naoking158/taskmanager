<template>
  <v-container>
    <h1 class="text-h4 mb-4">Tasks</h1>
    <v-alert v-if="!currentWorkspace" type="warning" class="mb-4">
      Please select a workspace to view tasks.
    </v-alert>
    <TaskList v-else :key="tasks" :tasks="tasks" />
  </v-container>
</template>

<script setup>
import { onMounted, watch } from 'vue'
import { storeToRefs } from 'pinia'

import TaskList from '@/components/Task/TaskList.vue'
import { useTasks } from '@/composables/useTask'
import { useWorkspaceStore } from '@/stores/workspace'

const workspaceStore = useWorkspaceStore()
const { currentWorkspace } = storeToRefs(workspaceStore)

const { data: tasks, refresh } = useTasks(currentWorkspace, { immediate: false })

watch(currentWorkspace, async () => {
  if (currentWorkspace.value?.id) {
    await refresh()
  }
})

onMounted(async () => {
  if (currentWorkspace.value?.id) {
    await refresh()
  }
})
</script>
