const viewModules = import.meta.glob('../view/**/*.vue')

export const asyncRouterHandle = (asyncRouter) => {
  asyncRouter.forEach((item) => {
    if (item.component && typeof item.component === 'string') {
      if (item.component.split('/')[0] === 'view') {
        console.log('item.component', item.component)
        item.component = dynamicImport(viewModules, item.component) //动态导入 Vue 组件，component是组件位置如../view/layout/index.vue
      }
    }
    if (item.children) {
      asyncRouterHandle(item.children)
    }
  })
}
//将服务器返回的 component 字段（一个字符串）转换为一个 Vue 组件的函数。
//component：传入对应前端路径的字符串，如'view/dashboard/index.vue' 过滤成-> '../view/dashboard/index.vue'
function dynamicImport(dynamicViewsModules, component) {
  const keys = Object.keys(dynamicViewsModules)
  const matchKeys = keys.filter((key) => {
    const k = key.replace('../', '')
    return k === component
  })
  const matchKey = matchKeys[0]

  return dynamicViewsModules[matchKey]
}
