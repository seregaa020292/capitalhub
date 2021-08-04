import { createStore, createLogger, Store, ModuleTree, Module } from 'vuex'
import { config } from '@/data/config/app'
import { IUserState } from '@/app/store/modules/user'
import { IAssetState } from '@/app/store/modules/asset'
import { IAuthState } from '@/app/store/modules/auth'
import { IMarketState } from '@/app/store/modules/market'
import { IPortfolioState } from '@/app/store/modules/portfolio'
import { IApplicationState } from '@/app/store/modules/application'

export interface IStateRoot {
  user: IUserState
  auth: IAuthState
  asset: IAssetState
  market: IMarketState
  portfolio: IPortfolioState
  application: IApplicationState
}

const storeRoot: Store<IStateRoot> = createStore({
  strict: config.isProd,
  modules: { ...importModules() },
  plugins: config.isDev ? [createLogger()] : [],
})

export type StoreRoot = typeof storeRoot

type ModuleRoot = ModuleTree<IStateRoot>

function importModules(): ModuleRoot {
  const files = import.meta.globEager('./modules/*.ts')

  const modules: ModuleRoot = {}

  for (const fileName in files) {
    if (!Object.prototype.hasOwnProperty.call(files, fileName)) continue

    const path: string = fileName.replace(/(\.\/|\.ts)/g, '')
    const [, namespace]: Array<string> = path.split('/')
    const module: Module<IStateRoot, IStateRoot> = files[fileName]

    modules[namespace] = {
      namespaced: false,
      ...module,
    }
  }

  return modules
}

export default storeRoot
