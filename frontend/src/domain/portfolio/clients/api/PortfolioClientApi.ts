import { injectable } from 'inversify'
import { http } from '@/infrastructure/network/http'
import { parsePatternUrl, urls } from '@/infrastructure/network/urls'
import {
  IPortfolioAdd,
  IPortfolioStats,
  IPortfolioTotal,
} from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioClientApi {
  fetchActiveTotal(): Promise<IPortfolioTotal>
  fetchAllStats(): Promise<IPortfolioStats[]>
  add(portfolio: IPortfolioAdd): Promise<IPortfolioStats>
  choose(portfolioId: string): Promise<boolean>
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

  choose(portfolioId: string): Promise<boolean> {
    return http.put(parsePatternUrl(urls.api_v1.PORTFOLIO_CHOOSE, portfolioId));
  }
}
