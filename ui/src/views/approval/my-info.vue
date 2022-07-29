<template>
  <el-row v-loading="loading" style="padding: 0px 40px">
    <el-row type="flex" align="middle" tag="div">
      <el-tag :type="tag.type" effect="dark" style="margin-right: 10px">{{ tag.label }}</el-tag>
      <h1>{{ data.title }}</h1>
    </el-row>

    <el-timeline v-if="data.id">
      <el-timeline-item :timestamp="timestamp" placement="top">
        <template v-if="content.type === 'base'">
          <el-card>
            <div v-html="content.description"></div>
          </el-card>
        </template>

        <template v-if="content.type === 'sign'">
          <el-card>
            <el-descriptions title="申请补签" :column="1" border>
              <el-descriptions-item label="用户手机号">{{content.phone}}</el-descriptions-item>
              <el-descriptions-item label="签到时间">{{content.time}}</el-descriptions-item>
              <el-descriptions-item label="原因">{{content.reason}}</el-descriptions-item>
            </el-descriptions>
          </el-card>
        </template>
      </el-timeline-item>

      <template v-if="comments.length">
        <el-timeline-item v-for="item in comments" :key="item.id" :timestamp="timestampFormat(item.create_user, item.created_at)" placement="top">
          <el-card>
            <div  v-html="item.opinion"></div>
          </el-card>
        </el-timeline-item>
      </template>

      <template v-if="data.status == 1">
        <el-timeline-item>
          <el-card>
            <div>{{data.approver}} 审批中</div>
          </el-card>
        </el-timeline-item>
      </template>

    </el-timeline>

    <el-row v-if="showApply">
      <my-apply v-if="id" :id="id" @edit-success="close" />
    </el-row>
  </el-row>
</template>
<script>
import { format, parseISO } from 'date-fns'
import { mapState } from 'vuex'
import filter from 'lodash/filter'
import get from 'lodash/get'
import myApply from './my-apply.vue'

export default {
  components: { myApply },
  props: ['id'],
  data() {
    return {
      loading: true,
      data: {},
      tag: {
        type: '',
        label: '',
      },
      comments:[]
    }
  },
  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('global', {
      approvalStatusOptions: (state) => state.options.approvalStatusOptions,
    }),
    ...mapState('user', {
      user: 'user',
    }),
    content() {
      return JSON.parse(get(this.data, 'content') || '{}')
    },
    timestamp() {
      return `申请人:  ${this.data.create_user}         ${format(parseISO(this.data.created_at), 'yyyy-MM-dd HH:mm:ss')}`
    },
    showApply() {
      return this.data.status == 1 && this.user.uname === this.data.approver
    }
  },
  watch: {
    data: function (val) {
      this.tag = {
        type: ['', 'info', 'success', 'danger'][this.data.status],
        label: filter(this.approvalStatusOptions, { value: this.data.status })[0]['label'],
      }
    },
  },

  created() {
    this.loadData()
  },

  methods: {
    timestampFormat(name, time) {
      return `${name}         ${format(parseISO(time), 'yyyy-MM-dd HH:mm:ss')}`
    },
    async loadData() {
      this.loading = true
      await this.axios.get(`/approval/${this.id}`).then((response) => {
        const { status, msg, data } = response.data
        if (status === 0) {
          this.data = { ...data }
        } else {
          this.$message.error(msg || '系统错误，请重试')
        }
        this.loading = false
      })

     await this.axios.get(`/progress/inApproval/${this.id}`, {params: {operate:5,updated_at:'asc'}}).then((response) => {
        const { status, msg, data } = response.data
        if (status === 0) {
          this.comments = [...this.comments,...data]
        } else {
          this.$message.error(msg || '系统错误，请重试')
        }
        this.loading = false
      })

     await this.axios.get(`/progress/inApproval/${this.id}`, {params: {operate:3,updated_at:'asc'}}).then((response) => {
        const { status, msg, data } = response.data
        if (status === 0) {
          this.comments = [...this.comments,...data]
        } else {
          this.$message.error(msg || '系统错误，请重试')
        }
        this.loading = false
      })

     await this.axios.get(`/progress/inApproval/${this.id}`, {params: {operate:4,updated_at:'asc'}}).then((response) => {
        const { status, msg, data } = response.data
        if (status === 0) {
          this.comments = [...this.comments,...data]
        } else {
          this.$message.error(msg || '系统错误，请重试')
        }
        this.loading = false
      })
    },
    close() {
      console.log(11111)
      this.$emit('edit-success')
    }
  },
}
</script>

<style lang="scss" scoped>
/deep/ .el-timeline-item__timestamp {
  padding-top: 0;
}
</style>
