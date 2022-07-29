<template>
  <el-row type="flex">
    <el-tag v-for="item in items" :key="item" type="info" style="margin-right: 10px" @close="handleRemove(item)" closable> {{ item }}</el-tag>
    <el-input v-model.trim="input" @keyup.enter.native="handleInputConfirm" @blur="handleInputConfirm"></el-input>
  </el-row>
</template>
<script>
import uniq from 'lodash/uniq'
import remove from 'lodash/remove'
import cloneDeep from 'lodash/cloneDeep'
export default {
  props: ['value'],
  data() {
    return {
      input: '',
      items: [],
    }
  },
  created() {
    this.items = cloneDeep(this.value)
  },
  methods: {
    handleInputConfirm() {
      if (this.input) {
        this.items = uniq([...this.items, this.input])
        this.input = ''
      }

      this.$emit('input', this.items)
    },
    handleRemove(item) {
      this.items = remove(this.items, (i) => i !== item)

      this.$emit('input', this.items)
    },
  },
}
</script>
