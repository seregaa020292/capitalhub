import { DIContainer } from '@/infrastructure/di'
import { types } from '@/domain/portfolio/module/types'
import { IPortfolioPresenter } from '@/domain/portfolio/presenters/PortfolioPresenter'
import { IPortfolioFetchUseCase } from '@/domain/portfolio/interactors/PortfolioFetchUseCase'
import { IPortfoliosFetchUseCase } from '@/domain/portfolio/interactors/PortfoliosFetchUseCase'
import { IPortfolioAddUseCase } from '@/domain/portfolio/interactors/PortfolioAddUseCase'

export const PortfolioPresenterDI = (): IPortfolioPresenter => DIContainer.get(types.IPortfolioPresenter)
export const PortfolioFetchUseCaseDI = (): IPortfolioFetchUseCase => DIContainer.get(types.IPortfolioFetchUseCase)
export const PortfoliosFetchUseCaseDI = (): IPortfoliosFetchUseCase => DIContainer.get(types.IPortfoliosFetchUseCase)
export const PortfolioAddUseCaseDI = (): IPortfolioAddUseCase => DIContainer.get(types.IPortfolioAddUseCase)
