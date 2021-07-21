import DefaultLayout from './DefaultLayout.vue'
import DashboardLayout from './DashboardLayout.vue'
import AuthLayout from './AuthLayout.vue'

const collects = {
  defaultLayout: DefaultLayout,
  dashboardLayout: DashboardLayout,
  authLayout: AuthLayout,
}

export default collects as typeof collects
