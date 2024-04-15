import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import '@/assets/main.scss'
import pinia from '@/pinia/index'

import '@/permission' //路由守卫，每次请求前预处理、动态路由持久化、权限控制

// @description 导入加载进度条，防止首屏加载时间过长，用户等待
import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'
Nprogress.configure({ showSpinner: false, ease: 'ease', speed: 500 })
Nprogress.start()
// 无需在这块结束，会在路由中间件中结束此块内容

const app = createApp(App)
app.use(pinia)
app.use(router)

app.mount('#app')
