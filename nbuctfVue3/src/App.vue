<template>
  <div id="app">
    <!-- 直接根据路由加载对应组件 -->
    <router-view></router-view>
  </div>
</template>
<script setup>
defineOptions({ name: 'App' })
</script>
<style scoped></style>
