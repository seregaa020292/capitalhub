import { inject, injectable } from 'inversify'
import { StoreRoot } from '@/app/store'
import { baseTypes } from '@/infrastructure/di/types'
import { IUser } from '@/domain/user/entities/UserEntity'
import { ClearUserPersonData, FetchUserPersonData } from '@/app/store/modules/user'

export interface IUserRepository {
  getUser(): IUser
  savePersonData(user: IUser): void
  clearPersonData(): void
}

@injectable()
export class UserRepository implements IUserRepository {
  constructor(@inject(baseTypes.IStoreRoot) private store: StoreRoot) {}

  getUser(): IUser {
    return this.store.state.user.personData
  }

  savePersonData(user: IUser): void {
    this.store.commit(new FetchUserPersonData(user))
  }

  clearPersonData(): void {
    this.store.commit(new ClearUserPersonData())
  }
}
