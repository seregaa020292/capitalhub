import { MutationTree } from 'vuex'
import { IPortfolio, IPortfolioStats } from '@/domain/portfolio/entities/PortfolioEntity'

export interface IPortfolioState {
  active: IPortfolio
  all: IPortfolioStats[]
  loadingAll: boolean
}

/**
 ******************************
 * @State
 ******************************
 */
export const state = (): IPortfolioState => ({
  active: {} as IPortfolio,
  all: [],
  loadingAll: false,
})

/**
 ******************************
 * @Mutation
 ******************************
 */
export enum Types {
  FETCH_PORTFOLIO = 'portfolio/fetch',
  FETCH_PORTFOLIOS = 'portfolio/fetch-all',
  FETCH_LOADING_PORTFOLIOS = 'portfolio/loading-all',
  ADD_PORTFOLIO = 'portfolio/add',
  EDIT_PORTFOLIO = 'portfolio/edit',
}

export class FetchPortfolio implements FluxStandardAction {
  public type = Types.FETCH_PORTFOLIO
  constructor(public payload: IPortfolio) {}
}

export class FetchPortfolios implements FluxStandardAction {
  public type = Types.FETCH_PORTFOLIOS
  constructor(public payload: IPortfolioStats[]) {}
}

export class LoadingPortfolios implements FluxStandardAction {
  public type = Types.FETCH_LOADING_PORTFOLIOS
  constructor(public payload: boolean) {}
}

export class AddPortfolio implements FluxStandardAction {
  public type = Types.ADD_PORTFOLIO
  constructor(public payload: IPortfolioStats) {}
}

export class EditPortfolio implements FluxStandardAction {
  public type = Types.EDIT_PORTFOLIO
  constructor(public payload: IPortfolioStats) {}
}

export const mutations: MutationTree<IPortfolioState> = {
  [Types.FETCH_PORTFOLIO]: (state, action: FetchPortfolio) => {
    state.active = action.payload
  },
  [Types.FETCH_PORTFOLIOS]: (state, action: FetchPortfolios) => {
    state.all = action.payload
  },
  [Types.FETCH_LOADING_PORTFOLIOS]: (state, action: LoadingPortfolios) => {
    state.loadingAll = action.payload
  },
  [Types.ADD_PORTFOLIO]: (state, action: AddPortfolio) => {
    state.all.push(action.payload)
  },
  [Types.EDIT_PORTFOLIO]: (state, action: EditPortfolio) => {
    const portfolioIdx = state.all.findIndex(({ portfolioId }) => portfolioId === action.payload.portfolioId)
    if (portfolioIdx !== -1) {
      state.all[portfolioIdx] = action.payload
    }
  },
}
