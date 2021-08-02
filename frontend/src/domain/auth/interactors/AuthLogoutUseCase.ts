import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/auth/module/types'
import { types as userTypes } from '@/domain/user/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IAuthClientApi } from '@/domain/auth/clients/api/AuthClientApi'
import { ITokenRepository } from '@/domain/auth/repositories/TokenRepository'
import { IUserRepository } from '@/domain/user/repositories/UserRepository'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'

export interface IAuthLogoutUseCase extends BaseUseCase<void, Promise<void>> {}

@injectable()
export class AuthLogoutUseCase implements IAuthLogoutUseCase {
  @inject(types.IAuthClientApi)
  private authClient!: IAuthClientApi

  @inject(types.ITokenRepository)
  private tokenRepository!: ITokenRepository

  @inject(userTypes.IUserRepository)
  private userRepository!: IUserRepository

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(): Promise<void> {
    try {
      await this.authClient.logout()
    } catch (error) {
      this.errorHandler.handle(error)
    } finally {
      this.tokenRepository.removeAccessToken()
      this.userRepository.clearPersonData()
    }
  }
}
