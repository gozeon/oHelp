<template>
  <el-row>
    <el-row type="flex" justify="end">
      <router-link :to="{ path: '/order/order-create' }" tag="span">
        <el-button style="width: 90px" icon="el-icon-plus" type="primary" size="small">新建</el-button>
      </router-link>
    </el-row>
    <el-form ref="searchFormRef" :model="searchForm" :rules="{}" size="mini" inline>
      <el-form-item label="ID" prop="id">
        <el-input v-model="searchForm.id" placeholder="ID"></el-input>
      </el-form-item>
      <el-form-item label="标签" prop="tag">
        <el-input v-model="searchForm.tag" placeholder="标签"></el-input>
      </el-form-item>

      <el-form-item label="状态" prop="status">
        <el-select v-model="searchForm.status" placeholder="状态">
          <el-option v-for="(item, index) in statusOptions" :key="index" :label="item.label" :value="item.value"> </el-option>
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
      <el-table-column prop="status" label="状态">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status ? 'success' : 'danger'" effect="dark">{{ scope.row.status ? '开启' : '关闭' }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="tags" label="标签"> </el-table-column>
      <el-table-column prop="create_user" label="创建人"></el-table-column>
      <el-table-column prop="created_at" label="创建时间" :formatter="timeFormater"></el-table-column>
      <el-table-column prop="updated_at" label="修改时间" :formatter="timeFormater"></el-table-column>
      <el-table-column prop="id" fixed="right" label="操作" width="160">
        <template slot-scope="scope">
          <router-link :to="{ path: `/order/order-info/${scope.row.id}` }" tag="span" style="margin-right: 10px">
            <el-button  type="text" size="small">详情</el-button>
          </router-link>


          <el-popconfirm v-if="scope.row.status" title="确定关闭该项?" @confirm="updateStatus({ id: scope.row.id, status: false, successMessage: '关闭成功', source: scope.row })">
            <el-button style="color: #f56c6c" type="text" slot="reference" size="small">关闭</el-button>
          </el-popconfirm>
          <el-popconfirm v-else title="确定重新开启该项?" @confirm="updateStatus({ id: scope.row.id, status: true, successMessage: '开启成功', source: scope.row })">
            <el-button type="text" slot="reference" size="small">开启</el-button>
          </el-popconfirm>

        </template>
      </el-table-column>
    </el-table>

    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage" :page-sizes="[10, 20, 50, 100]" :page-size="pageSize" layout="total, sizes, prev, pager, next" :total="pageTotal"> </el-pagination>
  </el-row>
</template>
<script>
import { format, parseISO } from 'date-fns'
export default {
  data() {
    return {
      statusOptions: [
        { label: '开启', value: 1 },
        { label: '关闭', value: 0 },
      ],
      searchForm: {},
      loading: false,
      pageTotal: 0,
      pageSize: 10,
      currentPage: 1,
      tableData: [],
    }
  },

  computed: {},
  mounted() {
    this.loadData()
  },
  methods: {
    timeFormater(row, column, cellValue, index) {
      return cellValue ? format(parseISO(cellValue), 'yyyy-MM-dd HH:mm:ss') : cellValue
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
    updateStatus({ id, status, successMessage,source }) {
      this.axios
        .put(
          `/order/status/${id}`,
          {...source, status}
        )
        .then((res) => {
          const { status, msg } = res.data
          if (status !== 0) {
            this.$message.error(msg || '系统错误，请重试')
          } else {
            this.$message.success(successMessage || '操作成功')
            this.loadData()
          }
        })
    },
    loadData() {
      this.loading = true
      this.axios
        .get('/order/', {
          params: {
            page: this.currentPage,
            size: this.pageSize,
            ...this.searchForm,
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
