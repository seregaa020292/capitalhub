import { injectable } from 'inversify'
import { RouteLocationRaw } from 'vue-router'
import router from '@/app/router'

export interface IRouterService {
  currentPath: string
  push(to: RouteLocationRaw): void
}

type RouteParams = RouteLocationRaw

@injectable()
export class RouterService implements IRouterService {
  public get currentPath() {
    return router.currentRoute.value.path
  }

  public push(to: RouteParams) {
    return router.push(to)
  }
}
