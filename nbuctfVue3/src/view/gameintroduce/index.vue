<template>
  <div id="index">
    <div class="block">
      <div class="container backgroundForm px-10 py-5">
        <div class="flex justify-center items-center">
          <dv-decoration5
            :dur="3"
            style="width: 300px; height: 40px"
            :color="['#80ffff', '#80ffff']"
          />
        </div>
        <h1 style="text-align: center" class="text-2xl text-cyan-300">
          比赛介绍
          <el-popover
            v-if="userStore.userInfo.authorityId === 888"
            placement="top-start"
            width="50"
            trigger="hover"
          >
            <div>
              <span style="margin-left: 5px">可创建比赛介绍</span>
            </div>
            <template #reference>
              <el-button
                size="small"
                icon="plus"
                circle
                autofocus
                @click="showDialog = true"
                class="text-gray-900 bg-gradient-to-r from-pink-400 to-cyan-200 hover:bg-gradient-to-l hover:from-pink-400 hover:to-cyan-200 focus:ring-4 focus:outline-none focus:ring-cyan-200 font-medium rounded-lg text-sm px-5"
              ></el-button>
            </template>
          </el-popover>
        </h1>

        <!-- </el-divider> -->
        <el-timeline v-if="reload">
          <el-timeline-item
            v-for="(ann, idx) in sortedAnncs"
            :key="idx"
            placement="top"
            :color="nodeColor"
            class="text-left"
          >
            <el-card
              class="text-gray-900 bg-transparent focus:ring-4 focus:outline-none focus:ring-lime-200 font-medium rounded-lg text-sm px-5 py-0 me-0 mb-2 border-0"
            >
              <!-- <template #header> </template> -->
              <div class="clearfix text-left">
                <!-- <span>{{ ann.title }}</span> -->
              </div>

              <!-- <h4
                class="text-left text-cyan-300"
                v-html="md.render(ann.content)"
              ></h4> -->

              <MdPreview
                :editorId="id"
                :modelValue="ann.content"
                class="bg-transparent"
              />
              <!-- <MdCatalog :editorId="id" :scrollElement="scrollElement" /> -->
              <p
                style="float: right; font-size: 12px; color: #8f8f8f"
                class="text-cyan-500"
              >
                管理员 发布于: {{ formatDate(ann.CreatedAt) }}
              </p>

              <el-popconfirm
                v-if="userStore.userInfo.authorityId === 888"
                @confirm="DelAnnc(ann.id)"
                confirm-button-text="确定"
                cancel-button-text="取消"
                :icon="InfoFilled"
                icon-color="red"
                title="确定删除该介绍吗？"
                confirm-button-type="danger"
              >
                <template #reference>
                  <el-button style="float: right; padding: 3px 0" link>
                    <el-icon style="color: #f56c6c"><Close /></el-icon>
                  </el-button>
                </template>
              </el-popconfirm>
              <el-button
                v-if="userStore.userInfo.authorityId === 888"
                @click="openDialog(ann)"
                style="float: right; padding: 3px 3px"
                link
              >
                <el-icon style="color: pink"><EditPen /></el-icon>
              </el-button>
            </el-card>
          </el-timeline-item>
        </el-timeline>
      </div>
    </div>
    <el-dialog
      center
      title="创建介绍"
      v-model="showDialog"
      width="60%"
      :show-close="false"
    >
      <el-form label-position="left" label-width="50px">
        <div>
          建议使用字体颜色：&lt;font color= #80ffff size=4&gt;text&lt;/font&gt;
        </div>
        <el-form-item label="标题:">
          <el-input
            v-model="title"
            placeholder="必须输入'比赛介绍'作为标题,否则为公告"
          ></el-input>
        </el-form-item>
        <el-form-item label="内容:">
          <MdEditor v-model="body" class="h-80" />
          <!-- <el-input type="textarea" v-model="body"></el-input> -->
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="handleClose">取 消</el-button>
          <el-button type="primary" @click="Submit">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog
      center
      title="更新介绍"
      v-model="showDialogUpdate"
      width="60%"
      :show-close="false"
    >
      <el-form label-position="left" label-width="50px">
        <el-form-item label="提示:">
          建议使用字体颜色：&lt;font color= #80ffff size=4&gt;text&lt;/font&gt;
        </el-form-item>
        <el-form-item label="内容:">
          <MdEditor v-model="body" class="h-80" />
          <!-- <el-input type="textarea" v-model="body"></el-input> -->
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="handleClose">取 消</el-button>
          <el-button type="primary" @click="UpdAnnc(annc)">更新</el-button>
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
  AddAnnouncement,
  UpdateAnnouncement
} from '@/api/announcement.js'
import { formatDate } from '@/utils/format'
import { useUserStore } from '@/pinia'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'

import { MdPreview } from 'md-editor-v3'
import 'md-editor-v3/lib/preview.css'

// import MarkdownIt from 'markdown-it'
// const md = new MarkdownIt()
const id = 'preview-only'

defineOptions({
  name: 'announcement'
})

const userStore = useUserStore()

const showDialog = ref(false)
const showDialogUpdate = ref(false)

const nodeColor = ref('#80ffff')
const anncs = ref([])
const reload = ref(true)
const title = ref('')
const body = ref('')

const InitAnnc = async () => {
  try {
    let res = await GetAnnouncements()
    anncs.value = res.data.filter((item) => item.title === '比赛介绍') //只获取比赛介绍的公告

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
  showDialogUpdate.value = false
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

const UpdAnnc = async () => {
  console.log()
  const res = await UpdateAnnouncement({
    id: annc.value.id,
    title: annc.value.title,
    content: body.value
  })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更新成功'
    })
    showDialogUpdate.value = false
    InitAnnc()
  } else {
    ElMessage({
      type: 'error',
      message: '更新失败'
    })
  }
}

const annc = ref({})
const openDialog = (ann) => {
  body.value = ann.content
  annc.value = {
    id: ann.id,
    title: ann.title
  }
  showDialogUpdate.value = true
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
.text_color {
  color: #80ffff !important;
}
</style>
<style>
.el-divider__text {
  background-color: transparent !important;
}
.backgroundForm {
  height: 80vh;
  background-image: url('/src/assets/form_bg.png');
  background-size: 115% 110%;
  background-position: center;
}
</style>
