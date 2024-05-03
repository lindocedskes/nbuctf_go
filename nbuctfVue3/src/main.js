//按需引入el组件样式，才能正常显示组件
import 'element-plus/es/components/message/style/css'
import 'element-plus/es/components/loading/style/css'
import 'element-plus/es/components/notification/style/css'
import 'element-plus/es/components/message-box/style/css'
// import 'element-plus/es/components/list/style/css'

import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import '@/style/main.scss'
import pinia from '@/pinia/index'

import '@/permission' //路由守卫，每次请求前预处理、动态路由持久化、权限控制
import { register } from './core/global'

// @description 导入加载进度条，防止首屏加载时间过长，用户等待
import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'
Nprogress.configure({ showSpinner: false, ease: 'ease', speed: 500 })
Nprogress.start()
// 无需在这块结束，会在路由中间件中结束此块内容
import DataVVue3 from '@kjgl77/datav-vue3'

const app = createApp(App)
register(app) //注册全局变量 $N8US3C
app.use(pinia)
app.use(router)
app.use(DataVVue3)

app.mount('#app')
