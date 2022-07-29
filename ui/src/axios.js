import axios from 'axios'
import { Message } from 'element-ui'

axios.defaults.baseURL = process.env.VUE_APP_API

axios.interceptors.response.use(
  (response) => {
    const { data, msg, status } = response?.data

    if (+status === 403) {
      Message.error('登录信息失效, 请重新登录')
      location.href = data?.url
      return
    }
    return response
  },
  (error) => {
    return Promise.reject(error)
  }
)

export default axios
