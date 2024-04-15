<template>
  <div>
    <div>
      <div>
        <img src="../../assets/notFound.png" />
        <p>页面被神秘力量吸走了</p>
        <p>↓</p>
        <el-button @click="toDashboard">返回首页</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useUserStore } from '@/pinia/modules/user'
import { useRouter } from 'vue-router'

defineOptions({
  name: 'Error'
})

const userStore = useUserStore()
const router = useRouter()
const toDashboard = () => {
  router.push({ name: userStore.userInfo.authority.defaultRouter })
}
</script>
