// https://nuxt.com/docs/api/configuration/nuxt-config
// export default defineNuxtConfig({
//   compatibilityDate: '2024-04-03',
//   devtools: { enabled: true }
// })


export default defineNuxtConfig({
  devtools: { enabled: true },
  css: ['vuetify/lib/styles/main.sass', '@mdi/font/css/materialdesignicons.min.css'],

  build: {
    transpile: ['vuetify'],
  },

  experimental: {
    writeEarlyHints: false,
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

  compatibilityDate: '2024-08-27',
})
