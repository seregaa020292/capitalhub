import { inject, injectable } from 'inversify'
import { StoreRoot } from '@/app/store'
import { baseTypes } from '@/infrastructure/di/types'
import {
  AddPortfolio,
  FetchPortfolio,
  FetchPortfolios,
  LoadingPortfolios,
} from '@/app/store/modules/portfolio'
import { IPortfolio, IPortfolioStats } from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioRepository {
  getPortfolio(): IPortfolio
  getPortfolios(): IPortfolioStats[]
  getLoadingPortfolios(): boolean
  setPortfolio(portfolio: IPortfolio): void
  setPortfolios(portfolios: IPortfolioStats[]): void
  loadingPortfolios(status: boolean): void
  addPortfolio(portfolio: IPortfolioStats): void
}

@injectable()
export class PortfolioRepository implements IPortfolioRepository {
  constructor(@inject(baseTypes.IStoreRoot) private store: StoreRoot) {}

  getPortfolio(): IPortfolio {
    return this.store.state.portfolio.active
  }

  getPortfolios(): IPortfolioStats[] {
    return this.store.state.portfolio.all
  }

  getLoadingPortfolios(): boolean {
    return this.store.state.portfolio.loadingAll
  }

  setPortfolio(portfolio: IPortfolio): void {
    this.store.commit(new FetchPortfolio(portfolio))
  }

  loadingPortfolios(status: boolean): void {
    this.store.commit(new LoadingPortfolios(status))
  }

  setPortfolios(portfolios: IPortfolioStats[]): void {
    this.store.commit(new FetchPortfolios(portfolios))
  }

  addPortfolio(portfolio: IPortfolioStats): void {
    this.store.commit(new AddPortfolio(portfolio))
  }
}
