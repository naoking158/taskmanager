export default defineNuxtConfig({
  experimental: {
    writeEarlyHints: false,
  },
  
  unhead: {
    renderSSRHeadOptions: {
      omitLineBreaks: false
    }
  },
  
  devtools: { enabled: true },
  css: ['vuetify/lib/styles/main.sass', '@mdi/font/css/materialdesignicons.min.css'],

  modules: [
    '@pinia/nuxt',
    'nuxt-auth-utils',
  ],
  
  build: {
    transpile: ['vuetify'],
  },

  vite: {
    define: {
      'process.env.DEBUG': false,
    },
    optimizeDeps: {
      exclude: ['fsevents'],
    },
  },

  // vite: {
  //   define: {
  //     'process.env.DEBUG': false,
  //   },
  // },

  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080/api/v1',
    }
  },

  compatibilityDate: '2024-08-27',
})
