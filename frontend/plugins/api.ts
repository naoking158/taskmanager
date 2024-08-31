import { $fetch } from 'ofetch'
import type { NuxtApp } from "nuxt/app"
import type { Pinia } from 'pinia'
import { useAuthStore } from "@/stores/auth.js"

export default defineNuxtPlugin((nuxtApp: NuxtApp) => {
  const config = useRuntimeConfig()
  const baseURL = config.public.apiBase
  const authStore = useAuthStore(nuxtApp.$pinia as Pinia)

  const $api = $fetch.create({
    baseURL: baseURL,
    cache: 'reload',
    onRequest({ request, options, error }) {

      if (authStore.token) {
        const headers = options.headers ||= {}
        if (Array.isArray(headers)) {
          headers.push(['Authorization', `Bearer ${authStore.token}`]) 
        } else if (headers instanceof Headers) {
          headers.set('Authorization', `Bearer ${authStore.token}`)
        } else {
          headers.Authorization = `Bearer ${authStore.token}`
        }
      }
    },
    async onResponseError({ response }) {
      if (response.status === 401) {
        await nuxtApp.runWithContext(() => navigateTo('/login'))
      }
    }
  })

  // Expose to useNuxtApp().$api
  return {
    provide: {
      api: $api,
    }
  }
})
