import { MutationTree } from 'vuex'

export interface IAuthState {
  loggedIn: boolean
  expire: number | null
  csrf: string
}

/**
 ******************************
 * @State
 ******************************
 */
export const state = (): IAuthState => ({
  loggedIn: false,
  expire: null,
  csrf: '',
})

/**
 ******************************
 * @Mutation
 ******************************
 */
export enum Types {
  FETCH_AUTH = 'auth/fetch',
  CLEAR_AUTH = 'auth/clear',
}

export class FetchAuth implements FluxStandardAction {
  public type = Types.FETCH_AUTH
  constructor(public payload: IAuthState) {}
}

export class ClearAuth implements FluxStandardAction {
  public type = Types.CLEAR_AUTH
}

export const mutations: MutationTree<IAuthState> = {
  [Types.FETCH_AUTH]: (state, action: FetchAuth) => {
    state.expire = action.payload.expire
    state.loggedIn = action.payload.loggedIn
    state.csrf = action.payload.csrf
  },
  [Types.CLEAR_AUTH]: (state) => {
    state.expire = null
    state.loggedIn = false
    state.csrf = ''
  },
}
