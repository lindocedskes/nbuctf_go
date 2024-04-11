<script setup>
// 在 Vue3 CompositionAPI 中
// 1. 获取路由对象 router  useRouter
//    const router = useRouter()
// 2. 获取路由参数 route   useRoute
//    const route = useRoute()
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user.js'

const router = useRouter()
const route = useRoute()

const goList = () => {
  router.push('/list')
  console.log(router, route)
}

const userStore = useUserStore() //实例化对象
</script>

<template>
  <test-demo>直接引用components下的组件</test-demo>
  <button @click="$router.push('/home')">跳首页</button>
  <button @click="goList">跳列表页</button>

  <p>{{ userStore.token }}</p>
  <el-button @click="userStore.setToken('Bearer sidfhheiwhuvyweqrjh')">
    登录
  </el-button>
  <el-button @click="userStore.removeToken()">退出</el-button>
</template>

<style scoped></style>
