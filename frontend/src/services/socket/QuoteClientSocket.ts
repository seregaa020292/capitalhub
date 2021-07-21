import { Socket, Unsubscribe } from '@/infrastructure/network/socket'
import { StorageServiceContainer } from '@/infrastructure/di/containers'
import urls from '@/infrastructure/network/urls'
import { IAssetRepository } from '@/domain/asset/repositories/AssetRepository'
import { DIContainer } from '@/infrastructure/di'
import types from '@/infrastructure/di/types'

export type QuoteData = {
  event: string // Название события
  payload: { // Структура свечи
    figi: string, // FIGI
    interval: string, // Interval
    o: number, // Цена открытия
    c: number, // Цена закрытия
    h: number, // Наибольшая цена
    l: number, // Наименьшая цена
    v: number, // Объем торгов
  }
  time: string // Время в формате RFC3339
}

export default class QuoteClientSocket {
  private socket: Socket
  private assetRepository: IAssetRepository

  constructor() {
    this.socket = new Socket(`${urls.ws.QUOTES}?authorization=${StorageServiceContainer().getAccessTokenWithPrefix()}`)
    this.socket.init()
    this.assetRepository = DIContainer.get(types.IAssetRepository)
  }

  public subscribe(): Unsubscribe {
    return this.socket.subscribe('', '', (message: QuoteData) => {
      this.assetRepository.changePriceAsset({
        identify: message.payload.figi,
        currentPrice: message.payload.c,
      })
    })
  }

  terminate(): void {
    this.socket.terminate()
  }
}
