<template>
  <div id="index">
    <div class="block">
      <div class="container">
        <h1 style="text-align: center" class="text-2xl text-cyan-300">
          公告板
        </h1>
        <el-divider content-position="right">
          <el-popover
            v-if="userStore.userInfo.authorityId === 888"
            placement="top-start"
            width="50"
            trigger="hover"
          >
            <div>
              <span style="margin-left: 5px">可创建公告</span>
            </div>
            <template #reference>
              <el-button
                size="small"
                icon="plus"
                circle
                autofocus
                @click="showDialog = true"
                class="text-gray-900 bg-gradient-to-r from-teal-200 to-lime-200 hover:bg-gradient-to-l hover:from-teal-200 hover:to-lime-200 focus:ring-4 focus:outline-none focus:ring-lime-200 dark:focus:ring-teal-700 font-medium rounded-lg text-sm px-5"
              ></el-button>
            </template>
          </el-popover>
        </el-divider>
        <el-timeline v-if="reload">
          <el-timeline-item
            v-for="(ann, idx) in sortedAnncs"
            :key="idx"
            :timestamp="formatDate(ann.CreatedAt).trim().split(/\s+/)[0]"
            placement="top"
            :color="nodeColor"
            class="text-left"
          >
            <el-card
              class="text-gray-900 bg-gradient-to-r from-pink-400 to-cyan-300 hover:bg-gradient-to-l hover:from-pink-400 hover:to-cyan-300 focus:ring-4 focus:outline-none focus:ring-lime-200 dark:focus:ring-teal-700 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2"
            >
              <template #header>
                <div class="clearfix text-left">
                  <span>{{ ann.title }}</span>
                  <el-popconfirm
                    v-if="userStore.userInfo.authorityId === 888"
                    @confirm="DelAnnc(ann.id)"
                    confirm-button-text="确定"
                    cancel-button-text="取消"
                    :icon="InfoFilled"
                    icon-color="red"
                    title="确定删除该公告吗？"
                    confirm-button-type="danger"
                  >
                    <template #reference>
                      <el-button style="float: right; padding: 3px 0" link>
                        <el-icon style="color: #99154b"><Close /></el-icon>
                      </el-button>
                    </template>
                  </el-popconfirm>
                </div>
              </template>

              <h4 class="text-left">{{ ann.content }}</h4>
              <p style="float: right; font-size: 12px; color: #8f8f8f">
                管理员 发布于: {{ formatDate(ann.CreatedAt) }}
              </p>
            </el-card>
          </el-timeline-item>
        </el-timeline>
      </div>
    </div>
    <el-dialog
      center
      title="创建公告"
      v-model="showDialog"
      width="30%"
      :show-close="false"
    >
      <el-form label-position="left" label-width="50px">
        <el-form-item label="标题:">
          <el-input v-model="title" placeholder="请输入标题"></el-input>
        </el-form-item>
        <el-form-item label="内容:">
          <el-input type="textarea" v-model="body"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="handleClose">取 消</el-button>
          <el-button type="primary" @click="Submit">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { InfoFilled } from '@element-plus/icons-vue'

import { ref, onMounted, nextTick, computed } from 'vue'
import {
  GetAnnouncements,
  DeleteAnnouncement,
  AddAnnouncement
} from '@/api/announcement.js'
import { formatDate } from '@/utils/format'
import { useUserStore } from '@/pinia'

defineOptions({
  name: 'announcement'
})

const userStore = useUserStore()

const showDialog = ref(false)
const nodeColor = ref('#80ffff')
const anncs = ref([])
const reload = ref(true)
const title = ref('')
const body = ref('')

const InitAnnc = async () => {
  try {
    let res = await GetAnnouncements()

    anncs.value = res.data.filter((item) => item.title != '比赛介绍') //只获取非比赛介绍标题的公告
    console.log(anncs.value)
    //ref创建的reload变化，强制重新渲染
    reload.value = false
    nextTick(() => {
      reload.value = true
    })
  } catch (e) {
    ElMessage({
      type: 'error',
      message: '获取数据失败,请检查网络!'
    })
  }
}

onMounted(() => {
  InitAnnc()
})

const sortedAnncs = computed(() => {
  return [...anncs.value].sort(
    (a, b) => new Date(b.CreatedAt) - new Date(a.CreatedAt)
  )
})
const handleClose = () => {
  showDialog.value = false
}
const Submit = async () => {
  try {
    await AddAnnouncement({
      title: title.value,
      content: body.value
    })
    ElMessage({
      type: 'success',
      message: '添加成功'
    })
    showDialog.value = false
    InitAnnc() //重新获取公告
  } catch (e) {
    ElMessage({
      type: 'error',
      message: '添加失败'
    })
  }
}
const DelAnnc = async (id) => {
  const res = await DeleteAnnouncement({
    id: id
  })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    InitAnnc()
  } else {
    ElMessage({
      type: 'error',
      message: '删除失败'
    })
  }
}
</script>

<style scoped>
.container {
  margin: 0 auto;
  max-width: 45%;
}
.text-gradient {
  background: linear-gradient(to right, #38b2ac, #38b2ac);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
.el-divider__text {
  background-color: transparent !important;
}
</style>
<style>
.el-divider__text {
  background-color: transparent !important;
}
</style>
