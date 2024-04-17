import axios from 'axios'
import { useUserStore } from '@/pinia'
import router from '@/router'
import { ElMessage } from 'element-plus'
import { emitter } from '@/utils/bus.js'

const baseURL = ''

const service = axios.create({
  // TODO 1. 基础地址，超时时间
  baseURL: `${import.meta.env.VITE_BASE_PATH}:${import.meta.env.VITE_SERVER_PORT}`,
  timeout: 10000
})

let acitveAxios = 0
let timer
const showLoading = () => {
  acitveAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (acitveAxios > 0) {
      emitter.emit('showLoading')
    }
  }, 400)
}

const closeLoading = () => {
  acitveAxios--
  if (acitveAxios <= 0) {
    clearTimeout(timer)
    emitter.emit('closeLoading')
  }
}
// 请求拦截
service.interceptors.request.use(
  (config) => {
    if (!config.donNotShowLoading) {
      showLoading()
    }
    const userStore = useUserStore()
    config.headers = {
      // 请求头携带token
      'Content-Type': 'application/json',
      'x-token': userStore.token,
      'x-user-id': userStore.userInfo.ID,
      ...config.headers
    }
    return config
  },
  (error) => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response 拦截器
service.interceptors.response.use(
  (response) => {
    console.log('response:', response.data) // for debug
    const userStore = useUserStore()
    if (!response.config.donNotShowLoading) {
      closeLoading()
    }
    if (response.headers['new-token']) {
      userStore.setToken(response.headers['new-token'])
    }
    if (response.data.code === 0 || response.headers.success === 'true') {
      // 服务端定义的响应code码为0时请求成功，仅返回需要的部分response.data
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg)
      }
      return response.data
    } else {
      // 服务端定义的响应code码为非0时请求失败，弹窗提示错误，弹错
      ElMessage({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg),
        type: 'error'
      })
      //如果返回的数据中有reload字段，表示需要重新登录
      if (response.data.data && response.data.data.reload) {
        userStore.token = ''
        localStorage.clear() //清除浏览器的本地存储
        router.push({ name: 'Login', replace: true }) //导航到名为 'Login' 的路由。replace: true 表示这个导航将不会留下历史记录
      }
      return response.data.msg ? response.data : response
    }
  },
  (error) => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }

    if (!error.response) {
      ElMessageBox.confirm(
        `
        <p>检测到请求错误</p>
        <p>${error}</p>
        `,
        '请求报错',
        {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '稍后重试',
          cancelButtonText: '取消'
        }
      )
      return
    }

    switch (error.response.status) {
      case 500:
        ElMessageBox.confirm(
          `
        <p>检测到接口错误${error}</p>
        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>
        `,
          '接口报错',
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '清理缓存',
            cancelButtonText: '取消'
          }
        ).then(() => {
          const userStore = useUserStore()
          userStore.ClearStorage()
          router.push({ name: 'Login', replace: true })
        })
        break
      case 404:
        ElMessageBox.confirm(
          `
          <p>检测到接口错误${error}</p>
          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>
          `,
          '接口报错',
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '我知道了',
            cancelButtonText: '取消'
          }
        )
        break
    }

    return error
  }
)
export default service

export { baseURL }
