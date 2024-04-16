//路由守卫，每次请求前和后处理、动态路由持久化、权限控制
import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router'
import getPageTitle from '@/utils/page'
import router from '@/router'
import Nprogress from 'nprogress'

const whiteList = ['Login', 'Init']

//获取用户的路由信息，并将这些路由添加到路由器中。实现已经登录的用户动态路由持久化
const getRouter = async (userStore) => {
  const routerStore = useRouterStore()
  await routerStore.SetAsyncRouter()
  await userStore.GetUserInfo()
  const asyncRouters = routerStore.asyncRouters
  console.log('动态路由持久化：', asyncRouters)
  asyncRouters.forEach((asyncRouter) => {
    router.addRoute(asyncRouter) //路由配置对象，根据path和component属性动态加载对应组件
  })
}
//处理需要保持活动状态的路由
async function handleKeepAlive(to) {
  if (to.matched.some((item) => item.meta.keepAlive)) {
    if (to.matched && to.matched.length > 2) {
      for (let i = 1; i < to.matched.length; i++) {
        const element = to.matched[i - 1]
        //如果 element 的名称是 'layout'，那么从 to.matched 中移除这个元素，并递归调用
        if (element.name === 'layout') {
          to.matched.splice(i, 1)
          await handleKeepAlive(to)
        }
        // 如果没有按需加载完成则等待加载
        if (typeof element.components.default === 'function') {
          await element.components.default()
          await handleKeepAlive(to)
        }
      }
    }
  }
}
//路由守卫，它在每次路由跳转之前执行。由于动态添加的路由在页面刷新时会被清除，要持久化存储路由，需要在路由跳转之前获取动态路由。
router.beforeEach(async (to, from) => {
  const routerStore = useRouterStore()
  Nprogress.start() // 开启 Nprogress 进度条
  const userStore = useUserStore()
  //to.matched包含了即将要进入的路由的所有嵌套路径片段的路由记录
  to.meta.matched = [...to.matched] //to.meta是一个可以用来存储路由元信息的对象。
  handleKeepAlive(to) // 处理需要保持活动状态的路由
  const token = userStore.token
  // 在白名单中的判断情况
  document.title = getPageTitle(to.meta.title, to) //在浏览器标签页上显示的标题
  if (to.meta.client) {
    return true
  }
  // 如果在白名单（/Login）中
  if (whiteList.indexOf(to.name) > -1) {
    //检查用户的 token 是否存在
    if (token) {
      if (!routerStore.asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        await getRouter(userStore)
      }
      // token 可以解析但是却是不存在的用户 id 或角色 id 会导致无限调用
      if (userStore.userInfo?.authority?.defaultRouter != null) {
        if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
          return { name: userStore.userInfo.authority.defaultRouter }
        } else {
          return { path: '/layout/404' }
        }
      } else {
        // 强制退出账号
        userStore.ClearStorage()
        return {
          name: 'Login',
          query: {
            redirect: document.location.hash //重定向到当前的 URL #号后哈希值。
          }
        }
      }
    } else {
      return true
    }
  } else {
    // 不在白名单中并且已经登录的时候
    if (token) {
      // 添加flag防止多次获取动态路由和栈溢出
      if (!routerStore.asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        await getRouter(userStore)
        if (userStore.token) {
          if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
            return { ...to, replace: true }
          } else {
            return { path: '/layout/404' }
          }
        } else {
          return {
            name: 'Login',
            query: { redirect: to.href } //重定向到登录页
          }
        }
      } else {
        if (to.matched.length) {
          return true
        } else {
          return { path: '/layout/404' }
        }
      }
    }
    // 不在白名单中并且未登录的时候
    if (!token) {
      return {
        name: 'Login',
        query: {
          redirect: document.location.hash //重定向到当前的 URL #号后哈希值。
        }
      }
    }
  }
})
//路由跳转之后执行
router.afterEach(() => {
  // 路由加载完成后关闭进度条
  document.getElementsByClassName('main-cont main-right')[0]?.scrollTo(0, 0)
  Nprogress.done()
})

router.onError(() => {
  // 路由发生错误后销毁进度条
  Nprogress.remove()
})
