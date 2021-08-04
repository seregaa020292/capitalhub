import { MutationTree } from 'vuex'
import { IDashboard } from '@/domain/application/entities/ApplicationEntity'

export interface IApplicationState {
  dashboard: IDashboard
}

/**
 ******************************
 * @State
 ******************************
 */
export const state = (): IApplicationState => ({
  dashboard: {} as IDashboard,
})

/**
 ******************************
 * @Mutation
 ******************************
 */
export enum Types {
  SAVE_DASHBOARD = 'application/save_dashboard',
}

export class SaveAppDashboard implements FluxStandardAction {
  public type = Types.SAVE_DASHBOARD
  constructor(public payload: IDashboard) {}
}

export const mutations: MutationTree<IApplicationState> = {
  [Types.SAVE_DASHBOARD]: (state, action: SaveAppDashboard) => {
    state.dashboard = action.payload
  },
}
