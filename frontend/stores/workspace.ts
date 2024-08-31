import { defineStore } from 'pinia'
import { useAPI } from '@/composables/useApi'
import { Workspace } from '@/types'
import { useNuxtApp } from 'nuxt/app'


export const useWorkspaceStore = defineStore(
  'workspace',
  () => {
    const workspaces = ref<Workspace[]>([])
    const currentWorkspace = ref<Workspace | null>(null)
    // const nuxtApp = useNuxtApp()
    
    // async function fetchWorkspaces() {
    //   const { data, error } = await useAPI(
    //     '/workspaces',
    //     {
    //       getCachedData(key: string) {
    //         return nuxtApp.payload.data[key] || nuxtApp.static.data[key]
    //       },
    //     },
    //   )
    //   if (error?.value) {
    //     console.log('Failed to fetch workspaces: ', error.value)
    //   }

    //   console.log(data.value)
    //   console.log(error.value)

    //   if (data.value == null) {
    //     workspaces.value = []
    //   } else {
    //     workspaces.value = data.value
    //   }
                  
    //   if (workspaces.value.length > 0 && !currentWorkspace.value) {
    //     currentWorkspace.value = workspaces.value[0]
    //   }

    //   console.log(workspaces.value)
    //   console.log(currentWorkspace.value)
    // }

    function setCurrentWorkspace(workspace: Workspace) {
      currentWorkspace.value = workspace
    }

    function addWorkspace(workspace: Workspace) {
      workspaces.value.push(workspace.value)
      if (!currentWorkspace.value) {
        currentWorkspace.value = workspace.value
      }
    }

    return { workspaces, currentWorkspace, setCurrentWorkspace, addWorkspace }
    // return { workspaces, currentWorkspace, fetchWorkspaces, setCurrentWorkspace, addWorkspace }
  },
  {
    persist: true,
  }
)
