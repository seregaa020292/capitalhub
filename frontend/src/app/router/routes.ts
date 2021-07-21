import { RouteRecordRaw } from 'vue-router'
import Home from '@/app/view/pages/Home.vue'
import NotFound from '@/app/view/pages/NotFound.vue'

const files = import.meta.globEager('./routes/*.ts')
const includeModuleRoutes: RouteRecordRaw[] = []

for (const path in files) {
  if (!Object.prototype.hasOwnProperty.call(files, path)) continue

  const router: Array<RouteRecordRaw> = files[path].default
  includeModuleRoutes.push(...router)
}

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  ...includeModuleRoutes,
  {
    path: '/:catchAll(.*)',
    component: NotFound,
  },
]

export default routes
