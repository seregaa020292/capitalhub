import { urls } from '@/infrastructure/network/urls'
import { Socket, Unsubscribe } from '@/infrastructure/network/socket'
import { TokenRepositoryDI } from '@/domain/auth/module/di'
import { IAssetRepository } from '@/domain/asset/repositories/AssetRepository'
import { AssetRepositoryDI } from '@/domain/asset/module/di'

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
    this.socket = new Socket(`${urls.ws.QUOTES}?authorization=${TokenRepositoryDI().getAccessTokenWithPrefix()}`)
    this.socket.init()
    this.assetRepository = AssetRepositoryDI()
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
