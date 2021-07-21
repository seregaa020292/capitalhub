import { createRouter, createWebHistory } from 'vue-router'
import { config } from '@/data/config/app'
import routes from './routes'
import { loggedInMiddleware } from '@/app/router/middlewares'

const router = createRouter({
  history: createWebHistory(config.baseUrl),
  routes,
})

router.beforeEach(loggedInMiddleware)

export default router
