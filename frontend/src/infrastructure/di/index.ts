import { Container } from 'inversify'
import getDecorators from 'inversify-inject-decorators'
import types from '@/infrastructure/di/types'

import Main from '@/app/main'
import store, { StoreRoot } from '@/app/store'
import { IUserRepository, UserRepository } from '@/domain/user/repositories/UserRepository'
import { IUserClientApi, UserClientApi } from '@/services/api/UserClientApi'
import { AuthClientApi, IAuthClientApi } from '@/services/api/AuthClientApi'
import { IStorageService, StorageService } from '@/services/auth/StorageService'
import { AssetClientApi, IAssetClientApi } from '@/services/api/AssetClientApi'
import { AssetRepository, IAssetRepository } from '@/domain/asset/repositories/AssetRepository'
import { AuthService, IAuthService } from '@/services/auth/AuthService'
import { ErrorHandler, IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { IMessageService, MessageService } from '@/services/message/MessageService'
import { ILoggerService, LoggerService } from '@/services/logger/LoggerService'
import { INotifyService, NotifyService } from '@/services/notify/NotifyService'
import { AuthRepository, IAuthRepository } from '@/services/auth/AuthRepository'
import { IRouterService, RouterService } from '@/services/router/RouterService'
import { AuthPresenter, IAuthPresenter } from '@/services/auth/AuthPresenter'
import { IMarketClientApi, MarketClientApi } from '@/services/api/MarketClientApi'
import { IMarketRepository, MarketRepository } from '@/domain/market/repositories/MarketRepository'
import { AssetPresenter, IAssetPresenter } from '@/domain/asset/presenter/AssetPresenter'
import { AssetFetchUseCase, IAssetFetchUseCase } from '@/domain/asset/interactors/AssetFetchUseCase'
import { UserPresenter, IUserPresenter } from '@/domain/user/presenter/UserPresenter'
import { IUserFetchUseCase, UserFetchUseCase } from '@/domain/user/interactors/UserFetchUseCase'
import {
  IMarketSearchUseCase,
  MarketSearchUseCase,
} from '@/domain/market/interactors/MarketSearchUseCase'
import { AssetAddUseCase, IAssetAddUseCase } from '@/domain/asset/interactors/AssetAddUseCase'
import { IPortfolioClientApi, PortfolioClientApi } from '@/services/api/PortfolioClientApi'
import {
  IPortfolioRepository,
  PortfolioRepository,
} from '@/domain/portfolio/repositories/PortfolioRepository'
import {
  IPortfolioFetchUseCase,
  PortfolioFetchUseCase,
} from '@/domain/portfolio/interactors/PortfolioFetchUseCase'
import {
  IPortfolioPresenter,
  PortfolioPresenter,
} from '@/domain/portfolio/presenter/PortfolioPresenter'

const DIContainer = new Container()

DIContainer.bind<Main>(Main).toSelf()
DIContainer.bind<StoreRoot>(types.IStoreRoot).toConstantValue(store)

DIContainer.bind<IUserClientApi>(types.IUserClientApi).to(UserClientApi).inSingletonScope()
DIContainer.bind<IAuthClientApi>(types.IAuthClientApi).to(AuthClientApi).inSingletonScope()
DIContainer.bind<IAssetClientApi>(types.IAssetClientApi).to(AssetClientApi).inSingletonScope()
DIContainer.bind<IMarketClientApi>(types.IMarketClientApi).to(MarketClientApi).inSingletonScope()
DIContainer.bind<IPortfolioClientApi>(types.IPortfolioClientApi).to(PortfolioClientApi).inSingletonScope()

DIContainer.bind<IAssetFetchUseCase>(types.IAssetFetchUseCase).to(AssetFetchUseCase).inSingletonScope()
DIContainer.bind<IAssetAddUseCase>(types.IAssetAddUseCase).to(AssetAddUseCase).inSingletonScope()
DIContainer.bind<IUserFetchUseCase>(types.IUserFetchUseCase).to(UserFetchUseCase).inSingletonScope()
DIContainer.bind<IMarketSearchUseCase>(types.IMarketSearchUseCase).to(MarketSearchUseCase).inSingletonScope()
DIContainer.bind<IPortfolioFetchUseCase>(types.IPortfolioFetchUseCase).to(PortfolioFetchUseCase).inSingletonScope()

DIContainer.bind<IAuthPresenter>(types.IAuthPresenter).to(AuthPresenter).inSingletonScope()
DIContainer.bind<IUserPresenter>(types.IUserPresenter).to(UserPresenter).inSingletonScope()
DIContainer.bind<IAssetPresenter>(types.IAssetPresenter).to(AssetPresenter).inSingletonScope()
DIContainer.bind<IPortfolioPresenter>(types.IPortfolioPresenter).to(PortfolioPresenter).inSingletonScope()

DIContainer.bind<IUserRepository>(types.IUserRepository).to(UserRepository).inSingletonScope()
DIContainer.bind<IAuthRepository>(types.IAuthRepository).to(AuthRepository).inSingletonScope()
DIContainer.bind<IAssetRepository>(types.IAssetRepository).to(AssetRepository).inSingletonScope()
DIContainer.bind<IMarketRepository>(types.IMarketRepository).to(MarketRepository).inSingletonScope()
DIContainer.bind<IPortfolioRepository>(types.IPortfolioRepository).to(PortfolioRepository).inSingletonScope()

DIContainer.bind<IAuthService>(types.IAuthService).to(AuthService).inSingletonScope()
DIContainer.bind<IStorageService>(types.IStorageService).to(StorageService).inSingletonScope()
DIContainer.bind<IMessageService>(types.IMessageService).to(MessageService).inSingletonScope()
DIContainer.bind<INotifyService>(types.INotifyService).to(NotifyService).inSingletonScope()
DIContainer.bind<ILoggerService>(types.ILoggerService).to(LoggerService).inSingletonScope()
DIContainer.bind<IRouterService>(types.IRouterService).to(RouterService).inSingletonScope()

DIContainer.bind<IErrorHandler>(types.IErrorHandler).to(ErrorHandler).inSingletonScope()

const { lazyInject } = getDecorators(DIContainer)
export { lazyInject, DIContainer }
