export interface ICurrency {
  currencyId: string
  title: string
  description: string
  createdAt: string
  updatedAt: string
}

export interface IDashboard {
  currencies: ICurrency[]
}
