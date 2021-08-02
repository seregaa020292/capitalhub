import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/portfolio/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IPortfolioRepository } from '@/domain/portfolio/repositories/PortfolioRepository'
import { IPortfolioClientApi } from '@/domain/portfolio/clients/api/PortfolioClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'

export interface IPortfoliosFetchUseCase extends BaseUseCase<unknown, Promise<void>> {}

@injectable()
export class PortfoliosFetchUseCase implements IPortfoliosFetchUseCase {
  @inject(types.IPortfolioClientApi)
  private portfolioClient!: IPortfolioClientApi

  @inject(types.IPortfolioRepository)
  private portfolioRepository!: IPortfolioRepository

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute() {
    try {
      this.portfolioRepository.loadingPortfolios(true)
      const response = await this.portfolioClient.fetchAllStats()

      this.portfolioRepository.setPortfolios(response)
    } catch (error) {
      this.errorHandler.handle(error)
    } finally {
      this.portfolioRepository.loadingPortfolios(false)
    }
  }
}
