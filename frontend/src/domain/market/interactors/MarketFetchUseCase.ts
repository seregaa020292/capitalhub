import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import types from '@/infrastructure/di/types'
import { IMarketRepository } from '@/domain/market/repositories/MarketRepository'
import { IMarket } from '@/domain/market/entities/MarketEntity'
import { IMarketClientApi } from '@/services/api/MarketClientApi'

@injectable()
export class MarketFetchUseCase implements BaseUseCase<void, Promise<void>> {
  @inject(types.IMarketClientApi)
  private marketClient!: IMarketClientApi

  @inject(types.IMarketRepository)
  private marketRepository!: IMarketRepository

  async execute(): Promise<void> {
    const response: IMarket[] = await this.marketClient.fetchMarkets()
    this.marketRepository.saveMarkets(response)
  }
}
