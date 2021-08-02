import { inject, injectable } from 'inversify'
import { types } from '@/domain/portfolio/module/types'
import { IPortfolioRepository } from '@/domain/portfolio/repositories/PortfolioRepository'
import { IPortfolio, IPortfolioStats } from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioPresenter {
  portfolio(): IPortfolio
  portfolios(): IPortfolioStats[]
  loadingPortfolios(): boolean
}

@injectable()
export class PortfolioPresenter implements IPortfolioPresenter {
  @inject(types.IPortfolioRepository)
  private portfolioRepository!: IPortfolioRepository

  portfolio(): IPortfolio {
    return this.portfolioRepository.getPortfolio()
  }

  portfolios(): IPortfolioStats[] {
    return this.portfolioRepository.getPortfolios()
  }

  loadingPortfolios(): boolean {
    return this.portfolioRepository.getLoadingPortfolios()
  }
}
