export interface IMarket {
  title: string
  ticker: string
  content?: string
  marketId: string
  titleInstrument: string
  descInstrument: string
}

export interface IMarketSearch {
  hasMore: boolean
  markets: IMarket[]
  page: number
  size: number
  totalCount: number
  totalPages: number
}

export interface IMarketGroup {
  [propKey: string]: IMarketOption
}

export interface IMarketOption {
  label: string
  description: string
  options: IMarket[]
}
