import urls from '@/infrastructure/network/urls'
import { Socket, Unsubscribe } from '@/infrastructure/network/socket'

export type CurrencyTicker = {
  currentPrice: number
  prevPrice: number
}

type Pair = 'BTCUSD' | 'BTCEUR'

export default class CurrencyClientSocket {
  private socket: Socket

  constructor() {
    // document.location.host
    this.socket = new Socket(urls.ws.BITFINEX)
    this.socket.init()
  }

  public tickerSubscribe(pair: Pair, handler: (response: CurrencyTicker) => void): Unsubscribe {
    const message = { event: 'subscribe', channel: 'ticker', pair }
    const callback = (message: Array<number>) => {
      if (!Array.isArray(message) || (typeof message[1] === 'string' && message[1] === 'hb')) {
        return
      }

      const valueTicker: CurrencyTicker = {
        currentPrice: message[1],
        prevPrice: message[7],
      }

      handler(valueTicker)
    }
    return this.socket.subscribe('', message, callback)
  }

  tickerTerminate(): void {
    this.socket.terminate()
  }
}
