import type { UseFetchOptions } from "nuxt/app"

export function useAPI<T>(
  url: string | (() => string),
  options?: UseFetchOptions<T>,
) {
  return useFetch(url, {
    ...options,
    $fetch: useNuxtApp().$api
  })
}

// export const useApi = () => {
//   const config = useRuntimeConfig()
//   const baseURL = config.public.apiBase

//   const apiFetch = (endpoint: string, options = {}) => {
//     return useFetch(endpoint, {
//       baseURL,
//       ...options,
//     })
//   }

//   return {
//     apiFetch
//   }
// }
