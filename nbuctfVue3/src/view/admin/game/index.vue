<template>
  <div>
    <div class="table-box">
      <div class="l-list">
        <label
          for="helper-text"
          class="block text-sm font-medium text-gray-900 w-1/8"
          >容器镜像:</label
        >
        <input
          v-model="ImageInfo.containerImage"
          type="email"
          id="helper-text"
          aria-describedby="helper-text-explanation"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-1/4 p-2.5"
          placeholder="请输入测试容器镜像地址"
        />
        <label
          for="helper-text"
          class="block mb-2 text-sm font-medium text-gray-900 w-1/8 ml-4"
          >服务端口:</label
        >
        <input
          v-model="ImageInfo.innerPort"
          type="email"
          id="helper-text"
          aria-describedby="helper-text-explanation"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-1/6 p-2.5"
          placeholder="测试容器内部运行端口（一般都为80"
        />

        <div class="w-1/3 ml-4">
          <div v-if="testcontainInfo.containerAddr">
            <label
              class="text-gray-900 bg-gradient-to-r from-lime-200 via-lime-400 to-lime-500 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-lime-300 dark:focus:ring-lime-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
            >
              {{
                testcontainInfo.containerAddr + ':' + testcontainInfo.outPort
              }}
            </label>
            <button
              @click="openWebPage"
              type="button"
              class="ml-2 text-gray-900 bg-gradient-to-r from-teal-200 to-lime-200 hover:bg-gradient-to-l hover:from-teal-200 hover:to-lime-200 focus:ring-4 focus:outline-none focus:ring-lime-200 dark:focus:ring-teal-700 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
            >
              Open
              <el-icon class="#dc2626"><Share /></el-icon>
            </button>
          </div>

          <label
            v-else
            class="text-white bg-gradient-to-r from-red-400 via-red-500 to-red-600 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-red-300 dark:focus:ring-red-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"
          >
            测试容器未开启
          </label>
        </div>
      </div>
      <div class="btn-list">
        <a
          class="group inline-block rounded-full bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 p-[2px] hover:text-white focus:outline-none focus:ring active:text-opacity-75"
          href="#"
        >
          <span
            @click="openTestContainer"
            class="block rounded-full bg-white px-4 py-1 text-sm font-medium group-hover:bg-transparent"
          >
            开启测试容器
          </span>
        </a>
        <a
          class="group inline-block rounded-full bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 p-[2px] hover:text-white focus:outline-none focus:ring active:text-opacity-75"
          href="#"
        >
          <span
            @click="closeTestContainer"
            class="block rounded-full bg-white px-4 py-1 text-sm font-medium group-hover:bg-transparent"
          >
            删除测试容器
          </span>
        </a>
        <a
          class="group inline-block rounded-full bg-gradient-to-r from-pink-500 via-red-500 to-yellow-500 p-[2px] hover:text-white focus:outline-none focus:ring active:text-opacity-75"
          href="#"
        >
          <span
            @click="checkTestContainer"
            class="block rounded-full bg-white px-4 py-1 text-sm font-medium group-hover:bg-transparent"
          >
            查询测试容器状态
          </span>
        </a>

        <button
          type="button"
          @click="closeAllContainer"
          class="text-white bg-gradient-to-br from-pink-500 to-orange-400 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-pink-200 dark:focus:ring-pink-800 font-medium rounded-lg text-sm px-4 py-1.5 text-center"
        >
          关闭所有比赛容器
        </button>
      </div>
      <div class="btn-list">
        <a
          class="group relative inline-block focus:outline-none focus:ring"
          href="#"
        >
          <span
            class="absolute inset-0 translate-x-1.5 translate-y-1.5 bg-yellow-300 transition-transform group-hover:translate-x-0 group-hover:translate-y-0"
          ></span>

          <span
            @click="addQuestion"
            class="relative inline-block border-2 border-current px-4 py-1 text-sm font-bold uppercase tracking-widest text-black group-active:text-opacity-75"
          >
            新增赛题
          </span>
        </a>
      </div>

      <el-table :data="tableData" row-key="id">
        <el-table-column align="left" label="题目类型" min-width="100">
          <template #default="scope">
            <el-tag :type="tagType(scope.row.queType)" disable-transitions
              >{{ scope.row.queType }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="id" min-width="50" prop="id" />
        <el-table-column
          align="left"
          label="题目名称"
          min-width="150"
          prop="queName"
        />
        <el-table-column
          align="left"
          label="题目描述"
          min-width="150"
          prop="queDescribe"
        />
        <el-table-column
          align="left"
          label="题目分值"
          min-width="150"
          prop="queMark"
        />
        <el-table-column
          align="left"
          label="解题flag"
          min-width="180"
          prop="queFlag"
        />
        <el-table-column align="left" label="启用" min-width="150">
          <template #default="scope">
            <el-switch
              v-model="scope.row.ifHidden"
              inline-prompt
              :active-value="false"
              :inactive-value="true"
              @change="
                () => {
                  switchEnable(scope.row)
                }
              "
            />
          </template>
        </el-table-column>
        <el-table-column label="附件" min-width="400">
          <template #default="scope">
            <div style="display: flex; flex-wrap: wrap">
              <div
                v-for="(file, index) in scope.row.files"
                :key="index"
                class="mt-2 flex items-center"
              >
                <CustomPic
                  :pic-src="file.url"
                  picType="file"
                  class="w-14 h-12"
                />
                <div class="text-center ml-2">{{ file.name }}</div>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="靶机地址"
          min-width="180"
          prop="imageUrl"
        />
        <el-table-column
          align="left"
          label="靶机内部运行端口"
          min-width="180"
          prop="innerPort"
        />

        <el-table-column label="操作" min-width="350" fixed="right">
          <template #default="scope">
            <el-popover
              v-model:visible="scope.row.visible"
              placement="top"
              width="160"
            >
              <p>确定要删除此题目吗</p>
              <div style="text-align: right; margin-top: 8px">
                <el-button
                  type="primary"
                  link
                  @click="scope.row.visible = false"
                  >取消</el-button
                >
                <el-button type="primary" @click="deleteQuestionFunc(scope.row)"
                  >确定</el-button
                >
              </div>
              <template #reference>
                <el-button
                  type="primary"
                  link
                  icon="delete"
                  class="text-emerald-400"
                  >删除</el-button
                >
              </template>
            </el-popover>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="openEdit(scope.row)"
              class="text-emerald-400"
              >编辑</el-button
            >
            <el-button
              type="primary"
              link
              icon="Link"
              @click="openFiles(scope.row)"
              class="text-emerald-400"
              >添加附件</el-button
            >
            <el-button
              type="primary"
              link
              icon="Link"
              @click="deleteFiles(scope.row)"
              class="text-emerald-400"
              >清空附件</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog
      v-model="addQuestionDialog"
      title="赛题"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height: 60vh; overflow: auto; padding: 0 12px">
        <el-form
          ref="questionForm"
          :rules="rules"
          :model="questionInfo"
          label-width="100px"
        >
          <el-form-item label="题目类型" prop="queType">
            <el-input v-model="questionInfo.queType" />
          </el-form-item>
          <el-form-item label="题目名称" prop="queName">
            <el-input v-model="questionInfo.queName" />
          </el-form-item>
          <el-form-item label="题目描述" prop="queDescribe">
            <el-input v-model="questionInfo.queDescribe" />
          </el-form-item>
          <el-form-item label="题目分值" prop="queMark">
            <el-input v-model.number="questionInfo.queMark" />
          </el-form-item>
          <el-form-item label="解题flag" prop="queFlag">
            <el-input v-model="questionInfo.queFlag" />
          </el-form-item>
          <el-form-item label="靶机地址" prop="queFlag">
            <el-input v-model="questionInfo.imageUrl" />
          </el-form-item>
          <el-form-item label="靶机内部端口" prop="queFlag">
            <el-input v-model="questionInfo.innerPort" />
          </el-form-item>
          <el-form-item label="启用" prop="ifHidden">
            <el-switch
              v-model="questionInfo.ifHidden"
              inline-prompt
              :active-value="false"
              :inactive-value="true"
            />
          </el-form-item>
        </el-form>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAddQuestionDialog">取 消</el-button>
          <el-button type="primary" @click="enterAddQuestionDialog"
            >确 定</el-button
          >
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="addFileDialog"
      title="添加附件"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height: 60vh; overflow: auto; padding: 0 12px">
        <el-form ref="fileForm" label-width="80px" :model="fileInfo">
          <el-form-item label="附件" label-width="80px">
            <div style="display: inline-block" @click="openHeaderChange">
              <!-- 空数组被视为真值 -->
              <div v-if="fileInfo.files.length > 0">
                <CustomPic :pic-src="fileInfo.files[0].url" picType="file" />
                <div class="text-center ml-2">
                  {{ fileInfo.files[0].name }}
                </div>
              </div>
              <div v-else class="header-img-box" w-16 h-16>从文件库选择</div>
            </div>
          </el-form-item>
        </el-form>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAddFileDialog">取 消</el-button>
          <el-button type="primary" @click="enterAddFileDialog"
            >确 定</el-button
          >
        </div>
      </template>
    </el-dialog>
    <ChooseImg
      ref="chooseImg"
      :target="fileInfo.files[0]"
      :target-key="`url`"
      @updateFiles="updateFiles"
    />
  </div>
</template>

<script setup>
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import {
  getGameList,
  deleteGame,
  addGame,
  editGame,
  addFile,
  deleteFile
} from '@/api/gameadmin'
import {
  openContainerApi,
  checkContainerApi,
  closeContainerApi,
  closeAllContainerApi
} from '@/api/k8s'

import { nextTick, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'GameAdmin'
})
const tableData = ref([])

