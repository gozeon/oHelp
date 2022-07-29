import Vue from 'vue'
import VueRouter from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

Vue.use(VueRouter)

const routes = [
  {
    path: '*',
    redirect: '/board/system-overview',
  },
]

// dynamic init routers with views
const requireRouter = require.context('../views', true, /\.js$/)

for (const key of requireRouter.keys()) {
  routes.push(requireRouter(key).default)
}

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes,
})

// bar
router.beforeEach((to, from, next) => {
  NProgress.start()
  next()
})

router.afterEach(() => {
  NProgress.done()
})

export default router
