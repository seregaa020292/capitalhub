import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import types from '@/infrastructure/di/types'
import { IMarketGroup, IMarketOption } from '@/domain/market/entities/MarketEntity'
import { IMarketClientApi } from '@/services/api/MarketClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'

export interface IMarketSearchUseCase extends BaseUseCase<string, Promise<IMarketOption[] | null>> {
}

@injectable()
export class MarketSearchUseCase implements IMarketSearchUseCase {
  @inject(types.IMarketClientApi)
  private marketClient!: IMarketClientApi

  @inject(types.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(title: string) {
    try {
      const response = await this.marketClient.searchMarkets(title)

      return Object.values(response.markets.reduce((curObj, curItem) => {
        if (!curObj.hasOwnProperty(curItem.title_instrument)) {
          curObj[curItem.title_instrument] = {
            label: curItem.title_instrument,
            description: curItem.desc_instrument,
            options: [],
          }
        }
        curObj[curItem.title_instrument].options.push({ ...curItem })
        return curObj
      }, {} as IMarketGroup))
    } catch (error) {
      this.errorHandler.handle(error)
    }
    return null
  }
}
