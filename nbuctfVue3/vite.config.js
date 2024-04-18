import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import tailwindcss from 'tailwindcss'
import autoprefixer from 'autoprefixer'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    open: true,
    port: 8888, //前端运行端口
    // 前端代理，解决跨域问题
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8889', // 本地后端地址
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''), // 去掉/api
        secure: false
      }
    }
  },

  // server: {
  //   // 如果使用docker-compose开发模式，设置为false
  //   open: true,
  //   port: import.meta.env.VITE_CLI_PORT,
  //   proxy: {
  //     // 把key的路径代理到target位置
  //     [import.meta.env.VITE_BASE_API]: {
  //       // 需要代理的路径   例如 '/api' ->
  //       target: `${import.meta.env.VITE_BASE_PATH}:${import.meta.env.VITE_SERVER_PORT}/`, // 代理到 目标路径
  //       changeOrigin: false,
  //       rewrite: (path) =>
  //         path.replace(new RegExp('^' + import.meta.env.VITE_BASE_API), '')
  //     }
  //   }
  // },
  plugins: [
    vue(),
    AutoImport({
      resolvers: [ElementPlusResolver()]
    }),
    Components({
      resolvers: [ElementPlusResolver()]
    })
  ],
  css: {
    postcss: {
      plugins: [tailwindcss, autoprefixer]
    }
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
