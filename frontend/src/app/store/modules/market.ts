import { MutationTree } from 'vuex'
import { IMarket } from '@/domain/market/entities/MarketEntity'

export interface IMarketState {
  markets: IMarket[]
}

/**
 ******************************
 * @State
 ******************************
 */
export const state = (): IMarketState => ({
  markets: [],
})

/**
 ******************************
 * @Mutation
 ******************************
 */
export enum Types {
  FETCH_MARKETS = 'market/fetch',
  CLEAR_MARKETS = 'market/clear',
}

export class FetchMarkets implements FluxStandardAction {
  public type = Types.FETCH_MARKETS
  constructor(public payload: IMarket[]) {}
}

export class ClearMarkets implements FluxStandardAction {
  public type = Types.CLEAR_MARKETS
}

export const mutations: MutationTree<IMarketState> = {
  [Types.FETCH_MARKETS]: (state, action: FetchMarkets) => {
    state.markets = action.payload
  },
  [Types.CLEAR_MARKETS]: (state) => {
    state.markets = []
  },
}
