import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/portfolio/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IPortfolioClientApi } from '@/domain/portfolio/clients/api/PortfolioClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { IPortfolioRepository } from '@/domain/portfolio/repositories/PortfolioRepository'

export interface IPortfolioRemoveUseCase extends BaseUseCase<string, Promise<boolean>> {}

@injectable()
export class PortfolioRemoveUseCase implements IPortfolioRemoveUseCase {
  @inject(types.IPortfolioClientApi)
  private portfolioClient!: IPortfolioClientApi

  @inject(types.IPortfolioRepository)
  private portfolioRepository!: IPortfolioRepository

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(portfolioId: string) {
    try {
      const response = await this.portfolioClient.remove(portfolioId)

      this.portfolioRepository.remove(portfolioId)
      return response
    } catch (error) {
      this.errorHandler.handle(error).report()
    }

    return false
  }
}
