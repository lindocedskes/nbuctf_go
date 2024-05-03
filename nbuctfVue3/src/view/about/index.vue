<template>
  <div id="about">
    <div class="container backgroundForm px-10 py-5">
      <h1 style="text-align: center; color: #0ea5e9" class="text-xl">关于</h1>
      <el-divider content-position="right">
        <p style="color: #8f8f8f">NBUCTF简介</p>
      </el-divider>
      <div class="about text-cyan-300">
        <p>NBUCTF个性化CTF靶场。</p>
        <p>前端使用VUE3+element-Plus+Datav+tailwindcss等</p>
        <p>后端使用golang+gin+gorm+casbin等</p>

        <p>靶机使用docker，通过k3s实现管理。</p>
        <p style="float: right">1u0Y_T</p>
      </div>
    </div>
  </div>
</template>

<script setup></script>

<style lang="scss" scoped>
.container {
  margin: auto;
  max-width: 45%;
}
.about {
  text-indent: 35px;
}
.backgroundForm {
  height: 80vh;
  background-image: url('/src/assets/form_bg.png');
  background-size: 115% 110%;
  background-position: center;
}
</style>
