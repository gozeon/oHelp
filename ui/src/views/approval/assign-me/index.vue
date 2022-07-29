<template>
  <my-table v-if="user.uname" :approver="user.uname"></my-table>
</template>
<script>
import myTable from '../my-table.vue'
import { mapState } from 'vuex'
export default {
  components: { myTable },
  computed: {
    // mix the getters into computed with object spread operator
    ...mapState('user', {
      user: 'user',
    }),
  },
  data() {
    return {
    }
  },
  
}
</script>
<style lang="scss" scoped></style>
