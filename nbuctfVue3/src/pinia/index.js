import { createPinia } from 'pinia'
import persist from 'pinia-plugin-persistedstate' //持久化

const pinia = createPinia()
pinia.use(persist)

export default pinia

//统一导出
export * from './modules/user'
// export * from './modules/xxx'
