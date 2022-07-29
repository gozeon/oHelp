<template>
  <div class="side" :class="{ collapse: collapse }">
    <el-menu router :default-active="$route.meta.active || $route.name" :collapse="collapse">
      <el-menu-item index="system-overview" :route="{ name: 'system-overview' }">
        <i class="iconfont icon-shujukanban"></i>
        <span slot="title">看板</span>
      </el-menu-item>

      <el-submenu index="1">
        <template slot="title">
          <i class="iconfont icon-dingdan"></i>
          <span>工单管理</span>
        </template>

        <el-menu-item index="order-base" :route="{ name: 'order-base' }">
          <span slot="title">工单列表</span>
        </el-menu-item>
      </el-submenu>

      <el-submenu index="2">
        <template slot="title">
          <i class="el-icon-s-check"></i>
          <span>审批管理</span>
        </template>

        <el-menu-item index="approval-new" :route="{ name: 'approval-new' }">
          <span slot="title">发起审批</span>
        </el-menu-item>

        <el-menu-item index="approval-my" :route="{ name: 'approval-my' }">
          <span slot="title">我发起的</span>
        </el-menu-item>

        <el-menu-item index="approval-assign-me" :route="{ name: 'approval-assign-me' }">
          <span slot="title">我审批的</span>
        </el-menu-item>

        <el-menu-item index="approval-base" :route="{ name: 'approval-base' }">
          <span slot="title">所有审批</span>
        </el-menu-item>
      </el-submenu>
    </el-menu>
  </div>
</template>
<script>
import { mapState } from 'vuex'
export default {
  data() {
    return {}
  },
  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('global', {
      collapse: 'collapse',
    }),
  },
  created() {},
  methods: {},
}
</script>
<style lang="scss" scoped>
.side {
  width: 160px;
  position: fixed;
  left: 0;
  top: 50px;
  bottom: 0;
  overflow-y: auto;
  overflow-x: hidden;
  -webkit-transition: width 0.3s;
  transition: width 0.3s;
  background: var(--N2);

  &.collapse {
    width: 48px;
  }

  /deep/.el-menu {
    background: var(--N2);
    border: 0;
  }

  /deep/.el-menu-item {
    // font-size: 12px;

    height: 40px;
    line-height: 40px;

    min-width: unset;

    border-radius: 2px;
    margin: 4px;
    &:hover {
      background: var(--N4);
    }

    i {
      margin-right: 5px;
    }

    &.is-active {
      color: var(--N1);
      background: var(--green-base);
    }

    * {
      vertical-align: unset;
    }
  }

  /deep/.el-submenu__title {
    // font-size: 12px;
    margin: 4px;
    height: 40px;
    line-height: 40px;
    border-radius: 2px;

    i {
      &:first-child {
        margin-right: 5px;
      }
    }

    &:hover {
      background: var(--N4);
    }
  }

  /deep/.el-menu--collapse {
    width: 48px;

    .el-menu-item {
      border-radius: 2px;

      i {
        margin-right: 0;
      }
      .el-tooltip {
        padding: 0 !important;
        text-align: center;
      }
    }

    .el-submenu {
      margin: 3px;
      border-radius: 2px;
      i {
        margin-right: 0;
      }
      &.is-active {
        color: var(--N1);
        i {
          color: var(--N1);
        }
        background: var(--green-base);
      }

      &:hover {
        background: var(--green-base);
      }
    }

    .el-submenu__title {
      padding: 0 !important;
      text-align: center;

      &:hover {
        background: var(--green-base);
      }
    }
  }
}

.el-menu--vertical {
  .el-menu-item {
    height: 40px;
    line-height: 40px;

    &.is-active {
      color: var(--N1);
      background: var(--green-base);
    }

    &:hover {
      background: var(--N4);
    }
  }
}
</style>
