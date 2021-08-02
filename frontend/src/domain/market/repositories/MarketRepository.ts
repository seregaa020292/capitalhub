import { inject, injectable } from 'inversify'
import { StoreRoot } from '@/app/store'
import { baseTypes } from '@/infrastructure/di/types'
import { IMarket } from '@/domain/market/entities/MarketEntity'
import { FetchMarkets, ClearMarkets } from '@/app/store/modules/market'

export interface IMarketRepository {
  getMarkets(): IMarket[]
  saveMarkets(markets: IMarket[]): void
  clearMarkets(): void
}

@injectable()
export class MarketRepository implements IMarketRepository {
  constructor(@inject(baseTypes.IStoreRoot) private store: StoreRoot) {}

  getMarkets(): IMarket[] {
    return this.store.state.market.markets
  }

  saveMarkets(markets: IMarket[]): void {
    this.store.commit(new FetchMarkets(markets))
  }

  clearMarkets(): void {
    this.store.commit(new ClearMarkets())
  }
}
