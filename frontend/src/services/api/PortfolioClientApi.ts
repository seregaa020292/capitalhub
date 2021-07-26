import { injectable } from 'inversify'
import { http } from '@/infrastructure/network/http'
import urls from '@/infrastructure/network/urls'
import {
  IPortfolioAdd,
  IPortfolioStats,
  IPortfolioTotal,
} from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioClientApi {
  fetchActiveTotal(): Promise<IPortfolioTotal>
  fetchAllStats(): Promise<IPortfolioStats[]>
  add(portfolio: IPortfolioAdd): Promise<IPortfolioStats>
}

@injectable()
export class PortfolioClientApi implements IPortfolioClientApi {
  fetchActiveTotal(): Promise<IPortfolioTotal> {
    return http.get(urls.api_v1.PORTFOLIO_ACTIVE_TOTAL)
  }

  fetchAllStats(): Promise<IPortfolioStats[]> {
    return http.get(urls.api_v1.PORTFOLIO_ALL_STATS)
  }

  add(portfolio: IPortfolioAdd): Promise<IPortfolioStats> {
    return http.post(urls.api_v1.PORTFOLIO_ADD, portfolio)
  }
}
