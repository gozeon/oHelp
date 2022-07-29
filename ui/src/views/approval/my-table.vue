<template>
  <el-row>
    <el-row type="flex" justify="end">
      <router-link :to="{ path: '/approval/approval-new' }" tag="span">
        <el-button style="width: 90px" icon="el-icon-plus" type="primary" size="small">新建</el-button>
      </router-link>
    </el-row>
    <el-form ref="searchFormRef" :model="searchForm" :rules="{}" size="mini" inline>
      <el-form-item label="ID" prop="id">
        <el-input v-model="searchForm.id" placeholder="ID"></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="searchForm.status" placeholder="状态">
          <el-option v-for="(item, index) in approvalStatusOptions" :key="index" :label="item.label" :value="item.value"> </el-option>
        </el-select>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="handerSearch">查询</el-button>
        <el-button @click="handerReset">重置</el-button>
      </el-form-item>
    </el-form>
    <el-table ref="tableEl" @header-dragend="$refs.tableEl.doLayout()" :data="tableData" v-loading="loading" height="634px" style="width: 100%" :header-cell-style="{ background: 'var(--N4)' }" border>
      <el-table-column prop="id" label="ID"> </el-table-column>
      <el-table-column prop="title" label="标题"> </el-table-column>
      <el-table-column prop="status" label="状态" :formatter="statusFormater"></el-table-column>
      <el-table-column prop="approver" label="审批人"></el-table-column>
      <el-table-column prop="create_user" label="创建人"></el-table-column>
      <el-table-column prop="created_at" label="创建时间" :formatter="timeFormater"></el-table-column>
      <el-table-column prop="updated_at" label="修改时间" :formatter="timeFormater"></el-table-column>
      <el-table-column prop="id" fixed="right" label="操作" width="120">
        <template slot-scope="scope">
          <el-button
            type="text"
            size="small"
            @click="
              activeId = scope.row.id
              drawer = true
            "
            >详情</el-button
          >
          <el-button
            type="text"
            size="small"
            @click="
              activeId = scope.row.id
              drawer1 = true
            "
            >进度</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage" :page-sizes="[10, 20, 50, 100]" :page-size="pageSize" layout="total, sizes, prev, pager, next" :total="pageTotal"> </el-pagination>

    <el-drawer direction="ltr" append-to-body :visible.sync="drawer" size="50%" title="审批详情" @close="loadData">
      <my-info v-if="drawer" :id="activeId" @edit-success="drawer = false"></my-info>
    </el-drawer>

    <el-drawer direction="ltr" append-to-body :visible.sync="drawer1" size="50%" title="进度详情" @close="loadData">
      <my-progress v-if="drawer1" :id="activeId" @edit-success="drawer1 = false"></my-progress>
    </el-drawer>
  </el-row>
</template>
<script>
import { format, parseISO } from 'date-fns'
import { mapState } from 'vuex'
import filter from 'lodash/filter'
import MyInfo from './my-info.vue'
import MyProgress from './my-progress.vue'
export default {
  components: {
    MyInfo,
    MyProgress,
  },
  props: {
    status: {
      type: Number | String,
      default: '',
    },
    showStatus: {
      type: Boolean,
      default: true,
    },
    approver: {
      type: String,
      default: '',
    },
    createUser: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      searchForm: {},
      loading: false,
      pageTotal: 0,
      pageSize: 10,
      currentPage: 1,
      tableData: [],
      activeId: '',
      drawer: false,
      drawer1: false,
    }
  },

  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('global', {
      approvalStatusOptions: (state) => state.options.approvalStatusOptions,
    }),
  },
  mounted() {
    this.loadData()
  },
  methods: {
    timeFormater(row, column, cellValue, index) {
      return cellValue ? format(parseISO(cellValue), 'yyyy-MM-dd HH:mm:ss') : cellValue
    },
    statusFormater(row, column, cellValue, index) {
      return filter(this.approvalStatusOptions, { value: cellValue })[0]['label']
    },

    handerSearch() {
      this.currentPage = 1
      this.loadData()
    },
    handerReset() {
      this.$refs.searchFormRef.resetFields()
      this.currentPage = 1
      this.loadData()
    },
    handleSizeChange(pageSize) {
      this.pageSize = pageSize
      this.loadData()
    },
    handleCurrentChange(currentPage) {
      this.currentPage = currentPage
      this.loadData()
    },
    loadData() {
      this.loading = true
      this.axios
        .get('/approval/', {
          params: {
            page: this.currentPage,
            size: this.pageSize,
            ...this.searchForm,

            approver: this.approver,
            createUser: this.createUser,
          },
        })
        .then((response) => {
          const { status, msg, data } = response.data
          this.loading = false
          this.pageTotal = data.total

          this.tableData = data.list
        })
    },
  },
}
</script>
<style lang="scss" scoped></style>
