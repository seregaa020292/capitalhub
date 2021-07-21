import { inject, injectable } from 'inversify'
import types from '@/infrastructure/di/types'
import { IUserRepository } from '@/domain/user/repositories/UserRepository'
import { IUser } from '@/domain/user/entities/UserEntity'

export interface IUserPresenter {
  user(): IUser
  isAvatarEmpty(): boolean
}

@injectable()
export class UserPresenter implements IUserPresenter {
  @inject(types.IUserRepository)
  private userRepository!: IUserRepository

  isAvatarEmpty(): boolean {
    return !!this.user().avatar
  }

  user(): IUser {
    return this.userRepository.getUser()
  }
}
