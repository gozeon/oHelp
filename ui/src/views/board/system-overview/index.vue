<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="24">
        <area-chart />
      </el-col>
    </el-row>

    <el-row :gutter="10">
      <el-col :span="12">
        <radar-chart />
      </el-col>
      <el-col :span="12">
        <pie-chart />
      </el-col>
    </el-row>

    <el-row :gutter="10">
      <el-col :span="12">
        <bar-chart />
      </el-col>
      <el-col :span="12">
        <line-chart />
      </el-col>
    </el-row>
  </div>
</template>
<script>
import AreaChart from './components/AreaChart.vue'
import BarChart from './components/BarChart.vue'
import LineChart from './components/LineChart.vue'
import PieChart from './components/PieChart.vue'
import RadarChart from './components/RadarChart.vue'

export default {
  components: {
    AreaChart,
    BarChart,
    LineChart,
    PieChart,
    RadarChart,
  },
}
</script>
