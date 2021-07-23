import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import types from '@/infrastructure/di/types'
import { IPortfolioRepository } from '@/domain/portfolio/repositories/PortfolioRepository'
import { IPortfolioClientApi } from '@/services/api/PortfolioClientApi'
import { IPortfolioTotal } from '@/domain/portfolio/entities/PortfolioEntity'
import { IAssetRepository } from '@/domain/asset/repositories/AssetRepository'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'

export interface IPortfolioFetchUseCase extends BaseUseCase<unknown, Promise<void>> {}

@injectable()
export class PortfolioFetchUseCase implements IPortfolioFetchUseCase {
  @inject(types.IPortfolioClientApi)
  private portfolioClient!: IPortfolioClientApi

  @inject(types.IPortfolioRepository)
  private portfolioRepository!: IPortfolioRepository

  @inject(types.IAssetRepository)
  private assetRepository!: IAssetRepository

  @inject(types.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute() {
    try {
      this.assetRepository.loadingAssets(true)

      const response: IPortfolioTotal = await this.portfolioClient.fetchActiveTotal()

      this.portfolioRepository.setPortfolio(response.portfolio)
      this.assetRepository.saveAssets(response.assetTotal)
    } catch (error) {
      this.errorHandler.handle(error)
    } finally {
      this.assetRepository.loadingAssets(false)
    }
  }
}
