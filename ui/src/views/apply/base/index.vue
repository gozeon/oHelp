<template>
  <el-form ref="form" :model="form" label-position="top" :rules="rules" size="small">
    <el-form-item label="申请标题" prop="title">
      <el-input v-model="form.title" maxlength="120" show-word-limit></el-input>
    </el-form-item>

    <el-form-item label="申请内容" prop="description">
      <my-tinymce id="editor" v-model="form.description"></my-tinymce>
    </el-form-item>

    <el-form-item label="审批人" prop="approver">
      <el-select v-model="form.approver" placeholder="请选择">
        <el-option v-for="item in userOption" :key="item.value" :label="item.label" :value="item.value"> </el-option>
      </el-select>
    </el-form-item>

    <el-form-item label="抄送" prop="assigner">
      <el-select v-model="form.assigner" multiple placeholder="请选择">
        <el-option v-for="item in userOption" :key="item.value" :label="item.label" :value="item.value"> </el-option>
      </el-select>
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
import MyTinymce from '../../../components/my-tinymce.vue'
import join from 'lodash/join'
import pick from 'lodash/pick'
import { mapState } from 'vuex'
export default {
  components: {
    MyTinymce,
  },
  data() {
    return {
      form: {
        type: 'base',
        description: '',
      },
      rules: {
        title: [{ required: true, message: '必填项', trigger: ['blur', 'change'] }],
        description: [{ required: true, message: '必填项', trigger: ['blur', 'change'] }],
        approver: [{ required: true, message: '必填项', trigger: ['blur', 'change'] }],
      },
    }
  },
  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('global', {
      userOption: (state) => state.options.users,
    }),
  },
  methods: {
    onSubmit() {
      console.log('submit!')
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.axios
            .post('/approval/', {
              ...this.form,
              assigner: join(this.form.assigner, ','),
              content: JSON.stringify(pick(this.form, ['type', 'description'])),
            })
            .then((res) => {
              console.log(res)
              const { status, msg } = res.data
              if (status === 0) {
                this.$message.success('操作成功')
                this.$router.push({ path: '/approval/approval-my' })
              } else {
                this.$message.error(msg || '系统错误，请重试')
              }
            })
        }
      })
    },
  },
}
</script>
