module.exports = {
  namespaced: true,
  state: {
    collapse: false,
    options: {
      approvalStatusOptions: [
        { label: '审批中', value: 1 },
        { label: '通过', value: 2 },
        { label: '驳回', value: 3 },
      ],
      approvalNewOptions: [
        { label: '常用审批', link: '/apply/apply-base' },
        { label: '补签', link: '/apply/apply-sign' },
      ],
      applyOperateOptions: [
        { label: '申请', value: 1 },
        { label: '移交', value: 2 },
        { label: '通过', value: 3 },
        { label: '驳回', value: 4 },
        { label: '回复', value: 5 },
      ],
      users: [
        { value: 'admin', label: '管理员' },
        { value: 'admin_vas', label: '管理员2' },
      ],
    },
  },
  mutations: {
    toggleSideBarCollapse(state) {
      state.collapse = !state.collapse
    },
  },
}
