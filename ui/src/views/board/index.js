import Layout from '../../layout/index.vue'

export default {
  path: '/board',
  name: 'board',
  component: Layout,
  redirect: '/board/system-overview',
  meta: {
    text: '看板',
  },
  children: [
    {
      path: 'system-overview',
      name: 'system-overview',
      component: () => import('./system-overview/index.vue'),
    },
  ],
}
