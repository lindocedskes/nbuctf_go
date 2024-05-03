<template>
  <div id="challenges">
    <div style="height: auto" class="container">
      <div class="title flex flex-col items-center">
        <dv-decoration9
          class="text-green-300 font-semibold"
          style="width: 150px; height: 40px"
        >
          <h1>
            <el-icon><Aim /></el-icon>
            {{ challengeType }}
          </h1>
        </dv-decoration9>
        <dv-decoration3 style="width: 100%; height: 20px" />

        <!-- <el-divider content-position="right"> </el-divider> -->
      </div>
      <el-tabs
        tab-position="left"
        style="height: 100%"
        v-model="challengeType"
        type="border-card"
      >
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

::v-deep .el-tabs--border-card {
  //修改整体背景色与边框为透明
  background-color: transparent;
  border-color: transparent;
  .el-tabs__header {
    //修改标签背景色为透明，并且没有下边横线
    background-color: transparent;
    margin: 0;
    border-bottom: 0px solid transparent;
    .el-tabs__nav {
      //将标签整体轮廓设为绿色
      border: 0px solid #80ffff;
      .el-tabs__item {
        //标签内容颜色为白色，并加粗，每个标签轮廓为绿色
        color: #80ffff;
        font-weight: bold;
        border-left: 0px solid #80ffff;
      }
      .el-tabs__item.is-active {
        //标签点击时，背景色变为蓝色，字体变为白色，左右上边框变绿色，
        background-color: #1b4584;
        color: #e36ce9;
        border-right-color: #80ffff;
        border-left-color: #80ffff;
        border-top-color: #80ffff;
      }
    }
  }
}
</style>
