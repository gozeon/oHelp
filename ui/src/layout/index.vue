<template>
  <div class="wrapper">
    <my-header></my-header>
    <my-sidebar></my-sidebar>
    <my-breadcrumb></my-breadcrumb>
    <my-main></my-main>
  </div>
</template>
<script>
import MyHeader from './components/my-header.vue'
import MySidebar from './components/my-sidebar.vue'
import MyMain from './components/my-main.vue'
import MyBreadcrumb from './components/my-breadcrumb.vue'
export default {
  components: {
    MyHeader,
    MySidebar,
    MyMain,
    MyBreadcrumb,
  },
}
</script>
<style lang="scss" scoped>
.wrapper {
  background: var(--white-500);
  font-family: var(--font-ui);
  font-size: 14px;
  color: var(--neutral-dark);
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
}
</style>
