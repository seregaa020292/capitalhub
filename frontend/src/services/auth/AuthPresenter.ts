import { inject, injectable } from 'inversify'
import types from '@/infrastructure/di/types'
import { IAuthRepository } from '@/services/auth/AuthRepository'

export interface IAuthPresenter {
  loggedIn(): boolean
  getCsrf(): string
}

@injectable()
export class AuthPresenter implements IAuthPresenter {
  @inject(types.IAuthRepository)
  private authRepository!: IAuthRepository

  loggedIn(): boolean {
    return this.authRepository.getState().loggedIn
  }

  getCsrf(): string {
    return this.authRepository.getState().csrf
  }
}
