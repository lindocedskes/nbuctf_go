import axios from 'axios'
import { useUserStore } from '@/pinia'
import router from '@/router'
import { ElMessage } from 'element-plus'

const baseURL = ''

const instance = axios.create({
  // TODO 1. 基础地址，超时时间
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 10000
})
// 请求拦截
instance.interceptors.request.use(
  (config) => {
    // TODO 2. 携带token
    const useStore = useUserStore()
    if (useStore.token) {
      config.headers.Authorization = useStore.token // 每次请求 携带token
    }

    return config
  },
  (err) => Promise.reject(err)
)
// 响应拦截
instance.interceptors.response.use(
  (res) => {
    // TODO 4. 摘取核心响应数据
    if (res.data.code === 0) {
      return res
    }
    // TODO 3. 处理业务失败
    // 处理业务失败, 给错误提示，抛出错误
    ElMessage.error(res.data.message || '服务异常')
    return Promise.reject(res.data)
  },
  (err) => {
    // TODO 5. 处理401错误
    // 错误的特殊情况 => 401 权限不足 或 token 过期 => 拦截到登录
    if (err.response?.status === 401) {
      router.push('/login')
    }

    // 错误的默认情况 => 只要给提示
    ElMessage.error(err.response.data.message || '服务异常')
    return Promise.reject(err)
  }
)

export default instance
export { baseURL }
