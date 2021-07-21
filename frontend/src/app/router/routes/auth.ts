import { RouteRecordRaw } from 'vue-router'
import collects from '@/app/view/layout/collects'

const AuthRouter: Array<RouteRecordRaw> = [
  {
    path: '/login',
    name: 'login',
    component: () => import('@/app/view/pages/auth/Login.vue'),
    meta: {
      layout: collects.authLayout,
    },
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('@/app/view/pages/auth/Register.vue'),
    meta: {
      layout: collects.authLayout,
    },
  },
  {
    path: '/confirmed/:code',
    name: 'confirmed-code',
    component: () => import('@/app/view/pages/auth/Confirmed.vue'),
    meta: {
      layout: collects.authLayout,
    },
  }
]

export default AuthRouter
