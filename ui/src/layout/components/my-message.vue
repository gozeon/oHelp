<template>
  <div class="message" @click="handleClick">
    <div class="title" :class="{ tip: unred }">{{ title }}</div>
    <div class="des">{{ desc }}</div>
    <div class="time">{{ time }}</div>

    <el-dialog title="通知" :visible.sync="dialog" append-to-body>
      <p>
        王建，约生于唐代宗大历元年（七六六），约卒于唐宪宗元和十年以后，字仲初，颖川（今河南许昌）人。 得推荐，历任太府寺丞、秘书丞，入魏博节度使田弘正幕，其后得裴度力荐，任昭应县丞。 官至陕州司马。 王建有才华，阅历也深，善写塞外和宫闱秘事的题材，工于乐府和歌行，曲调幽怨，意境深远。 有诗集十卷传世
      </p>
      <p>
        王建，约生于唐代宗大历元年（七六六），约卒于唐宪宗元和十年以后，字仲初，颖川（今河南许昌）人。 得推荐，历任太府寺丞、秘书丞，入魏博节度使田弘正幕，其后得裴度力荐，任昭应县丞。 官至陕州司马。 王建有才华，阅历也深，善写塞外和宫闱秘事的题材，工于乐府和歌行，曲调幽怨，意境深远。 有诗集十卷传世
      </p>
      <p>
        王建，约生于唐代宗大历元年（七六六），约卒于唐宪宗元和十年以后，字仲初，颖川（今河南许昌）人。 得推荐，历任太府寺丞、秘书丞，入魏博节度使田弘正幕，其后得裴度力荐，任昭应县丞。 官至陕州司马。 王建有才华，阅历也深，善写塞外和宫闱秘事的题材，工于乐府和歌行，曲调幽怨，意境深远。 有诗集十卷传世
      </p>

      <span slot="footer" class="dialog-footer">
        <el-button type="primary" size="mini" @click="dialog = false">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
export default {
  props: ['id', 'title', 'desc', 'time', 'unred'],
  data() {
    return {
      dialog: false,
    }
  },
  methods: {
    handleClick() {
      this.dialog = true
    },
  },
}
</script>
<style lang="scss" scoped>
.message {
  padding: 12px 24px;

  cursor: pointer;

  .title {
    font-size: 14px;
    line-height: 22px;
    color: var(--neutral-500);
    position: relative;

    &.tip {
      &::before {
        content: '';
        display: inline-block;
        width: 6px;
        height: 6px;
        border-radius: 50%;
        background: var(--red-500);
        position: absolute;
        left: -9px;
        bottom: 9px;
      }
    }
  }
  .des {
    font-size: 12px;
    line-height: 14px;
    color: var(--neutral-100);
  }
  .time {
    font-size: 12px;
    line-height: 20px;
    color: var(--neutral-125);
  }

  border-bottom: 1px solid var(--neutral-10);

  &:last-child {
    border-bottom: 0;
  }
}

/deep/ .el-button {
  &:hover {
    background-color: var(--green-300);
    border-color: var(--green-300);
  }
}
/deep/ .el-button--primary {
  background-color: var(--green-500);
  border-color: var(--green-500);
}
</style>
