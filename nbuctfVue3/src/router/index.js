import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL), // import.meta.env.BASE_URL 是路由的基准地址，默认是 ’/‘。vite.config.js  添加配置项  base: '/' 。
  routes: []
})

export default router
