import { injectable } from 'inversify'
import { http } from '@/infrastructure/network/http'
import { parsePatternUrl, urls } from '@/infrastructure/network/urls'
import {
  IPortfolioChangeFields,
  IPortfolioStats,
  IPortfolioTotal,
} from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioClientApi {
  fetchActiveTotal(): Promise<IPortfolioTotal>
  fetchAllStats(): Promise<IPortfolioStats[]>
  add(portfolio: IPortfolioChangeFields): Promise<IPortfolioStats>
  edit(portfolioId: string, portfolio: IPortfolioChangeFields): Promise<IPortfolioStats>
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

  add(portfolio: IPortfolioChangeFields): Promise<IPortfolioStats> {
    return http.post(urls.api_v1.PORTFOLIO_ADD, portfolio)
  }

  edit(portfolioId: string, portfolio: IPortfolioChangeFields): Promise<IPortfolioStats> {
    return http.put(parsePatternUrl(urls.api_v1.PORTFOLIO_EDIT, portfolioId), portfolio)
  }

  choose(portfolioId: string): Promise<boolean> {
    return http.put(parsePatternUrl(urls.api_v1.PORTFOLIO_CHOOSE, portfolioId));
  }
}
