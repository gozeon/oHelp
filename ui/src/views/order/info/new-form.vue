<template>
  <el-form v-loading="loading" style="padding: 0 20px" ref="form" label-position="top" :model="form" :rules="rules" size="small">

    <el-form-item label="描述" prop="description">
      <my-tinymce id="editor1" v-model="form.description"></my-tinymce>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="handerSubmit">提交</el-button>
    </el-form-item>
  </el-form>
</template>
<script>
import MyTinymce from '../../../components/my-tinymce.vue'
export default {
  components: {
    MyTinymce,
  },
  data() {
    return {
      loading: false,
      form: {
        description: '',
      },
      rules: {
        description: [{ required: true, message: '必填项', trigger: 'blur' }],
      },
    }
  },
  async mounted() {
  },
  methods: {
    handerSubmit() {
      this.$refs.form.validate((valid) => {
        
        if (valid) {
          this.loading = true
          this.axios
            .post('/comment/', {
              ...this.form,
              order_id: +this.$route.params.id
            })
            .then((res) => {
              const { status, msg } = res.data
              if (status !== 0) {
                this.$message.error(msg || '系统错误，请重试')
              } else {
                this.$message.success('提交成功')
                this.$emit('edit-success')
              }
              this.loading = false
            })
        }
      })
    },
    
  },
}
</script>
<style lang="scss" scoped></style>
