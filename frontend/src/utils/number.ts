
interface IFormatter {
  rub: Intl.NumberFormat
  usd: Intl.NumberFormat
  eur: Intl.NumberFormat
  cny: Intl.NumberFormat
}

const numberFormat = (currency: string) => new Intl.NumberFormat(undefined, {
  style: 'currency',
  currency,

  // Эти параметры необходимо округлить до целых чисел, если вы этого хотите.
  //minimumFractionDigits: 0,
  maximumFractionDigits: 4,
})

export const currencyFormatter: IFormatter = {
  rub: numberFormat('RUB'),
  usd: numberFormat('USD'),
  eur: numberFormat('EUR'),
  cny: numberFormat('CNY'),
}
