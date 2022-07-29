import './scss/style.scss'

import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import ElementUI from 'element-ui'
Vue.use(ElementUI)

import axios from './axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios, axios)

import * as utils from './utils.js'
Vue.prototype.$utils = utils

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app')
