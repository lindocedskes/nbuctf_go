<template>
  <div id="challenge-card">
    <el-card
      :body-style="{ padding: '0px' }"
      class="challenge-card"
      shadow="always"
    >
      <div class="card-text" @click="cardClick">
        <div
          class="bg-gray-100 text-gray-800 py-1 px-2 rounded-full font-semibold mb-2 shadow-sm uppercase"
        >
          {{ challenge.queType }}
        </div>
        <h4 style="color: #8f8f8f">{{ challenge.queName }}</h4>
        <p style="color: coral">Score: {{ challenge.queMark }}</p>
        <div style="font-size: 18px; font-weight: bold; color: #00e18b">
          Solvers: {{ challenge.queSolvers }}
        </div>
        <!-- challenge.solvers -->
        <div v-if="challenge.ifSolved" class="solve challenge-solved">
          You Sovled: <span style="color: #00dd30">✔</span>
        </div>
        <div v-if="!challenge.ifSolved" class="solve challenge-nosolved">
          You Not Sovled
        </div>
      </div>
    </el-card>
  </div>

  <el-dialog
    v-model="submitQuestionDialog"
    :title="challenge.queName"
    :show-close="false"
  >
    <div style="height: 60vh; overflow: auto; padding: 0 12px">
      <el-form
        ref="challengeForm"
        :rules="rules"
        :model="challengeInfo"
        label-width="100px"
      >
        <el-form-item label="题目类型:" prop="queType">
          <label class="block text-gray-700 font-bold mb-2">
            {{ challengeInfo.queType }}
          </label>
        </el-form-item>

        <el-form-item label="题目描述: （markdown）" prop="queDescribe">
          <div
            v-html="md.render(challengeInfo.queDescribe)"
            class="border border-gray-300 p-4 my-2 bg-white rounded shadow"
          />
        </el-form-item>
        <el-form-item label="题目分值:" prop="queMark">
          <label class="block text-gray-700 font-bold mb-2">
            {{ challengeInfo.queMark }}
          </label>
        </el-form-item>
        <el-form-item
          label="附件:"
          prop="Files"
          v-if="challengeInfo.Files && challengeInfo.Files.length"
        >
          <div
            v-for="(file, index) in challengeInfo.Files"
            :key="index"
            class="flex justify-between items-center"
          >
            <label class="block text-gray-700 font-bold">
              {{ file.name }}
            </label>
            <el-button
              @click="() => downloadFile(file)"
              class="bg-green-500 hover:bg-green-700 text-white font-bold py-1 px-4 mx-2 rounded"
              >下载</el-button
            >
          </div>
        </el-form-item>
        <el-form-item
          label="靶机:"
          prop="queMark"
          v-if="challengeInfo.imageUrl"
        >
          <label class="block text-gray-700 font-bold mb-2">
            {{ challengeInfo.imageUrl }}
          </label>
          <el-button
            @click="() => downloadFile(file)"
            class="bg-green-500 hover:bg-green-700 text-white font-bold py-1 px-4 mx-2 rounded"
            >开启</el-button
          >
        </el-form-item>
        <el-form-item label="解题flag:" prop="queFlag">
          <el-input v-model="challengeInfo.queFlag" />
          <el-button type="primary" @click="enterSubmitFlag" class="py-1 my-2"
            >提 交</el-button
          >
        </el-form-item>
      </el-form>
    </div>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="closeSubmitQuestionDialog">关闭</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ElMessage, ElMessageBox, ElNotification } from 'element-plus'
import { submitFlag } from '@/api/game.js'
import { computed, ref } from 'vue'
import MarkdownIt from 'markdown-it'
const md = new MarkdownIt()

const props = defineProps({
  data: Object, //接受父组件传递的数据
  challengeType: String
})
// console.log(props.challengeType)
const challenge = computed(() => {
  const data = { ...props.data }
  return data
})
// 挑战题目信息
const challengeForm = ref(null)
const rules = ref({
  queFlag: [{ required: true, message: '请输入flag', trigger: 'blur' }]
})
// 弹窗相关
const challengeInfo = ref({
  queType: challenge.value.queType,
  queName: challenge.value.queName,
  queDescribe: challenge.value.queDescribe,
  queMark: challenge.value.queMark,
  Files: challenge.value.files,
  imageUrl: challenge.value.imageUrl,
  queFlag: ''
})
const cardClick = () => {
  submitQuestionDialog.value = true
}
const closeSubmitQuestionDialog = () => {
  //   challengeForm.value.resetFields() //重置表单
  submitQuestionDialog.value = false
}
const submitQuestionDialog = ref(false)
const emit = defineEmits(['update']) //定义emit事件，用于触发父组件事件（update）
const enterSubmitFlag = async () => {
  challengeForm.value.validate(async (valid) => {
    if (valid) {
      console.log(challenge)
      const req = {
        questionId: challenge.value.id,
        submitFlag: challengeInfo.value.queFlag
      }
      const res = await submitFlag(req) //提交flag
      if (res.code === 0) {
        ElMessage({ type: 'success', message: '提交flag成功' })
        closeSubmitQuestionDialog()
        //刷新
        emit('update') //触发父组件的update事件
      } else {
        ElMessage({ type: 'error', message: res.msg })
      }
    }
  })
}

const downloadFile = (file) => {
  const url = `/api/${file.url}`
  const a = document.createElement('a')
  a.href = url
  a.download = file.name
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
}
</script>

<style scoped>
/* CSS styles remain unchanged */
.challenge-card {
  min-height: 200px;
  max-width: 180px;
  margin-top: 10px;
  margin-bottom: 10px;
}

.challenge-card:hover {
  cursor: pointer;
}

.card-text {
  margin-top: 10%;
}

.challenge-nosolved,
.challenge-solved,
.solve {
  margin-top: 15px;
  font-size: 18px;
  font-weight: bold;
}

.challenge-solved {
  color: #66b1ff;
}

.challenge-nosolved {
  color: #ff2222;
}
</style>
