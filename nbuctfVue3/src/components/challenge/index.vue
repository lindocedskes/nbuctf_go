<template>
  <div id="challenge" class="flex flex-wrap">
    <div
      class="w-1/4 p-4"
      v-for="(challenge, index) in filteredChallenges"
      :key="index"
    >
      <challenge-card
        :data="challenge"
        :challengeType="challengeType"
        @update="update"
      ></challenge-card>
    </div>
  </div>
</template>

<script setup>
import { getGameList } from '@/api/game.js'
import { ref, watch, computed } from 'vue'
import { ElMessage } from 'element-plus'
import challengeCard from '@/components/challenge/ChallengeCard.vue'

const type = ref('')
const loading = ref(true)
const challengeData = ref([])
// 查询
const getGameData = async () => {
  const res = await getGameList()
  if (res.code === 0) {
    challengeData.value = res.data.gameList
  } else {
    ElMessage({
      type: 'error',
      message: res.msg
    })
  }
  loading.value = false // 关闭loading
}

const props = defineProps({
  //接受父组件传递的数据
  challengeType: String
})
const initPage = async () => {
  getGameData()
  //   console.log(props.challengeType)
}
initPage()

watch(
  () => type.value,
  () => {
    initPage()
  }
)
const update = () => {
  initPage()
}
//实现筛选功能
const filteredChallenges = computed(() => {
  return challengeData.value.filter(
    (challenge) =>
      challenge.queType.toLowerCase() === props.challengeType.toLowerCase() ||
      props.challengeType.toLowerCase() === 'all'
  )
})
</script>
<style scoped></style>
