import { injectable } from 'inversify'
import { http } from '@/infrastructure/network/http'
import urls from '@/infrastructure/network/urls'
import { IPortfolioTotal } from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioClientApi {
  fetchActiveTotal(): Promise<IPortfolioTotal>
}

@injectable()
export class PortfolioClientApi implements IPortfolioClientApi {
  fetchActiveTotal(): Promise<IPortfolioTotal> {
    return http.get(urls.api_v1.PORTFOLIO_ACTIVE_TOTAL)
  }
}
