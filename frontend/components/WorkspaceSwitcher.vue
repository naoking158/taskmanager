<template>
  <div>
    <v-menu v-if="workspaces.length > 0" offset-y>
      <template v-slot:activator="{ props }">
        <v-btn v-bind="props" text style="text-transform: none;">
          {{ currentWorkspaceName }}
          <v-icon right>mdi-chevron-down</v-icon>
        </v-btn>
      </template>
      <v-list>
        <v-list-item
          v-for="workspace in workspaces"
          :key="workspace.name"
          @click="switchWorkspace(workspace)"
        >
          <v-list-item-title>{{ workspace.name }}</v-list-item-title>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item @click="showCreateWorkspaceDialog">
          <v-list-item-title class="text-primary">
            <v-icon left>mdi-plus</v-icon> Create New Workspace
          </v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
    <v-btn v-else color="primary" @click="showCreateWorkspaceDialog">
      Create Workspace
    </v-btn>
    
    <!-- ワークスペース作成ダイアログ -->
    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-card-title>Create New Workspace</v-card-title>
        <v-card-text>
          <v-text-field
            v-model="newWorkspaceName"
            label="Workspace Name"
            required
          ></v-text-field>
          <v-textarea
            v-model="newWorkspaceDescription"
            label="Description"
          ></v-textarea>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="dialog = false">Cancel</v-btn>
          <v-btn color="blue darken-1" text @click="createWorkspace">Create</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useWorkspaceStore } from '@/stores/workspace'
import type { Workspace } from '@/types'
import { useAPI } from '@/composables/useApi'

const workspaceStore = useWorkspaceStore()
const { workspaces, currentWorkspace } = storeToRefs(workspaceStore)

const currentWorkspaceName = computed(() => currentWorkspace.value.name || 'Select Workspace')

const dialog = ref(false)
const newWorkspaceName = ref('')
const newWorkspaceDescription = ref('')

const switchWorkspace = async (workspace: Workspace) => {
  await workspaceStore.setCurrentWorkspace(workspace)
  navigateTo('/tasks')
}

const showCreateWorkspaceDialog = () => {
  dialog.value = true
}

const createWorkspace = async () => {
  try {
    const { data } = await useAPI('/workspaces', {
      method: 'POST',
      body: {
        name: newWorkspaceName.value,
        description: newWorkspaceDescription.value
      }
    })
    // workspaceStore.addWorkspace(data)
    fetchWorkspaces()
    dialog.value = false
    newWorkspaceName.value = ''
    newWorkspaceDescription.value = ''

  } catch (error) {
    console.error('Failed to create workspace:', error)
    // エラー処理を追加（例：エラーメッセージの表示）
  }
}

async function fetchWorkspaces() {
  const { data, error } = await useAPI('/workspaces',{
    server: false,
  })
  if (error?.value) {
    console.log('Failed to fetch workspaces: ', error.value)
  }

  if (data.value == null) {
    workspaceStore.workspaces.value = []
  } else {
    workspaceStore.workspaces.value = data.value
  }

  if (workspaceStore.workspaces.value.length > 0 && !workspaceStore.currentWorkspace.value) {
    workspaceStore.currentWorkspace.value = workspaceStore.workspaces.value[0]
  }
}

// コンポーネントがマウントされたときにワークスペースを取得
onMounted(async () => {
  await fetchWorkspaces()
})
</script>
