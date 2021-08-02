import { injectable } from 'inversify'
import { http } from '@/infrastructure/network/http'
import urls from '@/infrastructure/network/urls'
import { IMarket, IMarketSearch } from '@/domain/market/entities/MarketEntity'

export interface IMarketClientApi {
  fetchMarkets(): Promise<IMarket[]>
  searchMarkets(title: string): Promise<IMarketSearch>
}

@injectable()
export class MarketClientApi implements IMarketClientApi {
  fetchMarkets(): Promise<IMarket[]> {
    return http.get(urls.api_v1.MARKETS)
  }

  searchMarkets(title: string): Promise<IMarketSearch> {
    return http.get(urls.api_v1.SEARCH_MARKETS, { params: { title } })
  }
}
