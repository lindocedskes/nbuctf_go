<template>
  <div class="chart" :style="{ height: height, width: width }" />
</template>

<script setup>
import { ref, watch, onMounted, onUnmounted } from 'vue'
import * as echarts from 'echarts'
import 'echarts/theme/macarons'

//传入折线图数据
const props = defineProps({
  chartData: {
    type: Object,
    default: null
  },
  className: {
    type: String,
    default: 'chart'
  },
  width: {
    type: String,
    default: '100%'
  },
  height: {
    type: String,
    default: '350px'
  },
  autoResize: {
    type: Boolean,
    default: true
  }
})

let chartData = ref(props.chartData)
const chart = ref(null)
// 观察 props.chartData 的变化
watch(
  () => props.chartData,
  (newChartData) => {
    // 当 props.chartData 变化时，更新 chartData
    chartData.value = newChartData
    // 更新图表
    setOptions(newChartData)
  },
  { deep: true }
)

onMounted(() => {
  // 基于准备好的dom，初始化echarts实例
  console.log('props.chartData:', props.chartData)
  chart.value = echarts.init(document.querySelector('.chart'), 'macarons')
  setOptions(props.chartData)
})
//卸载后清理和释放图表实例，以避免内存泄漏
onUnmounted(() => {
  if (chart.value) {
    chart.value.dispose()
    chart.value = null
  }
})
function createSerie(name, data) {
  return {
    name: name,
    smooth: true,
    type: 'line',
    lineStyle: { width: 2 }, //线宽
    data: data,
    animationDuration: 3000, //始动画的时长ms
    animationEasing: 'cubicInOut' //缓动效果
  }
}
function setOptions(data) {
  //获取 data 对象的所有键
  const legendData = Object.keys(data)
  const series = legendData.map((item) => createSerie(item, data[item])) //描述了数据系列的各种属性
  //设置图表的配置选项
  chart.value.setOption({
    xAxis: {
      type: 'time',
      boundaryGap: false,
      axisTick: { show: false }
    },
    dataZoom: [
      {
        type: 'inside',
        start: 0,
        end: 100
      }
    ],
    grid: {
      left: 10,
      right: 10,
      bottom: 20,
      top: 30,
      containLabel: true
    },
    //提示框组件
    tooltip: {
      trigger: 'item',
      axisPointer: {
        type: 'cross'
      },
      padding: [5, 10],
      formatter: function (params) {
        // console.log('params:', params)
        const param = params.value
        // console.log('param:', param)

        return `题目名称: ${param[2]}`
      }
    },
    yAxis: {
      axisTick: { show: false }
    },
    //图例的数据
    legend: {
      data: legendData
    },
    //图线数据
    series: series
  })
}
</script>

<style>
.chart {
  height: 400px;
  min-width: 100%;
}
</style>
