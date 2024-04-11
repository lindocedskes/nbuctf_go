import { defineStore } from 'pinia'
import { ref } from 'vue'

// 用户模块，export 导出函数，useUserStore常量被赋值为 defineStore 函数的返回值。组件可以
export const useUserStore = defineStore(
  'big-user',
  () => {
    const token = ref('') // 定义 token
    const setToken = (newToken) => (token.value = newToken) // 设置 token，接受一个参数 t，并将 t 的值赋给 token
    const removeToken = () => {
      token.value = ''
    }

    return {
      token,
      setToken,
      removeToken
    }
  },
  {
    persist: true // 持久化
  }
)
