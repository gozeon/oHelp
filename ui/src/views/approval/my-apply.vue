<template>
  <el-card shadow="nevera">
    <el-form ref="form" :model="form" label-width="100px" :rules="rules" size="small">
      <el-form-item label="审批操作" prop="operate">
        <el-radio v-for="(item, index) in operateOptions" :key="index" v-model="form.operate" :label="item.value">{{ item.label }}</el-radio>
      </el-form-item>

      <el-form-item v-if="form.operate == 2" label="移交到" prop="next">
        <el-select v-model="form.next" placeholder="请选择">
          <el-option v-for="item in userOption" :key="item.value" :label="item.label" :value="item.value"> </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="审批意见" prop="opinion">
        <my-tinymce id="editor" :height="200" v-model="form.opinion"></my-tinymce>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit" style="margin-right: 10px">提交</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
import MyTinymce from '../../components/my-tinymce.vue'
import { mapState } from 'vuex'
import filter from 'lodash/filter'

export default {
  props: ["id"],
  components: { MyTinymce },
  data() {
    return {
      form: {
        operate: 3,
        opinion: '',
      },
      rules: {
        next: [{ required: true, message: '必填项', trigger: 'blur' }],
      },
    }
  },
  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('global', {
      applyOperateOptions: (state) => state.options.applyOperateOptions,
      userOption: (state) => state.options.users,
      user: 'user',
    }),

    operateOptions() {
      return filter(this.applyOperateOptions, (item) => item.value > 1)
    },
  },
  methods: {
    onSubmit() {
      console.log('submit!')
      this.$refs.form.validate((valid) => {
        this.axios
          .post('/progress/', {
            ...this.form,
            approval_id: this.id
          })
          .then((res) => {
            console.log(res)
            const { status, msg } = res.data
            if (status === 0) {
              this.$message.success('操作成功')
              this.$emit('edit-success')
            } else {
              this.$message.error(msg || '系统错误，请重试')
            }
          })
      })
    },
  },
}
</script>
<style lang="scss" scoped></style>
