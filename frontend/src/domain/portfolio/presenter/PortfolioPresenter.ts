import { inject, injectable } from 'inversify'
import types from '@/infrastructure/di/types'
import { IPortfolioRepository } from '@/domain/portfolio/repositories/PortfolioRepository'
import { IPortfolio } from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioPresenter {
  portfolio(): IPortfolio
}

@injectable()
export class PortfolioPresenter implements IPortfolioPresenter {
  @inject(types.IPortfolioRepository)
  private portfolioRepository!: IPortfolioRepository

  portfolio(): IPortfolio {
    return this.portfolioRepository.getPortfolio()
  }
}
