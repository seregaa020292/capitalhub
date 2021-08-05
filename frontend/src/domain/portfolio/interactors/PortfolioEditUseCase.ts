import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/portfolio/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IPortfolioRepository } from '@/domain/portfolio/repositories/PortfolioRepository'
import { IPortfolioClientApi } from '@/domain/portfolio/clients/api/PortfolioClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { IPortfolioEditable } from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioEditUseCase extends BaseUseCase<IPortfolioEditable, Promise<void>> {}

@injectable()
export class PortfolioEditUseCase implements IPortfolioEditUseCase {
  @inject(types.IPortfolioClientApi)
  private portfolioClient!: IPortfolioClientApi

  @inject(types.IPortfolioRepository)
  private portfolioRepository!: IPortfolioRepository

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute({portfolioId, ...portfolio}: IPortfolioEditable): Promise<void> {
    try {
      const response = await this.portfolioClient.edit(portfolioId, portfolio)

      this.portfolioRepository.editPortfolio(response)
    } catch (error) {
      this.errorHandler.handle(error)
    }
  }
}
