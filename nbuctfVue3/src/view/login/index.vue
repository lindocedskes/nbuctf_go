<template>
  <div class="container2">
    <dv-border-box-11 style="height: 100%" title="NBUCTF">
      <LoginForm></LoginForm>
    </dv-border-box-11>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import LoginForm from './LoginForm.vue'

const box = ref(null)

// onMounted(() => {
//   setTimeout(() => {
//     box.value.initWH()
//   }, 100)
// })
</script>

<style scoped>
.full-screen {
  width: 100vw;
  height: 100vh;
  position: relative;
}
.container2 {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-image: url('/src/assets/login_bg.png');
  background-size: cover;
  background-position: center;
}
</style>
