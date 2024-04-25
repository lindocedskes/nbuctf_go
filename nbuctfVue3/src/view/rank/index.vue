<template>
  <div>
    <div
      class="table-box w-4/5 mx-auto relative overflow-x-auto shadow-md sm:rounded-lg"
    >
      <!-- 折线图 -->
      <el-row style="padding: 16px 16px 0; margin: 12px auto; max-width: 100%">
        <line-chart :chartData="transformData(chartData)"></line-chart>
      </el-row>
      <!-- 查询 -->
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
      <!-- 点击切换榜单 -->
      <el-tabs v-model="rankType" @tab-click="setRankType">
        <el-tab-pane
          v-for="type in types"
          :key="type.value"
          :label="type.label"
          :name="type.value"
        ></el-tab-pane>
      </el-tabs>

      <el-table :data="tableData" class="mx-auto flex">
        <el-table-column label="排名" prop="rank"> </el-table-column>
        <el-table-column label="用户名" prop="userName" class="flex-1">
        </el-table-column>
        <el-table-column label="总得分" prop="score" class="flex-1" />
        <el-table-column
          label="最后提交时间"
          prop="UpdatedAt"
          class="flex-1"
          width="200"
        >
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>
        <el-table-column label="Crypto" prop="scoreCrypto" class="flex-1" />
        <el-table-column label="Web" prop="scoreWeb" class="flex-1" />
        <el-table-column label="Misc" prop="scoreMisc" class="flex-1" />
        <el-table-column label="Pwn" prop="scorePwn" class="flex-1" />
        <el-table-column label="Reverse" prop="scoreReverse" class="flex-1" />
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
import { submitScoreChart } from '@/api/game.js'

import { ref, watch } from 'vue'
import { formatDate } from '@/utils/format'
import { ElMessage } from 'element-plus'
import LineChart from '@/components/chart/LineChart.vue'

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
const types = [
  { value: 'Score', label: '总榜' },
  { value: 'Crypto', label: 'Crypto榜' },
  { value: 'Web', label: 'Web榜' },
  { value: 'Misc', label: 'Misc榜' },
  { value: 'Pwn', label: 'Pwn榜' },
  { value: 'Reverse', label: 'Reverse榜' }
]
const rankType = ref('Score')
const setRankType = (tab) => {
  console.log('Setting rankType.value to:', tab)
  console.log('Type of rankType.value:', tab.paneName)
  rankType.value = tab.paneName
}
// 查询
const getTableData = async () => {
  const table = await getRankList({
    page: page.value,
    pageSize: pageSize.value,
    type: rankType.value
  })
  if (table.code === 0) {
    // 添加排名，rank属性，前端实现计算排名
    table.data.list.forEach((item, index) => {
      item.rank = index + 1 + (page.value - 1) * pageSize.value
    })

    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
const chartData = ref([])

// 查询
const getChartData = async () => {
  const res = await submitScoreChart()
  if (res.code === 0) {
    chartData.value = res.data
    console.log('chartData:', chartData.value)
  }
}

const initPage = async () => {
  getTableData()
  getChartData()
}
initPage()

const findRank = async () => {
  // 查询所有记录
  const table = await getRankList({
    page: 1,
    pageSize: 9999, // 一个足够大的数，以确保获取所有记录
    type: rankType.value
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
      await getTableData()
    }
  }
}
// rankType.value 改变时，重新获取数据
watch(
  () => rankType.value,
  () => {
    if (search.value.keyword) {
      console.log('search.value.keyword', search.value.keyword)
      findRank()
    } else {
      console.log(rankType.value)
      getTableData()
    }
  }
)
//数据转换为适合图表的格式，键是用户名，值是一个数组，数组中的每个元素都是一个包含提交时间和问题分数的数组
function transformData(data) {
  const users = {}
  data.forEach((item) => {
    if (!users[item.userName]) {
      users[item.userName] = []
    }
    //获取累积得分
    const lastScore =
      users[item.userName].length > 0
        ? users[item.userName][users[item.userName].length - 1][1]
        : 0
    users[item.userName].push([
      formatDate(item.submitTime),
      item.queMark + lastScore,
      item.queName
    ])
  })
  console.log('users累积积分折线数据:', users)
  return users
}
</script>

<style scoped>
.name {
  cursor: pointer;
}

.table-box {
  @apply p-6 bg-white rounded;
  .btn-list {
    .is-active {
      @apply rounded text-white;
      background: var(--el-color-primary) !important;
      color: #e21c1c !important;
    }
  }
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
    background: var(--el-color-primary) !important;
    color: #ffffff !important;
  }
}
</style>
