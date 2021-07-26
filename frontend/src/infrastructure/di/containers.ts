import types from '@/infrastructure/di/types'
import { DIContainer } from '@/infrastructure/di/index'
import { IAuthService } from '@/services/auth/AuthService'
import { IStorageService } from '@/services/auth/StorageService'
import { IAuthPresenter } from '@/services/auth/AuthPresenter'
import { IAssetPresenter } from '@/domain/asset/presenter/AssetPresenter'
import { IUserPresenter } from '@/domain/user/presenter/UserPresenter'
import { IMarketSearchUseCase } from '@/domain/market/interactors/MarketSearchUseCase'
import { IAssetFetchUseCase } from '@/domain/asset/interactors/AssetFetchUseCase'
import { IAssetAddUseCase } from '@/domain/asset/interactors/AssetAddUseCase'
import { IPortfolioFetchUseCase } from '@/domain/portfolio/interactors/PortfolioFetchUseCase'
import { IPortfolioPresenter } from '@/domain/portfolio/presenter/PortfolioPresenter'
import { IPortfoliosFetchUseCase } from '@/domain/portfolio/interactors/PortfoliosFetchUseCase'
import { IPortfolioAddUseCase } from '@/domain/portfolio/interactors/PortfolioAddUseCase'

/**
 * Make to: Services, Presenters, UseCases
 */
export const AuthPresenterContainer = (): IAuthPresenter => DIContainer.get(types.IAuthPresenter)
export const UserPresenterContainer = (): IUserPresenter => DIContainer.get(types.IUserPresenter)
export const AssetPresenterContainer = (): IAssetPresenter => DIContainer.get(types.IAssetPresenter)
export const PortfolioPresenterContainer = (): IPortfolioPresenter => DIContainer.get(types.IPortfolioPresenter)

export const AssetFetchUseCaseContainer = (): IAssetFetchUseCase => DIContainer.get(types.IAssetFetchUseCase)
export const AssetAddUseCaseContainer = (): IAssetAddUseCase => DIContainer.get(types.IAssetAddUseCase)
export const MarketSearchUseCaseContainer = (): IMarketSearchUseCase => DIContainer.get(types.IMarketSearchUseCase)
export const PortfolioFetchUseCaseContainer = (): IPortfolioFetchUseCase => DIContainer.get(types.IPortfolioFetchUseCase)
export const PortfoliosFetchUseCaseContainer = (): IPortfoliosFetchUseCase => DIContainer.get(types.IPortfoliosFetchUseCase)
export const PortfolioAddUseCaseContainer = (): IPortfolioAddUseCase => DIContainer.get(types.IPortfolioAddUseCase)

export const AuthServiceContainer = (): IAuthService => DIContainer.get(types.IAuthService)
export const StorageServiceContainer = (): IStorageService => DIContainer.get(types.IStorageService)
