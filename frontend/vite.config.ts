import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import styleImport from 'vite-plugin-style-import'
import path from 'path'
import { htmlPlugin } from './setup'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    htmlPlugin(),
    styleImport({
      libs: [
        {
          libraryName: 'element-plus',
          esModule: true,
          ensureStyleFile: true,
          // resolveStyle: (name) => {
          //   return `element-plus/packages/theme-chalk/src/${name.slice(3)}.scss`
          // },
          resolveComponent: (name) => {
            return `element-plus/lib/${name}`
          },
        },
      ],
    }),
  ],
  // css: {
  //   preprocessorOptions: {
  //     scss: {
  //       additionalData: `@use "@/app/presenter/ui-theme/element/index.scss";`,
  //     },
  //   },
  // },
  server: {
    host: '0.0.0.0',
    // // Следует ли автоматически открывать в браузере
    // open: true,
    // // Открывать ли https
    // https: false,
    // // Серверный рендеринг
    // ssr: false,
    hmr: {
      path: '/sockjs-node',
      port: 8080,
    },
    // proxy: {},
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
})
