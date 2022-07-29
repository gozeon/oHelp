/**
 * 表格header添加提示
 * @link https://blog.csdn.net/m0_60540878/article/details/120368512
 * @param {*} h
 * @param {*} label
 * @param {*} content
 * @param {*} placement
 * @returns
 */
export const renderHeaderWithToolTip = (h, label, content, placement = 'top') => [
  label,
  h(
    'el-tooltip',
    {
      props: { content, placement },
    },
    [
      h('span', {
        class: {
          'el-icon-question': true,
        },
      }),
    ]
  ),
]
