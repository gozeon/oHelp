<template>

  <el-row v-loading="loading">
    <el-row style="padding-left: 40px" type="flex" align="middle" tag="div" >
      <el-tag :type="data.status ? 'success' : 'danger'" effect="dark" style="margin-right:10px">{{ data.status ? '开启' : '关闭' }}</el-tag>
      <h1>{{data.title}}</h1>
    </el-row>

    <el-row style="padding:0 0 20px 40px" v-if="tags.length">
      <el-tag v-for="item in tags" :key="item" size="small" effect="plain" style="margin-right: 10px" @close="handleRemove(item)"> {{ item }}</el-tag>
    </el-row>

    <el-timeline v-if="data.id">
      <el-timeline-item :timestamp="timestampFormat(data.create_user, data.created_at)" placement="top">
        <el-card>
          <div  v-html="data.description"></div>
        </el-card>
      </el-timeline-item>

      <template v-if="comments.length">
        <el-timeline-item v-for="item in comments" :key="item.id" :timestamp="timestampFormat(item.create_user, item.created_at)" placement="top">
          <el-card>
            <div  v-html="item.description"></div>
          </el-card>
        </el-timeline-item>
      </template>
      
    </el-timeline>

    <el-row style="padding:0 0 20px 40px">
      <el-button type="primary" style="width:90px;" size="small" @click="drawer = true">添加评论</el-button>
    </el-row>

    <el-drawer direction="rtl" append-to-body :visible.sync="drawer" size="50%" :title="drawerTitle" @close="loadData">
      <new-form v-if="drawer" @edit-success="drawer = false"></new-form>
    </el-drawer>
  </el-row>
</template>
<script>
import split from 'lodash/split'
import { format, parseISO } from 'date-fns'
import NewForm from './new-form.vue'
export default {
  components: {NewForm},
  data() {
    return {
     drawer: false,
      drawerTitle: '新建评论',
      loading: true,
      data: {},
      form: {
        description: ''
      },
      comments: []
    }
  },
  computed:{
    tags() {
      if(this.data?.tags) {
        return split(this.data?.tags, ',')
      } else {
        return []
      }
    }
  },
  created(){
    this.loadData()
  },
  methods: {
    timestampFormat(name, time) {
      return `${name}         ${format(parseISO(time), 'yyyy-MM-dd HH:mm:ss')}`
    },
   loadData() {
      this.loading = true
      this.axios
        .get(`/order/${this.$route.params.id}`)
        .then((response) => {
          const { status, msg, data } = response.data
          this.loading = false
          if (status === 0) {
              this.data = {...data}
            } else {
              this.$message.error(msg || '系统错误，请重试')
            }
        })

      this.axios
        .get(`/comment/inOrder/${this.$route.params.id}`)
        .then((response) => {
          const { status, msg, data } = response.data
          this.comments = [...data]
        })
    },
  },
}
</script>
<style lang="scss" scoped>
/deep/ .el-timeline-item__timestamp {
  padding-top: 0;
}
</style>