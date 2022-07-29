<template>
  <div>
    <router-link v-for="(item, index) in approvalNewOptions" :key="index" :to="{ path: item.link }" tag="span" style="margin-right: 10px">
      <el-link>{{ item.label }}</el-link>
    </router-link>
  </div>
</template>
<script>
import { mapState } from 'vuex'
export default {
  data() {
    return {
      form: {},
    }
  },
  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('global', {
      approvalNewOptions: (state) => state.options.approvalNewOptions,
    }),
  },
}
</script>
