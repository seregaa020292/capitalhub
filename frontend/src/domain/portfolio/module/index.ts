import { ContainerModule } from 'inversify'
import { types } from '@/domain/portfolio/module/types'
import {
  IPortfolioClientApi,
  PortfolioClientApi,
} from '@/domain/portfolio/clients/api/PortfolioClientApi'
import {
  IPortfolioFetchUseCase,
  PortfolioFetchUseCase,
} from '@/domain/portfolio/interactors/PortfolioFetchUseCase'
import {
  IPortfoliosFetchUseCase,
  PortfoliosFetchUseCase,
} from '@/domain/portfolio/interactors/PortfoliosFetchUseCase'
import {
  IPortfolioAddUseCase,
  PortfolioAddUseCase,
} from '@/domain/portfolio/interactors/PortfolioAddUseCase'
import {
  IPortfolioPresenter,
  PortfolioPresenter,
} from '@/domain/portfolio/presenters/PortfolioPresenter'
import {
  IPortfolioRepository,
  PortfolioRepository,
} from '@/domain/portfolio/repositories/PortfolioRepository'
import {
  IPortfolioChooseUseCase,
  PortfolioChooseUseCase,
} from '@/domain/portfolio/interactors/PortfolioChooseUseCase'

export const portfolioModule = new ContainerModule((bind) => {
  bind<IPortfolioClientApi>(types.IPortfolioClientApi).to(PortfolioClientApi).inSingletonScope()

  bind<IPortfolioFetchUseCase>(types.IPortfolioFetchUseCase).to(PortfolioFetchUseCase).inSingletonScope()
  bind<IPortfoliosFetchUseCase>(types.IPortfoliosFetchUseCase).to(PortfoliosFetchUseCase).inSingletonScope()
  bind<IPortfolioAddUseCase>(types.IPortfolioAddUseCase).to(PortfolioAddUseCase).inSingletonScope()
  bind<IPortfolioChooseUseCase>(types.IPortfolioChooseUseCase).to(PortfolioChooseUseCase).inSingletonScope()

  bind<IPortfolioPresenter>(types.IPortfolioPresenter).to(PortfolioPresenter).inSingletonScope()

  bind<IPortfolioRepository>(types.IPortfolioRepository).to(PortfolioRepository).inSingletonScope()
})
