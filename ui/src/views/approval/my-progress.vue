<template>
  <el-row v-loading="loading" style="padding: 0px 40px">
    <el-timeline>
      <el-timeline-item v-for="(item, index) in data" :key="item.id" :timestamp="timestampFormat(item.create_user, item.created_at)" placement="top" :color='index == 0 ? "#0bbd87": ""'>
        <el-card>
          <div v-html="item.opinion"></div>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </el-row>
</template>
<script>
import { format, parseISO } from 'date-fns'

export default {
  props: ['id'],
  data() {
    return {
      loading: true,
      data: [],
    }
  },

  created() {
    this.loadData()
  },

  methods: {
    timestampFormat(name, time) {
      return `${name}         ${format(parseISO(time), 'yyyy-MM-dd HH:mm:ss')}`
    },
    loadData() {
      this.loading = true
      this.axios.get(`/progress/inApproval/${this.id}`).then((response) => {
        const { status, msg, data } = response.data
        if (status === 0) {
          this.data = { ...data }
        } else {
          this.$message.error(msg || '系统错误，请重试')
        }
        this.loading = false
      })
    },
  },
}
</script>

<style lang="scss" scoped>
/deep/ .el-timeline-item__timestamp {
  //   padding-top: 0;
}
</style>
