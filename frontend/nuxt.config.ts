import vuetify, { transformAssetUrls } from 'vite-plugin-vuetify'

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
    '@pinia-plugin-persistedstate/nuxt',
    (_options, nuxt) => {
      nuxt.hooks.hook('vite:extendConfig', (config) => {
        // @ts-expect-error
        config.plugins.push(vuetify({ autoImport: true }))
      })
    },
    
    // 'nuxt-auth-utils',
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
    server: {
      port: 3000,
      strictPort: true,
      hmr: {
        protocol: "ws",
        host: "localhost",
        port: 24687,
        clientPort: 3000,
    },
  },
  },

  vuetify: {
    /* vuetify options */
    vuetifyOptions: {
      icons: {
        defaultSet: 'mdi',
      },
    },
    moduleOptions: {
      /* nuxt-vuetify module options */
      treeshaking: true,
      useIconCDN: true,
      /* vite-plugin-vuetify options */
      styles: true,
      autoImport: true,
      // comment this line to check drawer behavior
      ssrClientHints: {
        viewportSize: true,
      },
    },
  },

  piniaPersistedstate: {
    cookieOptions: {
      sameSite: 'strict',
    },
    storage: 'cookies'
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