// 查询
const getTableData = async () => {
  const table = await getGameList()
  if (table.code === 0) {
    tableData.value = table.data
  }
}
//watch 函数接受两个参数,被监视的数据和回调函数
watch(
  () => tableData.value,
  () => {
    // setAuthorityIds()
  }
)

const initPage = async () => {
  getTableData()
}

initPage()

const tagType = (queType) => {
  switch (
    queType.toLowerCase() //不区分大小写
  ) {
    case 'web':
      return 'success'
    case 'misc':
      return 'info'
    case 'crypto':
      return 'warning'
    case 'pwn':
      return 'danger'
    default:
      return 'primary'
  }
}

const deleteQuestionFunc = async (row) => {
  const res = await deleteGame({ id: row.id })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    row.visible = false
    await getTableData()
  }
}

// 弹窗相关
const questionInfo = ref({
  queType: '',
  queName: '',
  queDescribe: '',
  queMark: '',
  queFlag: '',
  ifHidden: true,
  imageUrl: '',
  innerPort: '' //默认80端口
})

const rules = ref({
  queType: [{ required: true, message: '请输入题目类型', trigger: 'blur' }],
  queName: [{ required: true, message: '请输入题目名称', trigger: 'blur' }],
  queMark: [
    { required: true, message: '请输入题目分值', trigger: 'blur' },
    {
      pattern: /^\d+$/,
      message: '请输入合法整数',
      trigger: 'blur'
    }
  ]
})
const questionForm = ref(null)
const enterAddQuestionDialog = async () => {
  questionForm.value.validate(async (valid) => {
    if (valid) {
      const req = {
        ...questionInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await addGame(req) //新增赛题
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddQuestionDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await editGame(req) //编辑赛题
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddQuestionDialog()
        }
      }
    }
  })
}
const openEdit = (row) => {
  dialogFlag.value = 'edit'
  questionInfo.value = JSON.parse(JSON.stringify(row)) // 深拷贝row对象 转 string  再转 json，传入iscope.row 是当前行的数据对象
  addQuestionDialog.value = true
}
const addQuestionDialog = ref(false)
const closeAddQuestionDialog = () => {
  questionForm.value.resetFields() //重置表单
  addQuestionDialog.value = false
}
const dialogFlag = ref('add')

