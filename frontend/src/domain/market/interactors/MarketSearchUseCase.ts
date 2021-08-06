import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/market/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IMarketGroup, IMarketOption } from '@/domain/market/entities/MarketEntity'
import { IMarketClientApi } from '@/domain/market/clients/api/MarketClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'

export interface IMarketSearchUseCase extends BaseUseCase<string, Promise<IMarketOption[] | null>> {
}

@injectable()
export class MarketSearchUseCase implements IMarketSearchUseCase {
  @inject(types.IMarketClientApi)
  private marketClient!: IMarketClientApi

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(title: string) {
    try {
      const response = await this.marketClient.searchMarkets(title)

      return Object.values(response.markets.reduce((curObj, curItem) => {
        if (!curObj.hasOwnProperty(curItem.titleInstrument)) {
          curObj[curItem.titleInstrument] = {
            label: curItem.titleInstrument,
            description: curItem.descInstrument,
            options: [],
          }
        }
        curObj[curItem.titleInstrument].options.push({ ...curItem })
        return curObj
      }, {} as IMarketGroup))
    } catch (error) {
      this.errorHandler.handle(error)
    }
    return null
  }
}
