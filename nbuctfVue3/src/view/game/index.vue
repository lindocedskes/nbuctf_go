<template>
  <div id="challenges">
    <div style="height: auto" class="container">
      <div class="title">
        <h1>
          <el-icon><Aim /></el-icon>
          {{ challengeType }}
        </h1>
        <el-divider content-position="right"> </el-divider>
      </div>
      <el-tabs tab-position="left" style="height: 100%" v-model="challengeType">
        <el-tab-pane
          :name="name"
          v-for="(icon, name) in challengesIcon"
          :key="name"
        >
          <template v-slot:label>
            <span>{{ name }}</span>
          </template>
        </el-tab-pane>
        <challenge
          :challengeType="challengeType"
          style="min-height: 400px"
          v-if="reload"
        ></challenge>
      </el-tabs>
    </div>
  </div>
</template>

<script setup>
// import { getGameList } from '@/api/game.js'
import { ref, watch } from 'vue'
import challenge from '@/components/challenge/index.vue'

const reload = ref(true)
const challengesIcon = {
  All: 'mobile-phone',
  Crypto: 'lock',
  Web: 'monitor',
  Misc: 'connection',
  Pwn: 'cpu',
  Reverse: 'refresh'
}
const challengeType = ref('All')

// const gameData = ref([])
// // 查询
// const getGameData = async () => {
//   const datas = await getGameList({})
//   if (datas.code === 0) {
//     gameData.value = datas.data.gameList
//   }
// }

// const initPage = async () => {
//   getGameData()
//   console.log(challengeType.value)
// }
// initPage()

watch(
  () => challengeType.value,
  () => {
    // ...
  }
)
</script>

<style lang="scss" scoped>
.form-dialog {
  max-width: 500px;
}
.container {
  text-align: center;
  margin: 0 auto;
  max-width: 85%;
}
.title {
  margin: 0 auto;
  max-width: 65%;
}
</style>
