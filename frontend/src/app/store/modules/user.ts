import { MutationTree } from 'vuex'
import { IUser } from '@/domain/user/entities/UserEntity'

export interface IUserState {
  personData: IUser
}

/**
 ******************************
 * @State
 ******************************
 */
const initPersonData: IUser = {
  user_id: '',
  name: '',
  email: '',
  role: '',
  avatar: '',
}

export const state = (): IUserState => ({
  personData: { ...initPersonData },
})

/**
 ******************************
 * @Mutation
 ******************************
 */
export enum Types {
  FETCH_USER_PERSON_DATA = 'user/fetch_person_data',
  CLEAR_USER_PERSON_DATA = 'user/clear_person_data',
}

export class FetchUserPersonData implements FluxStandardAction {
  public type = Types.FETCH_USER_PERSON_DATA
  constructor(public payload: IUser) {}
}

export class ClearUserPersonData implements FluxStandardAction {
  public type = Types.CLEAR_USER_PERSON_DATA
}

export const mutations: MutationTree<IUserState> = {
  [Types.FETCH_USER_PERSON_DATA]: (state, action: FetchUserPersonData) => {
    state.personData = action.payload
  },
  [Types.CLEAR_USER_PERSON_DATA]: (state) => {
    state.personData = { ...initPersonData }
  },
}
