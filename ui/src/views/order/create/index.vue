<template>
  <el-form ref="form" :model="form" label-position="top" :rules="rules" size="small">
    <el-form-item label="标题" prop="title">
      <el-input v-model="form.title" maxlength="120" show-word-limit></el-input>
    </el-form-item>

    <el-form-item label="标签" prop="tag">
      <my-input-tag v-model="form.tag"></my-input-tag>
    </el-form-item>

    <el-form-item label="描述" prop="description">
      <my-tinymce id="editor" v-model="form.description"></my-tinymce>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="onSubmit" style="margin-right: 10px">创建</el-button>
      <router-link :to="{ path: '/order/order-base' }" tag="span">
        <el-button>取消</el-button>
      </router-link>
    </el-form-item>
  </el-form>
</template>
<script>
import MyInputTag from '../../../components/my-input-tag.vue'
import MyTinymce from '../../../components/my-tinymce.vue'
import join from 'lodash/join'
export default {
  components: {
    MyTinymce,
    MyInputTag,
  },
  data() {
    return {
      form: {
        description: '',
        tag: [],
      },
      rules: {
        title: [{ required: true, message: '必填项', trigger: 'blur' }],
        description: [{ required: true, message: '必填项', trigger: 'blur' }],
      },
    }
  },
  methods: {
    onSubmit() {
      console.log('submit!')
      this.$refs.form.validate((valid) => {
        this.axios
          .post('/order/', {
            ...this.form,
            tags: join(this.form.tag, ','),
          })
          .then((res) => {
            console.log(res)
            const { status, msg } = res.data
            if (status === 0) {
              this.$message.success('操作成功')
              this.$router.push({ path: '/order/order-base' })
            } else {
              this.$message.error(msg || '系统错误，请重试')
            }
          })
      })
    },
  },
}
</script>
