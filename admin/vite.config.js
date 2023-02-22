import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import { VitePWA } from 'vite-plugin-pwa'
import tailwindcss from 'tailwindcss'
import autoprefixer from 'autoprefixer'

export default defineConfig({
  base: process.env.NODE_ENV === "production" ? "/{{__DIR__}}/" : "/",
  server:{
    host:"0.0.0.0",
    port:3000,
    proxy:{
      '/admin/api': {
        target: 'http://127.0.0.1:8989/',
        changeOrigin: true,
      },
    },
  },
  resolve:{
    alias:{
      "@":path.resolve(__dirname, "src"),
      'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js',
    }
  },
  build:{
    outDir:"../main/resources/admin",
    emptyOutDir: true,
    rollupOptions: {
      output:{
        manualChunks(id) {
          if (id.includes('node_modules')) {
            return id.toString().split('node_modules/')[1].split('/')[0].toString();
          }
        },
      }
    },
  },
  plugins: [
    vue(),
    VitePWA({
      registerType: 'autoUpdate',
      includeAssets: ['favicon.ico', 'apple-touch-icon.png', 'mask_icon.svg'],
      manifest:{
        name:"moss",
        short_name:"moss",
        description: 'Moss Administration',
        //theme_color: '#ffffff',
        icons: [
          {
            src: 'icon_192.png',
            sizes: '192x192',
            type: 'image/png'
          },
          {
            src: 'icon_512.png',
            sizes: '512x512',
            type: 'image/png'
          },
          {
            src: 'icon_512.png',
            sizes: '512x512',
            type: 'image/png',
            purpose: 'any maskable'
          }
        ]
      }
    }),
  ],
  css: {
    postcss: {
      plugins: [tailwindcss, autoprefixer],
    },
    preprocessorOptions: {
      less: {
        modifyVars: {
          'dark-gray-1': "#2a2f38",
          'dark-gray-2':"#383e49",
          'dark-gray-3': "#49505e",
          'dark-gray-4': "#505664",
          'dark-gray-5': "#545b69",


          'dark-color-border-1':'#2a2f38',
          'dark-color-border-2':'#383e49',
          'dark-color-border-3':'#49505e',
          'dark-color-border-4':'#575f6e',

          'color-fill-1':'#fafafc',

          'dark-color-fill-1':'#282d36',
          'dark-color-fill-2':'#313641',
          'dark-color-fill-3':'#373e49',
          'dark-color-fill-4':'#444c59',

          'dark-color-text-1':'#93a6b9',
          'dark-color-text-2':'#8597a9',
          'dark-color-text-3':'#77889a',
          'dark-color-text-4':'#687888',

          'dark-color-bg-1':'#21252b',
          'dark-color-bg-2':'#282d36',
          'dark-color-bg-3':'#393f4d',
          'dark-color-bg-4':'#3f4554',
          'dark-color-bg-5':'#3f4554',

          'dark-color-spin-layer-bg':'rgba(18,25,37,0.6)',

          'color-menu-light-bg':'#f0f1f6',
          'color-menu-dark-bg':'#21252b',
        },
        javascriptEnabled: true,
      }
    },
  },
})
