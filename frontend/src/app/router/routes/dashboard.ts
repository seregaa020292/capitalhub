import { RouteRecordRaw } from 'vue-router'
import collects from '@/app/view/layout/collects'

const DashboardRouter: Array<RouteRecordRaw> = [
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('@/app/view/pages/dashboard/Main.vue'),
    children: [
      {
        path: 'asset-edit',
        components: {
          AssetEdit: () => import('@/app/view/pages/dashboard/AssetEdit.vue'),
        },
      },
    ],
  },
  {
    path: '/dashboard/:ticker',
    component: () => import('@/app/view/pages/dashboard/Ticker.vue'),
  },
]

DashboardRouter.forEach((route) => {
  route.meta = {
    ...route.meta,
    layout: collects.dashboardLayout,
    auth: true,
  }
})

export default DashboardRouter
