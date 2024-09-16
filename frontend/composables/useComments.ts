import { useAPI } from '@/composables/useApi'
import type { UseFetchOptions } from "nuxt/app"

export function useComments(
  taskID: string,
  options?: UseFetchOptions<T>,
) {
  return useAPI(`/tasks/${taskID}/comments`, {
    ...options,
    server: false,
  })
}