const addQuestion = () => {
  dialogFlag.value = 'add'
  addQuestionDialog.value = true
}

// 添加附件
const fileForm = ref(null)
const fileInfo = ref({
  files: [],
  id: ''
})
// 添加附件
const enterAddFileDialog = async () => {
  const req = {
    questionId: fileInfo.value.id,
    fileId: fileInfo.value.files[0].id
  }
  const res = await addFile(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '添加成功' })
    await getTableData()
    closeAddFileDialog()
  }
}

const openFiles = (row) => {
  fileInfo.value = JSON.parse(JSON.stringify(row)) // 深拷贝row对象 转 string  再转 json，传入iscope.row 是当前行的数据对象
  addFileDialog.value = true
}
const deleteFiles = async (row) => {
  const req = {
    id: row.id
  }
  const res = await deleteFile(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '清空成功' })
    await getTableData()
  } else {
    ElMessage({ type: 'error', message: '清空失败' })
  }
}

const addFileDialog = ref(false)
const closeAddFileDialog = () => {
  fileForm.value.resetFields() //重置表单
  addFileDialog.value = false
}

const switchEnable = async (row) => {
  questionInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  console.log(questionInfo.value)
  console.log(row)
  const req = {
    ...questionInfo.value
  }
  const res = await editGame(req)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: `${req.ifHidden === true ? '禁用' : '启用'}成功`
    })
    await getTableData()
  }
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}
const updateFiles = (newFiles) => {
  fileInfo.value.files[0] = newFiles
}

