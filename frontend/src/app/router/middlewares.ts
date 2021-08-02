import { AuthPresenterDI } from '@/domain/auth/module/di'
import { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'

export const loggedInMiddleware = (
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext,
): void => {
  if (to.matched.some((record) => record.meta.auth)) {
    if (AuthPresenterDI().loggedIn()) {
      next()
      return
    }
    next({ name: 'login' })
    return
  }
  next()
}
