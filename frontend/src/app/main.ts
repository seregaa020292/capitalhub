import { createApp } from 'vue'
import { injectable } from 'inversify'
import App from '@/app/App.vue'
import store from '@/app/store'
import router from '@/app/router'
import { initPlugins } from '@/app/plugins'
import { initDirectives } from '@/app/directives'
import { elementRegister } from '@/app/themes'
import registerServiceWorker from '@/app/registerServiceWorker'
import { AuthCheckUseCaseDI } from '@/domain/auth/module/di'
import { config } from '@/data/config/app'

@injectable()
export default class Main {
  constructor() {
    this.init()
  }

  async init() {
    const app: ReturnType<typeof createApp> = createApp(App)

    /**
     ******************************
     * If auth, then check logged
     ******************************
     */
    await AuthCheckUseCaseDI().execute()

    /**
     ******************************
     * Mock data
     ******************************
     */
    if (config.isMock) {
      import('@/data/mock').then((m) => m.default.init())
    }

    /**
     ******************************
     * Plugins
     ******************************
     */
    initPlugins(app)

    /**
     ******************************
     * Directives
     ******************************
     */
    initDirectives(app)

    /**
     ******************************
     * UI framework
     ******************************
     */
    elementRegister(app)

    /**
     ******************************
     * Vuex
     ******************************
     */
    app.use(store)

    /**
     ******************************
     * Router
     ******************************
     */
    app.use(router)

    /**
     ******************************
     * Register ServiceWorker
     ******************************
     */
    registerServiceWorker()

    app.mount('#app')
  }
}
