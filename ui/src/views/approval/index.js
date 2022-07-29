import Layout from '../../layout/index.vue'

export default {
  path: '/approval',
  name: 'approval',
  component: Layout,
  redirect: '/approval/approval-base',
  meta: {
    text: '审批管理',
  },
  children: [
    {
      path: 'approval-base',
      name: 'approval-base',
      meta: {
        text: '审批列表',
      },
      component: () => import('./base/index.vue'),
    },
    {
      path: 'approval-my',
      name: 'approval-my',
      meta: {
        text: '我发起的',
      },
      component: () => import('./my/index.vue'),
    },
    {
      path: 'approval-assign-me',
      name: 'approval-assign-me',
      meta: {
        text: '我审批的',
      },
      component: () => import('./assign-me/index.vue'),
    },
    {
      path: 'approval-new',
      name: 'approval-new',
      meta: {
        text: '发起审批',
      },
      component: () => import('./new/index.vue'),
    },
  ],
}
