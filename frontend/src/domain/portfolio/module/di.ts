import { DIContainer } from '@/infrastructure/di'
import { types } from '@/domain/portfolio/module/types'
import { IPortfolioPresenter } from '@/domain/portfolio/presenters/PortfolioPresenter'
import { IPortfolioFetchUseCase } from '@/domain/portfolio/interactors/PortfolioFetchUseCase'
import { IPortfoliosFetchUseCase } from '@/domain/portfolio/interactors/PortfoliosFetchUseCase'
import { IPortfolioAddUseCase } from '@/domain/portfolio/interactors/PortfolioAddUseCase'
import { IPortfolioEditUseCase } from '@/domain/portfolio/interactors/PortfolioEditUseCase'
import { IPortfolioChooseUseCase } from '@/domain/portfolio/interactors/PortfolioChooseUseCase'

export const PortfolioPresenterDI = (): IPortfolioPresenter => DIContainer.get(types.IPortfolioPresenter)
export const PortfolioFetchUseCaseDI = (): IPortfolioFetchUseCase => DIContainer.get(types.IPortfolioFetchUseCase)
export const PortfoliosFetchUseCaseDI = (): IPortfoliosFetchUseCase => DIContainer.get(types.IPortfoliosFetchUseCase)
export const PortfolioChooseUseCaseDI = (): IPortfolioChooseUseCase => DIContainer.get(types.IPortfolioChooseUseCase)
export const PortfolioAddUseCaseDI = (): IPortfolioAddUseCase => DIContainer.get(types.IPortfolioAddUseCase)
export const PortfolioEditUseCaseDI = (): IPortfolioEditUseCase => DIContainer.get(types.IPortfolioEditUseCase)
