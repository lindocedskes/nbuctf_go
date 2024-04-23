<template>
  <div>
    <div
      class="table-box w-2/3 mx-auto relative overflow-x-auto shadow-md sm:rounded-lg"
    >
      <div class="btn-list">
        <el-input
          v-model="search.keyword"
          class="keyword"
          placeholder="查询用户排名"
        />
        <el-button type="primary" icon="search" @click="findRank"
          >查询</el-button
        >
      </div>

      <el-table :data="tableData" class="mx-auto flex">
        <el-table-column label="排名" prop="rank" class="flex-1">
        </el-table-column>
        <el-table-column label="用户名" prop="userName" class="flex-1">
        </el-table-column>
        <el-table-column label="得分" prop="score" class="flex-1" />
        <el-table-column label="最后提交时间" prop="UpdatedAt" class="flex-1">
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
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
import { getRankList } from '@/api/rank.js'
import { ref } from 'vue'
import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'

const page = ref(1)
const total = ref(0)
const pageSize = ref(20)
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
  const table = await getRankList({
    page: page.value,
    pageSize: pageSize.value,
    ...search.value
  })
  if (table.code === 0) {
    // 添加排名，rank属性
    table.data.list.forEach((item, index) => {
      item.rank = index + 1 + (page.value - 1) * pageSize.value
    })

    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const initPage = async () => {
  getTableData()
}
initPage()

const findRank = async () => {
  // 查询所有记录
  const table = await getRankList({
    page: 1,
    pageSize: 9999 // 一个足够大的数，以确保获取所有记录
  })

  if (table.code === 0) {
    // 添加排名，rank属性
    table.data.list.forEach((item, index) => {
      item.rank = index + 1
    })

    // 找到指定用户的排名
    const user = table.data.list.find(
      (item) => item.userName === search.value.keyword
    )
    if (user) {
      // 显示指定用户的排名
      tableData.value = [user]
      total.value = 1
      page.value = 1
      pageSize.value = 1
    } else {
      //报错 查询失败
      ElMessage({
        type: 'error',
        message: '用户未参加比赛 或 输入框为空'
      })
      // 内容为空，显示所有用户
      page.value = 1
      pageSize.value = 20
      search.value.keyword = '' // 清空搜索关键字
      console.log('search.value.keyword', search.value.keyword)
      await getTableData()
    }
  }
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
