import { MutationTree } from 'vuex'
import { IPortfolio } from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioState {
  active: IPortfolio
}

/**
 ******************************
 * @State
 ******************************
 */
export const state = (): IPortfolioState => ({
  active: {} as IPortfolio,
})

/**
 ******************************
 * @Mutation
 ******************************
 */
export enum Types {
  FETCH_PORTFOLIO = 'portfolio/fetch',
}

export class FetchPortfolio implements FluxStandardAction {
  public type = Types.FETCH_PORTFOLIO
  constructor(public payload: IPortfolio) {}
}

export const mutations: MutationTree<IPortfolioState> = {
  [Types.FETCH_PORTFOLIO]: (state, action: FetchPortfolio) => {
    state.active = action.payload
  },
}
