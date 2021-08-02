import { DIContainer } from '@/infrastructure/di'
import { types } from '@/domain/market/module/types'
import { IMarketSearchUseCase } from '@/domain/market/interactors/MarketSearchUseCase'

export const MarketSearchUseCaseDI = (): IMarketSearchUseCase => DIContainer.get(types.IMarketSearchUseCase)
