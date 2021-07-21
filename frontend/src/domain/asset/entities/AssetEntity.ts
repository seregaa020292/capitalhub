export interface IAsset {
  ticker: string
  identify: string
  title: string
  imageUrl?: string
  marketId: string
  firstNotationAt: string
  totalAmount: number
  totalQuantity: number
  totalCount: number
  averagePurchasePrice: number
  currentPrice: number
  // currentValue: number
  // changeTotalPercent: number
  // changePerDayPercent: number
}

export interface IAssetPrice {
  identify: string
  currentPrice: number
}

export interface IAssetNotation {
  marketId: string
  notationAt: string
  amount: number
  quantity: number
}
