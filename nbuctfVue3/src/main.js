import { createApp } from 'vue'
import { createPinia } from 'pinia'
import persist from 'pinia-plugin-persistedstate' //持久化

import App from './App.vue'
import router from './router'
import '@/assets/main.scss'

const app = createApp(App)
app.use(createPinia().use(persist))
app.use(createPinia())
app.use(router)

app.mount('#app')
