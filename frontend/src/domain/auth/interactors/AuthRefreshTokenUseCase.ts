import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/auth/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IAuthClientApi } from '@/domain/auth/clients/api/AuthClientApi'
import { ITokenRepository } from '@/domain/auth/repositories/TokenRepository'
import { IAuthRepository } from '@/domain/auth/repositories/AuthRepository'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { JwtService } from '@/services/jwt/JwtService'
import { initFingerprintJS } from '@/utils/device'

export interface IAuthRefreshTokenUseCase extends BaseUseCase<void, Promise<void>> {}

@injectable()
export class AuthRefreshTokenUseCase implements IAuthRefreshTokenUseCase {
  @inject(types.IAuthClientApi)
  private authClient!: IAuthClientApi

  @inject(types.IAuthRepository)
  private authRepository!: IAuthRepository

  @inject(types.ITokenRepository)
  private tokenRepository!: ITokenRepository

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(): Promise<void> {
    try {
      const fingerprint = await initFingerprintJS()
      const response = await this.authClient.refresh({ fingerprint: fingerprint.visitorId })
      const accessToken = new JwtService(response.accessToken.token)

      this.tokenRepository.saveAccessToken(response.accessToken)
      this.authRepository.setCondition({
        loggedIn: true,
        csrf: response.csrf,
        expire: accessToken.payload.exp,
      })
    } catch (error) {
      this.errorHandler.handle(error)
      throw error
    }
  }
}
