import { defineStore } from 'pinia'
import { Workspace } from '@/types'
import { useAPI } from '@/composables/useApi'

export const useWorkspaceStore = defineStore(
  'workspace',
  () => {
    const workspaces = ref<Workspace[]>([])
    const currentWorkspace = ref<Workspace | null>(null)

    function setCurrentWorkspace(workspace: Workspace) {
      currentWorkspace.value = workspace
    }

    function setWorkspaces(fetchedWorkspaces: Workspace[]) {
      workspaces.value = fetchedWorkspaces
    }

    function addWorkspace(workspace: Workspace) {
      workspaces.value.push(workspace.value)
      if (!currentWorkspace.value) {
        currentWorkspace.value = workspace.value
      }
    }

    return { workspaces, currentWorkspace, setCurrentWorkspace, addWorkspace, setWorkspaces }
  },
  {
    persist: true,
  }
)
