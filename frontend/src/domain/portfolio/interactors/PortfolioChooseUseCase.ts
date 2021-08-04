import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/portfolio/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IPortfolioClientApi } from '@/domain/portfolio/clients/api/PortfolioClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'

export interface IPortfolioChooseUseCase extends BaseUseCase<string, Promise<boolean>> {}

@injectable()
export class PortfolioChooseUseCase implements IPortfolioChooseUseCase {
  @inject(types.IPortfolioClientApi)
  private portfolioClient!: IPortfolioClientApi

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(portfolioId: string) {
    try {
      return await this.portfolioClient.choose(portfolioId)
    } catch (error) {
      this.errorHandler.handle(error).report()
    }

    return false
  }
}
