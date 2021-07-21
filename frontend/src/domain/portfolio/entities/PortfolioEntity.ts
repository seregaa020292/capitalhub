import { IAsset } from '@/domain/asset/entities/AssetEntity'

export interface IPortfolio {
  active: boolean,
  createdAt: string
  currencyId: string
  portfolioId: string
  title: string
  updatedAt: string
  userId: string
}

export interface IPortfolioTotal {
  portfolio: IPortfolio
  assetTotal: IAsset[]
}
