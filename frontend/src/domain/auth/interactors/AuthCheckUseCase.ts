import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/auth/module/types'
import { types as userTypes } from '@/domain/user/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IAuthClientApi } from '@/domain/auth/clients/api/AuthClientApi'
import { ITokenRepository } from '@/domain/auth/repositories/TokenRepository'
import { IUserRepository } from '@/domain/user/repositories/UserRepository'
import { IAuthRepository } from '@/domain/auth/repositories/AuthRepository'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { JwtService } from '@/services/jwt/JwtService'

export interface IAuthCheckUseCase extends BaseUseCase<void, Promise<boolean>> {}

@injectable()
export class AuthCheckUseCase implements IAuthCheckUseCase {
  @inject(types.IAuthClientApi)
  private authClient!: IAuthClientApi

  @inject(types.IAuthRepository)
  private authRepository!: IAuthRepository

  @inject(types.ITokenRepository)
  private tokenRepository!: ITokenRepository

  @inject(userTypes.IUserRepository)
  private userRepository!: IUserRepository

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(): Promise<boolean> {
    try {
      if (this.tokenRepository.hasAccessToken()) {
        const response = await this.authClient.checkLogged()
        const accessToken = new JwtService(this.tokenRepository.getAccessToken()!.token)

        this.userRepository.savePersonData(response.user)
        this.authRepository.setCondition({
          loggedIn: true,
          csrf: response.csrf,
          expire: accessToken.payload.exp,
        })
        return true
      }
    } catch (error) {
      this.errorHandler.handle(error)
    }
    return false
  }
}
