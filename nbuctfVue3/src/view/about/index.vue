<template>
  <div id="about">
    <div class="container">
      <h1 style="text-align: center; color: #0ea5e9" class="text-xl">关于</h1>
      <el-divider content-position="right">
        <p style="color: #8f8f8f">N8US3C简介</p>
      </el-divider>
      <div class="about">
        <p>N8US3C个性化CTF靶场。</p>
        <p>前端使用VUE3+element-Plus,后端使用golang。</p>
        <p>靶机使用docker，通过k3s实现管理。</p>
        <p style="float: right">1u0Y_T</p>
      </div>
    </div>
  </div>
</template>

<script setup></script>

<style lang="scss" scoped>
.container {
  margin: 0 auto;
  max-width: 45%;
}
.about {
  text-indent: 35px;
}
</style>
