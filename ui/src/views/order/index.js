import Layout from '../../layout/index.vue'

export default {
  path: '/order',
  name: 'order',
  component: Layout,
  redirect: '/order/order-base',
  meta: {
    text: '工单管理',
  },
  children: [
    {
      path: 'order-base',
      name: 'order-base',
      meta: {
        text: '工单列表',
      },
      component: () => import('./base/index.vue'),
    },
    {
      path: 'order-create',
      name: 'order-base',
      meta: {
        text: '创建工单',
      },
      component: () => import('./create/index.vue'),
    },
    {
      path: 'order-info/:id',
      name: 'order-base',
      meta: {
        text: '工单详情',
      },
      component: () => import('./info/index.vue'),
    },
  ],
}
