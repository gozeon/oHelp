<template>
  <div class="header">
    <div class="logo" :class="{ collapse: collapse }">
      {{ collapse ? 'OH' : 'oHelp' }}
    </div>
    <div class="bar">
      <div class="left">
        <i class="iconfont" :class="collapse ? 'icon-zhankaicaidan' : 'icon-shouqicaidan'" @click="toggleCollapse"></i>
        <el-input placeholder="请输入关键字搜索" v-model="search" size="mini" prefix-icon="el-icon-search" clearable> </el-input>
      </div>
      <div class="right">
        <i class="iconfont icon-bangzhu"></i>
        <el-drawer title="消息通知" :visible.sync="drawer" :modal="false" direction="rtl" append-to-body>
          <el-empty v-if="messages.length < 1" description="空空如也"></el-empty>
        </el-drawer>
        <el-badge :value="200" :max="99" @click="showMessage">
          <i class="iconfont icon-lingdang" @click="showMessage"></i>
        </el-badge>
        <el-dropdown trigger="click" @command="handleCommand">
          <div class="avatar">
            <img src="../../assets/img/ic_head_default.png" />
            <i class="el-icon-arrow-down el-icon--right"></i>
          </div>

          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>{{ user.uname }}</el-dropdown-item>
            <el-dropdown-item command="logout" divided>退出</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>
<script>
import { mapState } from 'vuex'
import MyMessage from './my-message.vue'
export default {
  components: {
    MyMessage,
  },
  data() {
    return {
      search: '',
      drawer: false,
      messages: [],
    }
  },
  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('global', {
      collapse: 'collapse',
    }),
    ...mapState('user', {
      user: 'user',
    }),
  },
  created() {
    // TODO: get message
    // this.axios.get('/log/getLogPage').then(({ data }) => {
    //   this.messages = data.data
    // })
  },
  methods: {
    toggleCollapse() {
      this.$store.commit('global/toggleSideBarCollapse')
    },
    showMessage() {
      this.drawer = true
    },
    handleCommand(command) {
      if (command === 'logout') {
        this.axios.get('/logout').then((response) => {
          const { data, msg, status } = response.data
          if (+status === 0) {
            location.href = data.url
          }
        })
      }
    },
  },
}
</script>
<style lang="scss" scoped>
.header {
  height: 50px;
  background: var(--white-500);
  display: flex;

  .logo {
    width: 160px;
    flex-shrink: 0;
    background: var(--N2);

    display: flex;
    align-items: center;
    justify-content: center;

    &.collapse {
      width: 48px;
    }
    transition: width 0.3s;
  }

  .bar {
    flex: 1;
    height: 100%;
    border-bottom: 1px solid var(--N3);
    box-sizing: border-box;
    padding: 0 32px;

    display: flex;
    justify-content: space-between;

    .left {
      display: flex;
      align-items: center;

      i {
        cursor: pointer;
        margin-right: 32px;
      }

      /deep/ .el-input input {
        width: 230px;
        height: 28px;
        border: 0;
        border-radius: 1px;
        background: var(--N3);
      }
    }

    .right {
      display: flex;
      align-items: center;

      i {
        margin-right: 20px;
        cursor: pointer;
      }

      /deep/ .el-badge .el-badge__content {
        right: 27px;
        z-index: 20;
      }

      .avatar {
        cursor: pointer;
        margin-left: 20px;
        display: flex;
        align-items: center;
        img {
          width: 20px;
          height: 20px;
        }
      }
    }
  }
}

.el-dropdown-menu {
  border-radius: 1px;

  .el-dropdown-menu__item {
    width: 170px;

    &:hover {
      background-color: var(--green-200A);
      color: var(--white-500);
    }
  }
}

.el-drawer__wrapper {
  /deep/ .el-drawer__container {
    top: 60px;

    .el-drawer__body {
      border-top: 1px solid var(--neutral-10);
    }
  }
}
</style>
