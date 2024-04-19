<!-- 根据当前的路由渲染相应的组件，并应用过渡效果，同时保持某些组件的状态 -->
<template>
  <div>
    <router-view v-slot="{ Component }">
      <transition mode="out-in" name="el-fade-in-linear">
        <keep-alive :include="routerStore.keepAliveRouters">
          <component :is="Component" />
        </keep-alive>
      </transition>
    </router-view>
  </div>
</template>

<script setup>
import { useRouterStore } from '@/pinia/modules/router'
const routerStore = useRouterStore()

defineOptions({
  name: 'admin'
})
</script>
