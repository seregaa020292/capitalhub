import { onMounted, onUnmounted, reactive } from 'vue'
import CurrencyClientSocket, { CurrencyTicker } from '@/services/socket/CurrencyClientSocket'

interface ICurrencyBtcUsd {
  state: {
    currencyBtcUsd: CurrencyTicker
  }
}

export const useCurrencyBtc = (): ICurrencyBtcUsd => {
  const currencyClientSocket = new CurrencyClientSocket()

  const state = reactive({
    currencyBtcUsd: {
      currentPrice: 0,
      prevPrice: 0,
    },
  })

  onMounted(() => {
    currencyClientSocket.tickerSubscribe('BTCUSD', (response: CurrencyTicker) => {
      state.currencyBtcUsd = response
    })
  })

  onUnmounted(() => {
    currencyClientSocket.tickerTerminate()
  })

  return {
    state,
  }
}
