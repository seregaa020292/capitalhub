import AxiosMockAdapter from 'axios-mock-adapter'
import { urls } from '@/infrastructure/network/urls'
import { serverResponse } from '@/utils/server'

export default (mock: AxiosMockAdapter): void => {
  mock.onGet(urls.api_v1.TOTAL_ASSETS).reply(
    200,
    serverResponse({
      message: [
        {
          title: 'Ripple',
          ticker: 'XRP',
          totalQuantity: 2004,
          totalAmount: 69579.59,
          firstNotationAt: 4,
          currentPrice: 74.95,
          currentValue: 150206.23,
          averagePurchasePrice: 34.7,
          changeTotalPercent: 115.88,
          changePerDayPercent: 14.78,
        },
        {
          title: 'Stellar',
          ticker: 'XLM',
          totalQuantity: 104,
          totalAmount: 2579.59,
          firstNotationAt: 1,
          currentPrice: 34.54,
          currentValue: 10206.23,
          averagePurchasePrice: 12.7,
          changeTotalPercent: 65.88,
          changePerDayPercent: 12.8,
        },
      ],
    })
  )
}
