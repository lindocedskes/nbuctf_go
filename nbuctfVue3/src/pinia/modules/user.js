import { getUserInfo, login, setSelfInfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElLoading, ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'
import cookie from 'js-cookie'
import { useRouterStore } from './router' // 引入pina路由模块

// 用户模块，export 导出函数，useUserStore常量被赋值为 defineStore 函数的返回值。组件可以
export const useUserStore = defineStore(
  'user',
  () => {
    const loadingInstance = ref(null)
    const userInfo = ref({
      uuid: '',
      nickName: '',
      userName: '',
      headerImg: '',
      authority: {},
      sideMode: 'dark',
      activeColor: 'var(--el-color-primary)',
      baseColor: '#fff',
      email: '',
      phone: ''
    })
    const token = ref(
      window.localStorage.getItem('token') || cookie.get('x-token') || ''
    )
    const setUserInfo = (val) => {
      userInfo.value = val
    }

    const setToken = (val) => {
      console.log('调用setToken')
      token.value = val
    }
    //重置应用的状态并将用户重定向到初始化init页面。
    // const NeedInit = () => {
    //   token.value = ''
    //   window.localStorage.removeItem('token')
    //   localStorage.clear()
    //   router.push({ name: 'Init', replace: true })
    // }

    //更新 userInfo.value 的值
    const ResetUserInfo = (value = {}) => {
      // ... 展开运算符（...）来合并对象,value覆盖userInfo.value
      userInfo.value = {
        ...userInfo.value,
        ...value
      }
    }

    /* 获取用户信息*/
    const GetUserInfo = async () => {
      const res = await getUserInfo()
      if (res.code === 0) {
        setUserInfo(res.data.userInfo)
      }
      return res
    }

    /* 登录*/
    const LoginIn = async (loginInfo) => {
      loadingInstance.value = ElLoading.service({
        fullscreen: true,
        text: '登录中，请稍候...'
      })
      try {
        const res = await login(loginInfo) // 传入loginInfo，调用登录接口
        if (res.code === 0) {
          console.log(res.data)
          setUserInfo(res.data.user) // 设置用户信息
          setToken(res.data.token) // 设置token
          const routerStore = useRouterStore() // 实例化路由 store
          await routerStore.SetAsyncRouter() // 设置异步路由
          const asyncRouters = routerStore.asyncRouters // 获取处理后的动态路由
          asyncRouters.forEach((asyncRouter) => {
            router.addRoute(asyncRouter) // 添加动态路由，动态添加的路由在页面刷新时会被清除
          })
          //角色对应的路由
          if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
            ElMessage.error('请联系管理员进行授权')
          } else {
            await router.replace({
              name: userInfo.value.authority.defaultRouter // 跳转到默认路由
            })
          }

          loadingInstance.value.close() // 关闭loading

          const isWin = ref(/windows/i.test(navigator.userAgent)) // 判断是否为windows系统
          if (isWin.value) {
            window.localStorage.setItem('osType', 'WIN')
          } else {
            window.localStorage.setItem('osType', 'MAC')
          }
          return true
        }
      } catch (e) {
        loadingInstance.value.close()
      }
      loadingInstance.value.close()
    }

    /* 登出*/
    const LoginOut = async () => {
      const res = await jsonInBlacklist()
      if (res.code === 0) {
        await ClearStorage()
        router.push({ name: 'Login', replace: true })
        window.location.reload()
      }
    }
    /* 清理数据 */
    const ClearStorage = async () => {
      console.log('调用ClearStorage')
      token.value = ''
      sessionStorage.clear()
      localStorage.clear()
      cookie.remove('x-token')
    }
    /* 设置侧边栏模式*/
    const changeSideMode = async (data) => {
      const res = await setSelfInfo({ sideMode: data }) // 调用设置用户信息接口
      if (res.code === 0) {
        userInfo.value.sideMode = data
        ElMessage({
          type: 'success',
          message: '设置成功'
        })
      }
    }

    const mode = computed(() => userInfo.value.sideMode)
    const sideMode = computed(() => {
      if (userInfo.value.sideMode === 'dark') {
        return '#191a23'
      } else if (userInfo.value.sideMode === 'light') {
        return '#fff'
      } else {
        return userInfo.value.sideMode
      }
    })
    const baseColor = computed(() => {
      if (userInfo.value.sideMode === 'dark') {
        return '#fff'
      } else if (userInfo.value.sideMode === 'light') {
        return '#191a23'
      } else {
        return userInfo.value.baseColor
      }
    })
    const activeColor = computed(() => {
      return 'var(--el-color-primary)'
    })
    //当 token.value 的值改变时，会将新的 token.value 的值存储到浏览器的 localStorage 中
    watch(
      () => token.value,
      () => {
        console.log('token before:', window.localStorage.getItem('token'))
        window.localStorage.setItem('token', token.value)
        console.log('token after:', window.localStorage.getItem('token'))
      }
    )

    return {
      userInfo,
      token,
      // NeedInit,
      ResetUserInfo,
      GetUserInfo,
      LoginIn,
      LoginOut,
      changeSideMode,
      mode,
      sideMode,
      setToken,
      baseColor,
      activeColor,
      loadingInstance,
      ClearStorage
    }
  },
  {
    persist: true // 持久化
  }
)
