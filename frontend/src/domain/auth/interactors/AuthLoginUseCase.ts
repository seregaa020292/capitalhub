import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/auth/module/types'
import { types as userTypes } from '@/domain/user/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IAuthClientApi, ILoginParams } from '@/domain/auth/clients/api/AuthClientApi'
import { ITokenRepository } from '@/domain/auth/repositories/TokenRepository'
import { IUserRepository } from '@/domain/user/repositories/UserRepository'
import { IAuthRepository } from '@/domain/auth/repositories/AuthRepository'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { JwtService } from '@/services/jwt/JwtService'
import { initFingerprintJS } from '@/utils/device'

export interface IAuthLoginUseCase extends BaseUseCase<ILoginParams, Promise<boolean>> {}

@injectable()
export class AuthLoginUseCase implements IAuthLoginUseCase {
  @inject(types.IAuthClientApi)
  private authClient!: IAuthClientApi

  @inject(types.ITokenRepository)
  private tokenRepository!: ITokenRepository

  @inject(userTypes.IUserRepository)
  private userRepository!: IUserRepository

  @inject(types.IAuthRepository)
  private authRepository!: IAuthRepository

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute({ email, password }: ILoginParams): Promise<boolean> {
    try {
      const fingerprint = await initFingerprintJS()
      const response = await this.authClient.login({
        fingerprint: fingerprint.visitorId,
        email,
        password,
      })
      const accessToken = new JwtService(response.accessToken.token)

      this.userRepository.savePersonData(response.user)
      this.tokenRepository.saveAccessToken(response.accessToken)
      this.authRepository.setCondition({
        loggedIn: true,
        csrf: response.csrf,
        expire: accessToken.payload.exp,
      })
      return true
    } catch (error) {
      this.errorHandler.handle(error).report()
      return false
    }
  }
}