const ImageInfo = ref({ containerImage: '', innerPort: 80 })
const testcontainInfo = ref({ containerAddr: '', outPort: '' })

const openTestContainer = async () => {
  //判断是containerImage否为空
  if (!ImageInfo.value.containerImage) {
    ElMessage.error('请输入测试容器镜像地址')
    return
  }
  if (!ImageInfo.value.innerPort) {
    ElMessage.error('请输入测试容器内部运行端口')
    return
  }
  const req = {
    imageAddr: ImageInfo.value.containerImage,
    innerPort: ImageInfo.value.innerPort
  }
  console.log(req)
  const res = await openContainerApi(req)
  if (res.code === 0) {
    testcontainInfo.value = res.data
    ElMessage.success(res.msg)
  } else {
    // ElMessage.error(res.msg)
  }
}
const checkTestContainer = async () => {
  const req = { imageAddr: ImageInfo.value.containerImage }
  const res = await checkContainerApi(req)
  if (res.code === 0) {
    testcontainInfo.value = res.data
    ElMessage.success(res.msg)
  } else {
    // ElMessage.error(res.msg)
  }
}
const closeTestContainer = async () => {
  const req = { imageAddr: ImageInfo.value.containerImage }
  const res = await closeContainerApi(req)
  if (res.code === 0) {
    testcontainInfo.value = res.data // 清空测试容器信息
    ElMessage.success(res.msg)
  } else {
    // ElMessage.error(res.msg)
  }
}
// 关闭所有容器
const closeAllContainer = async () => {
  ElMessageBox.confirm('此操作将删除比赛所有容器, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      const res = await closeAllContainerApi()
      if (res.code === 0) {
        testcontainInfo.value = res.data // 清空测试容器信息
        ElMessage.success(res.msg)
      } else {
        // ElMessage.error(res.msg)
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除'
      })
    })
}

const openWebPage = () => {
  window.open(
    'http://' +
      testcontainInfo.value.containerAddr +
      ':' +
      testcontainInfo.value.outPort
  )
}
</script>

<style lang="scss">
.header-img-box {
  @apply w-32 h-32 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
}
.table-box {
  @apply p-6 bg-white rounded;
}
.btn-list {
  @apply mb-3 flex gap-3 items-center;
}
.l-list {
  @apply mb-3 flex gap-1 items-center;
}
.pagination {
  @apply flex justify-end;
  .el-pagination__editor {
    .el-input__inner {
      @apply h-8;
    }
  }

  .is-active {
    @apply rounded text-white;
    background: var(--el-color-primary);
    color: #ffffff !important;
  }
}
.el-table .cell {
  background-color: inherit;
}
</style>
