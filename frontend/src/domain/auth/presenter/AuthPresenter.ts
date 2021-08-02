import { inject, injectable } from 'inversify'
import { types } from '@/domain/auth/module/types'
import { IAuthRepository } from '@/domain/auth/repositories/AuthRepository'

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
