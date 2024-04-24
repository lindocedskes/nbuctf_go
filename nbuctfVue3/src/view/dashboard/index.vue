<template>
  <div id="challenge-card">
    <el-card
      :body-style="{ padding: '0px' }"
      class="box-card challenge-card"
      shadow="always"
    >
      <el-row>
        <div style="float: right">
          <el-popover
            v-if="role"
            placement="top-start"
            width="150"
            trigger="hover"
          >
            <div>
              <i class="el-icon-info" style="color: red"></i
              ><span style="margin-left: 5px">删除题目,高危操作!</span>
            </div>
            <template #reference>
              <el-button
                style="float: right; padding: 3px 0"
                type="text"
                @click="close"
              >
                <i class="el-icon-close" style="color: #f56c6c"></i>
              </el-button>
            </template>
          </el-popover>
        </div>
      </el-row>
      <div class="card-text" @click="cardClick">
        <h4 style="color: #8f8f8f">{{ challenge.title }}</h4>
        <p style="color: coral">Score: {{ challenge.score }}</p>
        <div style="font-size: 18px; font-weight: bold; color: #00e18b">
          Solvers: {{ challenge.solvers }}
        </div>
        <div v-if="challenge.solved" class="solve challenge-solved">
          You Sovled: <span style="color: #00dd30">✔</span>
        </div>
        <div v-if="!challenge.solved" class="solve challenge-nosolved">
          You Not Sovled
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useStore } from 'vuex'
import { ElMessage, ElMessageBox, ElNotification } from 'element-plus'
import { SubmitFlag } from '@/api/api.js'

const props = defineProps({
  data: Object
})

const store = useStore()
const role = computed(() => store.state.user.role)

const challenge = computed(() => {
  const data = { ...props.data }
  data.des = data.des.replace(/[\n\r]/g, '<br/>')
  return data
})

const close = () => {
  emit('delete', challenge.value)
}

const cardClick = async () => {
  try {
    const { value } = await ElMessageBox.prompt(
      challenge.value.des,
      challenge.value.title,
      {
        confirmButtonText: '提交',
        cancelButtonText: '取消',
        center: true,
        inputPattern: /^E-CTF{[\w]*}$/,
        inputErrorMessage: 'Flag格式不正确',
        dangerouslyUseHTMLString: true
      }
    )

    if (store.state.user) {
      const response = await SubmitFlag({
        cid: challenge.value.cid,
        flag: value,
        user: store.state.user
      })
      switch (response.data.code) {
        case 200:
          challenge.value.solved = true
          challenge.value.solvers++
          ElNotification({
            title: '成功',
            message: '提交flag正确',
            type: 'success'
          })
          break
        case 600:
          ElNotification({
            title: '错误',
            message: 'flag错误',
            type: 'error'
          })
          break
        case 700:
          ElNotification({
            title: '警告',
            message: '请勿重复提交!',
            type: 'warning'
          })
          break
        case 400:
          await ElMessageBox.confirm('请先登录', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          })
          break
        case 500:
          ElMessage.error('token已失效,请重新登录!')
          store.state.user = {}
          break
        default:
          ElMessage.error('未知错误，请联系管理员')
          break
      }
    }
  } catch (error) {
    if (error !== 'cancel') {
      ElNotification({
        title: '消息',
        message: '取消输入',
        type: 'info'
      })
    }
  }
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
