export interface IMarket {
  title: string
  ticker: string
  content?: string
  market_id: string
  title_instrument: string
  desc_instrument: string
}

export interface IMarketSearch {
  has_more: boolean
  markets: IMarket[]
  page: number
  size: number
  total_count: number
  total_pages: number
}

export interface IMarketGroup {
  [propKey: string]: IMarketOption
}

export interface IMarketOption {
  label: string
  description: string
  options: IMarket[]
}
