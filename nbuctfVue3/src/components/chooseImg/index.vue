<template>
  <el-drawer v-model="drawer" title="选择文件" size="650px">
    <div class="btn-list">
      <upload-common v-model:imageCommon="imageCommon" @on-success="open" />
      <!-- 子组件 emit('on-success', data.file.url)， 触发父组件@on-success="getTableData" -->
      <upload-image
        v-model:imageUrl="imageUrl"
        :file-size="512"
        :max-w-h="1080"
        @on-success="open"
      />
      <el-input
        v-model="search.keyword"
        class="keyword"
        placeholder="请输入文件名"
      />
      <el-button type="primary" icon="search" @click="open">查询</el-button>
    </div>
    <div class="media">
      <div v-for="(item, key) in picList" :key="key" class="media-box">
        <div class="header-img-box-list">
          <el-image
            :key="key"
            :src="getUrl(item.url)"
            @click="chooseImg(item.url, target, targetKey, item)"
          >
            <template #error>
              <div class="header-img-box-list">
                <el-icon>
                  <picture />
                </el-icon>
              </div>
            </template>
          </el-image>
        </div>
        <div class="img-title" @click="editFileNameFunc(item)">
          {{ item.name }}
        </div>
      </div>
    </div>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :total="total"
      :style="{ 'justify-content': 'center' }"
      layout="total, prev, pager, next, jumper"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
  </el-drawer>
</template>

<script setup>
import { getUrl } from '@/utils/image'
import { ref } from 'vue'
import { getFileList, editFileName } from '@/api/fileUploadAndDownload'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const imageUrl = ref('')
const imageCommon = ref('')

const search = ref({})
const page = ref(1)
const total = ref(0)
const pageSize = ref(20)

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  open()
}

const handleCurrentChange = (val) => {
  page.value = val
  open()
}

const emit = defineEmits(['enterImg', 'updateFiles']) // 新增：定义emit事件，用于触发父组件事件（enterImg,
defineProps({
  target: {
    type: Object,
    default: null
  },
  targetKey: {
    type: String,
    default: ''
  }
})

const drawer = ref(false)
const picList = ref([])
// 选择图片
const chooseImg = (url, target, targetKey, item) => {
  if (target && targetKey) {
    target[targetKey] = url
  }
  emit('enterImg', url) // 触发父组件enterImg事件
  emit('updateFiles', item) // 新增：触发父组件updateFiles事件 传递选中的整个file。 父组件用：@updateFiles="updateFiles" 实现监听
  drawer.value = false
}

const open = async () => {
  const res = await getFileList({
    page: page.value,
    pageSize: pageSize.value,
    ...search.value
  })
  if (res.code === 0) {
    picList.value = res.data.list
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
    drawer.value = true
  }
}

/**
 * 编辑文件名或者备注
 * @param row
 * @returns {Promise<void>}
 */
const editFileNameFunc = async (row) => {
  ElMessageBox.prompt('请输入文件名', '编辑', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /\S/,
    inputErrorMessage: '不能为空',
    inputValue: row.name
  })
    .then(async ({ value }) => {
      row.name = value
      // console.log(row)
      const res = await editFileName(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '编辑成功!'
        })
        open()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消修改'
      })
    })
}

defineExpose({ open })
</script>

<style lang="scss">
.upload-btn-media-library {
  margin-left: 20px;
}

.media {
  display: flex;
  flex-wrap: wrap;

  .media-box {
    width: 120px;
    margin-left: 20px;

    .img-title {
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      line-height: 36px;
      text-align: center;
      cursor: pointer;
    }

    .header-img-box-list {
      width: 120px;
      height: 120px;
      border: 1px dashed #ccc;
      border-radius: 8px;
      text-align: center;
      line-height: 120px;
      cursor: pointer;
      overflow: hidden;
      .el-image__inner {
        max-width: 120px;
        max-height: 120px;
        vertical-align: middle;
        width: unset;
        height: unset;
      }
    }
  }
}
</style>
