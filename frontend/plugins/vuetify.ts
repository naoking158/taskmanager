import type { NuxtApp } from 'nuxt/app'
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

export default defineNuxtPlugin((nuxtApp: NuxtApp) => {
  const vuetify = createVuetify({
    components,
    directives,
    ssr: true,
  })

  nuxtApp.vueApp.use(vuetify)
})
