<template>
  <div class="breadcrumb" :style="{ left: collapse ? '48px' : '160px' }">
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>

      <template v-for="(item, index) in breadcrumb">
        <el-breadcrumb-item v-if="item.text" :key="index" :to="item.to">{{ item.text }}</el-breadcrumb-item>
      </template>
    </el-breadcrumb>
  </div>
</template>
<script>
import { mapState } from 'vuex'
export default {
  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('global', {
      collapse: 'collapse',
    }),
    breadcrumb: function () {
      const { matched } = this.$route
      return matched.map((item) => {
        var obj = {}
        if (item?.redirect) {
          obj.to = { name: item?.meta.active || item?.name }
        }
        return { ...obj, ...item?.meta }
      })
    },
  },
}
</script>
<style lang="scss" scoped>
.breadcrumb {
  position: absolute;
  padding-left: 20px;

  top: 50px;
  right: 0;
  left: 0;
  background: var(--white-500);
  border-bottom: 1px solid var(--N3);
  height: 30px;
  display: flex;
  align-items: center;
  transition: left 0.3s;

  /deep/ .el-breadcrumb {
    .is-link {
      &:hover {
        color: var(--green-base);
      }
    }
  }
}
</style>
