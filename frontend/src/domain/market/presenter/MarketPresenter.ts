import { inject, injectable } from 'inversify'
import types from '@/infrastructure/di/types'
import { IMarketRepository } from '@/domain/market/repositories/MarketRepository'
import { IMarket } from '@/domain/market/entities/MarketEntity'

export interface IMarketPresenter {
  markets(): IMarket[]
}

@injectable()
export class MarketPresenter implements IMarketPresenter {
  @inject(types.IMarketRepository)
  private marketRepository!: IMarketRepository

  markets(): IMarket[] {
    return this.marketRepository.getMarkets()
  }
}
