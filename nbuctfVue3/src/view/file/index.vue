<template>
  <div>
    <div class="table-box">
      <div class="btn-list">
        <upload-common
          v-model:imageCommon="imageCommon"
          @on-success="getTableData"
        />
        <upload-image
          v-model:imageUrl="imageUrl"
          :file-size="512"
          :max-w-h="1080"
          @on-success="getTableData"
        />
        <el-input
          v-model="search.keyword"
          class="keyword"
          placeholder="请输入文件名"
        />
        <el-button type="primary" icon="search" @click="getTableData"
          >查询</el-button
        >
      </div>

      <el-table :data="tableData">
        <el-table-column align="left" label="预览" width="100">
          <template #default="scope">
            <!-- 预览图片 -->
            <CustomPic pic-type="file" :pic-src="scope.row.url" preview />
          </template>
        </el-table-column>
        <el-table-column align="left" label="日期" prop="UpdatedAt" width="180">
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="文件名" prop="name" width="180">
          <!-- 默认插槽，子组件可以访问父组件数据 -->
          <template #default="scope">
            <div class="name" @click="editFileNameFunc(scope.row)">
              {{ scope.row.name }}
              <el-button icon="edit" type="primary" link>编辑</el-button>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="链接" prop="url" min-width="300" />
        <el-table-column align="left" label="标签" prop="tag" width="100">
          <template #default="scope">
            <el-tag
              :type="scope.row.tag === 'jpg' ? 'info' : 'success'"
              disable-transitions
              >{{ scope.row.tag }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="160">
          <template #default="scope">
            <el-button
              icon="download"
              type="primary"
              link
              @click="downloadFile(scope.row)"
              >下载</el-button
            >
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteFileFunc(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :style="{ float: 'right', padding: '20px' }"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  getFileList,
  deleteFile,
  editFileName
} from '@/api/fileUploadAndDownload'
import { downloadImage } from '@/utils/downloadImg'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import { formatDate } from '@/utils/format'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'Upload'
})

const path = ref(import.meta.env.VITE_BASE_API)

const imageUrl = ref('')
const imageCommon = ref('')

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const search = ref({})
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}
//当前页码改变时
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getFileList({
    page: page.value,
    pageSize: pageSize.value,
    ...search.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

const deleteFileFunc = async (row) => {
  ElMessageBox.confirm('此操作将永久删除文件, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      const res = await deleteFile(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除'
      })
    })
}

const downloadFile = (row) => {
  if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
    downloadImage(row.url, row.name) //直接网址下载图片
  } else {
    //前端加载时，F12，出现可调试
    // debugger
    downloadImage(path.value + '/' + row.url, row.name) //本地下载
  }
}

/**
 * 编辑文件名或者备注
 * @param row
 * @returns {Promise<void>}
 */
const editFileNameFunc = async (row) => {
  ElMessageBox.prompt('请输入文件名或者备注', '编辑', {
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
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消修改'
      })
    })
}
</script>

<style scoped>
.name {
  cursor: pointer;
}
.table-box {
  @apply p-6 bg-white rounded;
}
.btn-list {
  @apply mb-3 flex gap-3 items-center;
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
</style>
