import { Container } from 'inversify'
import getDecorators from 'inversify-inject-decorators'
import { baseTypes } from '@/infrastructure/di/types'

import Main from '@/app/main'
import store, { StoreRoot } from '@/app/store'
import { ErrorHandler, IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { IMessageService, MessageService } from '@/services/message/MessageService'
import { ILoggerService, LoggerService } from '@/services/logger/LoggerService'
import { INotifyService, NotifyService } from '@/services/notify/NotifyService'
import { IRouterService, RouterService } from '@/services/router/RouterService'
import { userModule } from '@/domain/user/module'
import { portfolioModule } from '@/domain/portfolio/module'
import { marketModule } from '@/domain/market/module'
import { assetModule } from '@/domain/asset/module'
import { authModule } from '@/domain/auth/module'
import { applicationModule } from '@/domain/application/module'

const DIContainer = new Container()

DIContainer.bind<Main>(Main).toSelf()

DIContainer.bind<StoreRoot>(baseTypes.IStoreRoot).toConstantValue(store)

DIContainer.bind<IMessageService>(baseTypes.IMessageService).to(MessageService).inSingletonScope()
DIContainer.bind<INotifyService>(baseTypes.INotifyService).to(NotifyService).inSingletonScope()
DIContainer.bind<ILoggerService>(baseTypes.ILoggerService).to(LoggerService).inSingletonScope()
DIContainer.bind<IRouterService>(baseTypes.IRouterService).to(RouterService).inSingletonScope()

DIContainer.bind<IErrorHandler>(baseTypes.IErrorHandler).to(ErrorHandler).inSingletonScope()

DIContainer.load(userModule, portfolioModule, marketModule, assetModule, authModule, applicationModule)

const { lazyInject } = getDecorators(DIContainer)
export { lazyInject, DIContainer }
