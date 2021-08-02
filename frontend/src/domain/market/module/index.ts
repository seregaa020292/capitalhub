import { ContainerModule } from 'inversify'
import { types } from '@/domain/market/module/types'
import { IMarketClientApi, MarketClientApi } from '@/domain/market/clients/api/MarketClientApi'
import {
  IMarketSearchUseCase,
  MarketSearchUseCase,
} from '@/domain/market/interactors/MarketSearchUseCase'
import { IMarketRepository, MarketRepository } from '@/domain/market/repositories/MarketRepository'

export const marketModule = new ContainerModule((bind) => {
  bind<IMarketClientApi>(types.IMarketClientApi).to(MarketClientApi).inSingletonScope()

  bind<IMarketSearchUseCase>(types.IMarketSearchUseCase).to(MarketSearchUseCase).inSingletonScope()

  bind<IMarketRepository>(types.IMarketRepository).to(MarketRepository).inSingletonScope()
})
