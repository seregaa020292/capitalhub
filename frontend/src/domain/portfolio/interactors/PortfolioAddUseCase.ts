import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/portfolio/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IPortfolioRepository } from '@/domain/portfolio/repositories/PortfolioRepository'
import { IPortfolioClientApi } from '@/domain/portfolio/clients/api/PortfolioClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { IPortfolioAdd } from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioAddUseCase extends BaseUseCase<IPortfolioAdd, Promise<void>> {}

@injectable()
export class PortfolioAddUseCase implements IPortfolioAddUseCase {
  @inject(types.IPortfolioClientApi)
  private portfolioClient!: IPortfolioClientApi

  @inject(types.IPortfolioRepository)
  private portfolioRepository!: IPortfolioRepository

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(portfolio: IPortfolioAdd) {
    try {
      const response = await this.portfolioClient.add(portfolio)

      this.portfolioRepository.addPortfolio(response)
    } catch (error) {
      this.errorHandler.handle(error)
    }
  }
}
