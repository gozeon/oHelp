<template>
  <main :style="{ left: collapse ? '48px' : '160px' }" ref="main">
    <transition name="fade-transform" mode="out-in">
      <router-view />
    </transition>
  </main>
</template>
<script>
import Watermask from '../../watermark.js'
import { mapState } from 'vuex'
export default {
  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('global', {
      collapse: 'collapse',
    }),
  },
  mounted() {
    // console.log(this.$refs.main.firstChild)
    // new Watermask({
    //   el: this.$refs.main,
    //   text: '',
    //   logoText: '',
    // })

    this.axios.get('/me').then((response) => {
      const { data, msg, status } = response.data
      if (+status === 0) {
        this.$store.commit('user/setUser', data)
      }
    })
  },
}
</script>
<style lang="scss" scoped>
main {
  position: absolute;

  top: 80px;
  right: 0;
  bottom: 0;
  left: 0;
  transition: left 0.3s;

  overflow: auto;
  padding: 28px 32px;
  z-index: 0;
}
</style>
