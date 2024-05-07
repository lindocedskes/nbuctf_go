<template>
  <div
    id="challenge-card-bg"
    class="hover:animate-background rounded-xl bg-gradient-to-r from-green-300 via-blue-500 to-purple-600 p-0.5 shadow-xl transition hover:bg-[length:400%_400%] hover:shadow-sm hover:[animation-duration:_4s]"
  >
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
        <p style="color: coral">分值: {{ challenge.queMark }}</p>
        <div style="font-size: 18px; font-weight: bold; color: #00e18b">
          已解决人数: {{ challenge.queSolvers }}
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
    <div style="height: 64vh; overflow: auto; padding: 0 12px">
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

        <el-form-item
          label="题目描述: （markdown）"
          prop="queDescribe"
          class="text-left"
        >
          <!-- <div class="border border-gray-300 p-4 my-2 rounded shadow"> -->
          <MdPreview :modelValue="challengeInfo.queDescribe" />
          <!-- </div> -->
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
          <!-- <label class="block text-gray-700 font-bold mb-2">
            {{ challengeInfo.imageUrl }}
          </label> -->
          <div>
            <div v-if="containInfo.containerAddr" class="btn-list">
              <label
                class="text-gray-900 bg-gradient-to-r from-lime-200 via-lime-400 to-lime-500 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-lime-300 dark:focus:ring-lime-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
              >
                {{ containInfo.containerAddr + ':' + containInfo.outPort }}
              </label>
              <button
                @click="openWebPage"
                type="button"
                class="text-gray-900 bg-gradient-to-r from-teal-200 to-lime-200 hover:bg-gradient-to-l hover:from-teal-200 hover:to-lime-200 focus:ring-4 focus:outline-none focus:ring-lime-200 dark:focus:ring-teal-700 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
              >
                Open
                <el-icon class="#dc2626"><Share /></el-icon>
              </button>
            </div>

            <label
              v-else
              class="text-white bg-gradient-to-r from-red-400 via-red-500 to-red-600 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-red-300 dark:focus:ring-red-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"
            >
              靶机未开启
            </label>
          </div>
          <div class="btn-list mt-1">
            <a
              class="group inline-block rounded-full bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 p-[2px] hover:text-white focus:outline-none focus:ring active:text-opacity-75"
              href="#"
            >
              <span
                @click="openContainer"
                class="block rounded-full bg-white px-4 py-1 text-sm font-medium group-hover:bg-transparent"
              >
                开启靶机
              </span>
            </a>
            <a
              class="group inline-block rounded-full bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 p-[2px] hover:text-white focus:outline-none focus:ring active:text-opacity-75"
              href="#"
            >
              <span
                @click="closeContainer"
                class="block rounded-full bg-white px-4 py-1 text-sm font-medium group-hover:bg-transparent"
              >
                删除靶机
              </span>
            </a>
            <a
              class="group inline-block rounded-full bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 p-[2px] hover:text-white focus:outline-none focus:ring active:text-opacity-75"
              href="#"
            >
              <span
                @click="checkContainer"
                class="block rounded-full bg-white px-4 py-1 text-sm font-medium group-hover:bg-transparent"
              >
                查询靶机状态
              </span>
            </a>
          </div>
        </el-form-item>

        <el-form-item label="解题flag:" prop="queFlag">
          <el-input v-model="challengeInfo.queFlag" />
          <el-button
            @click="enterSubmitFlag"
            class="mt-2 text-white bg-gradient-to-br from-pink-500 to-orange-400 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-pink-200 dark:focus:ring-pink-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"
            >提 交</el-button
          >
        </el-form-item>
      </el-form>
    </div>

    <template #footer>
      <div class="dialog-footer">
        <button
          @click="closeSubmitQuestionDialog"
          class="text-white bg-gradient-to-r from-cyan-500 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-cyan-300 dark:focus:ring-cyan-800 font-medium rounded-lg text-sm px-5 py-1.5 text-center me-2 mb-2"
        >
          关闭
        </button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ElMessage } from 'element-plus'
import { submitFlag } from '@/api/game.js'
import { computed, ref } from 'vue'
import {
  openContainerApi,
  checkContainerApi,
  closeContainerApi
} from '@/api/k8s'

import { MdPreview } from 'md-editor-v3'
import 'md-editor-v3/lib/preview.css'

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
  if (challenge.value.imageUrl) {
    //如果有测试容器镜像地址，自动检查容器状态
    checkContainer() //每次打开卡片自动检查容器状态
  }
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
  const url = import.meta.env.VITE_FILE_API + '/' + file.url
  // const url = `/api/${file.url}`
  const a = document.createElement('a')
  a.href = url
  a.download = file.name
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
}

const ImageInfo = ref({
  containerImage: challenge.value.imageUrl,
  innerPort: challenge.value.innerPort
})
const containInfo = ref({ containerAddr: '', outPort: '' })

const openContainer = async () => {
  console.log(ImageInfo.value)
  //判断是containerImage否为空
  if (!ImageInfo.value.containerImage) {
    ElMessage.error('未注册靶机镜像地址，请联系管理员')
    return
  }
  if (!ImageInfo.value.innerPort) {
    ElMessage.error('未注册靶机内部运行端口，请联系管理员')
    return
  }
  const req = {
    imageAddr: ImageInfo.value.containerImage,
    innerPort: parseInt(ImageInfo.value.innerPort)
  }
  console.log(req)
  const res = await openContainerApi(req)
  if (res.code === 0) {
    containInfo.value = res.data
    ElMessage.success(res.msg)
  } else {
    // ElMessage.error(res.msg)
  }
}
const checkContainer = async () => {
  const req = { imageAddr: ImageInfo.value.containerImage }
  const res = await checkContainerApi(req)
  if (res.code === 0) {
    containInfo.value = res.data
    ElMessage.success(res.msg)
  } else {
    // ElMessage.error(res.msg)
  }
}
const closeContainer = async () => {
  const req = { imageAddr: ImageInfo.value.containerImage }
  const res = await closeContainerApi(req)
  if (res.code === 0) {
    containInfo.value = res.data // 清空测试容器信息
    ElMessage.success(res.msg)
  } else {
    // ElMessage.error(res.msg)
  }
}

const openWebPage = () => {
  window.open(
    'http://' +
      containInfo.value.containerAddr +
      ':' +
      containInfo.value.outPort
  )
}
</script>

<style scoped>
/* CSS styles remain unchanged */
.challenge-card {
  min-height: 200px;
  max-width: 180px;
  margin-top: 8px;
  margin-bottom: 2px;
  margin-left: 2px;
  margin-right: 4px;
}
#challenge-card-bg {
  min-height: 200px;
  max-width: 180px;
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
.btn-list {
  @apply flex gap-3 items-center;
}
</style>
