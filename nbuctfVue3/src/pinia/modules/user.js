import { getUserInfo, setSelfInfo } from '@/api/user'
import router from '@/router/index'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import cookie from 'js-cookie'
// 用户模块，export 导出函数，useUserStore常量被赋值为 defineStore 函数的返回值。组件可以
export const useUserStore = defineStore(
  'user',
  () => {
    const loadingInstance = ref(null)
    const userInfo = ref({
      uuid: '',
      nickName: '',
      headerImg: '',
      authority: {},
      sideMode: 'dark',
      activeColor: 'var(--el-color-primary)',
      baseColor: '#fff'
    })
    const token = ref(
      window.localStorage.getItem('token') || cookie.get('x-token') || ''
    )
    const setUserInfo = (val) => {
      userInfo.value = val
    }

    const setToken = (val) => {
      token.value = val
    }

    const NeedInit = () => {
      token.value = ''
      window.localStorage.removeItem('token')
      localStorage.clear()
      router.push({ name: 'Init', replace: true })
    }

    const ResetUserInfo = (value = {}) => {
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

    return {
      userInfo,
      token,
      NeedInit,
      ResetUserInfo,
      GetUserInfo,
      setToken,
      loadingInstance
    }
  },
  {
    persist: true // 持久化
  }
)
