import { asyncRouterHandle } from '@/utils/asyncRouter'
import { emitter } from '@/utils/bus.js'
import { asyncMenu } from '@/api/menu'
import { defineStore } from 'pinia'
import { ref } from 'vue'

const notLayoutRouterArr = [] // 无需layout的路由，即顶级路由
const keepAliveRoutersArr = [] // 记录需要keep-alive的路由
const nameMap = {} // 路由name ->  路由对象的映射

// 格式化路由，formatRouter(asyncRouter, routeMap)，asyncRouter是从后台获取的动态路由，routeMap是一个空对象
//构造routeMap，将menus的name属性作为键，将menus的值包根据父路由格式化（统一并赋给routeMap
const formatRouter = (routes, routeMap, parent) => {
  routes &&
    routes.forEach((item) => {
      item.parent = parent // 按父路由
      item.meta.btns = item.btns
      item.meta.hidden = item.hidden
      if (item.meta.defaultMenu === true) {
        //如果没有父路由，即顶级路由，将其 path 改为 /，否则不变
        if (!parent) {
          item = { ...item, path: `/${item.path}` }
          notLayoutRouterArr.push(item) //将顶级路由添加到 notLayoutRouterArr 数组中
        }
      }
      routeMap[item.name] = item //按menus的name属性作为键，将menus的值赋给routeMap
      //递归调用自身来格式化子路由
      if (item.children && item.children.length > 0) {
        formatRouter(item.children, routeMap, item)
      }
    })
}
//找出需要保持活动状态的路由，Vue.js 中使用 <keep-alive> 标签包裹的路由。这种路由的组件在被切换到后台后，不会被销毁，而是被缓存起来，提高了页面的切换效率。
const KeepAliveFilter = (routes) => {
  routes &&
    routes.forEach((item) => {
      // 子菜单中有 keep-alive 属性的，父菜单也必须 keep-alive，否则无效。这里将子菜单中有 keep-alive 的父菜单也加入。
      if (
        (item.children && item.children.some((ch) => ch.meta.keepAlive)) ||
        item.meta.keepAlive
      ) {
        item.component &&
          item.component().then((val) => {
            keepAliveRoutersArr.push(val.default.name) //将需要keep-alive的路由添加到keepAliveRoutersArr数组中
            nameMap[item.name] = val.default.name //将路由name映射到nameMap中
          })
      }
      //递归调用自身来格式化子路由
      if (item.children && item.children.length > 0) {
        KeepAliveFilter(item.children)
      }
    })
}

export const useRouterStore = defineStore('router', () => {
  const keepAliveRouters = ref([]) // 空的 队列数组
  const asyncRouterFlag = ref(0) // 异步路由标志
  const setKeepAliveRouters = (history) => {
    const keepArrTemp = []
    //forEach 是数组的一个方法，item 是 history 数组中的每一个元素
    history.forEach((item) => {
      //item 对象的 name 属性是否在 nameMap 中
      if (nameMap[item.name]) {
        //存在nameMap中的路由，添加到keepArrTemp数组中
        keepArrTemp.push(nameMap[item.name])
      }
    })
    keepAliveRouters.value = Array.from(new Set(keepArrTemp))
  }
  emitter.on('setKeepAlive', setKeepAliveRouters)

  const asyncRouters = ref([])
  const routeMap = {}
  // 从后台获取动态路由
  const SetAsyncRouter = async () => {
    asyncRouterFlag.value++ // 异步路由标志加一
    const baseRouter = [
      {
        path: '/layout',
        name: 'layout',
        component: 'view/layout/index.vue',
        meta: {
          title: '底层layout'
        },
        children: []
      }
    ]
    //await 异步函数，在此期间，其他代码可以继续执行
    const asyncRouterRes = await asyncMenu() // ! 从后台获取动态路由，得到是Promise对象
    const asyncRouter = asyncRouterRes.data.menus

    //如果 asyncRouter 存在，添加一个新的路由定义，（重载在被访问时才加载组件
    asyncRouter &&
      asyncRouter.push({
        path: 'reload',
        name: 'Reload',
        hidden: true,
        meta: {
          title: '',
          closeTab: true
        },
        component: 'view/error/reload.vue' // 当路由的路径为 'reload' 时，渲染 'view/error/reload.vue' 这个组
      })
    formatRouter(asyncRouter, routeMap) //构造routeMap，asyncRouter是从后台获取的动态路由，将menus的name属性作为键，将menus的值包根据父路由格式化（统一并赋给routeMap
    baseRouter[0].children = asyncRouter //将 asyncRouter 作为 baseRouter（view/layout/index.vue） 的子路由。
    //如果 顶级路由个数不为 0，就将其所有元素添加到 baseRouter 的末尾
    if (notLayoutRouterArr.length !== 0) {
      baseRouter.push(...notLayoutRouterArr)
    }
    asyncRouterHandle(baseRouter) //得到baseRouter，动态导入 Vue 组件，并处理异步路由
    KeepAliveFilter(asyncRouter) // 找出需要保持活动状态的路由
    asyncRouters.value = baseRouter //更新，ref创建，任何依赖于 asyncRouters.value 的地方都会被自动更新
    console.log('最终动态路由：', asyncRouterRes)
    return true
  }

  return {
    asyncRouters,
    keepAliveRouters,
    asyncRouterFlag,
    SetAsyncRouter,
    routeMap
  }
})
