import Layout from '../../layout/index.vue'

export default {
  path: '/apply',
  name: 'apply',
  component: Layout,
  redirect: '/apply/apply-base',
  meta: {
    text: '发起审批',
    active: 'approval-new',
  },
  children: [
    {
      path: 'apply-base',
      name: 'apply-base',
      meta: {
        text: '常用审批',
        active: 'approval-new',
      },
      component: () => import('./base/index.vue'),
    },
    {
      path: 'apply-sign',
      name: 'apply-sign',
      meta: {
        text: '补签',
        active: 'approval-new',
      },
      component: () => import('./sign/index.vue'),
    },
  ],
}
