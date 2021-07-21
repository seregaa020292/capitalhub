import { inject, injectable } from 'inversify'
import { StoreRoot } from '@/app/store'
import { FetchAuth, IAuthState } from '@/app/store/modules/auth'
import types from '@/infrastructure/di/types'

export interface IAuthRepository {
  getState(): IAuthState
  setCondition(condition: IAuthState): void
}

@injectable()
export class AuthRepository implements IAuthRepository {
  @inject(types.IStoreRoot)
  private store!: StoreRoot

  getState(): IAuthState {
    return this.store.state.auth
  }

  setCondition(condition: IAuthState): void {
    this.store.commit(new FetchAuth(condition))
  }
}
