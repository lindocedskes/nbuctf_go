<template>
  <div />
</template>

<script setup>
import { useRouter } from 'vue-router'
//定义了组件的名字为 'Reload'
defineOptions({
  name: 'Reload'
})

const router = useRouter() //获取了路由实例
router.go(-1) //类似于浏览器的后退按钮
</script>
