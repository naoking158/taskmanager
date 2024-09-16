import type { Ref } from 'vue'
import { Workspace } from '@/types/index'
import { useAPI } from '@/composables/useApi'
import type { UseFetchOptions } from "nuxt/app"

export function useTasks(
  currentWorkspace: Ref<Workspace>,
  options?: UseFetchOptions<T>,
) {
  return useAPI(() => `/workspaces/${currentWorkspace.value.id}/tasks`, {
    ...options, 
    server: false,
  })
}

export function useTaskByID(
  taskID: string,
  options?: UseFetchOptions<T>,
) {
  return useAPI(() => `/tasks/${taskID}`, {
    ...options,
    server: false,
  })
}
