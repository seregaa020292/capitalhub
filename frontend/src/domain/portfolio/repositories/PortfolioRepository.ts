import { inject, injectable } from 'inversify'
import { StoreRoot } from '@/app/store'
import types from '@/infrastructure/di/types'
import { FetchPortfolio } from '@/app/store/modules/portfolio'
import { IPortfolio } from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioRepository {
  getPortfolio(): IPortfolio
  setPortfolio(portfolio: IPortfolio): void
}

@injectable()
export class PortfolioRepository implements IPortfolioRepository {
  constructor(@inject(types.IStoreRoot) private store: StoreRoot) {}

  getPortfolio(): IPortfolio {
    return this.store.state.portfolio.active
  }

  setPortfolio(portfolio: IPortfolio): void {
    this.store.commit(new FetchPortfolio(portfolio))
  }
}
