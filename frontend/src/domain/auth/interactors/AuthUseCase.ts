import { inject, injectable } from 'inversify'
import { baseTypes } from '@/infrastructure/di/types'
import { types } from '@/domain/auth/module/types'
import { types as userTypes } from '@/domain/user/module/types'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { ITokenRepository } from '@/domain/auth/repositories/TokenRepository'
import { INotifyService } from '@/services/notify/NotifyService'
import { IMessageService } from '@/services/message/MessageService'
import { IRouterService } from '@/services/router/RouterService'
import { JwtService } from '@/services/jwt/JwtService'
import { IUserRepository } from '@/domain/user/repositories/UserRepository'
import { IAuthRepository } from '@/domain/auth/repositories/AuthRepository'
import {
  IAuthClientApi,
  ILoginParams,
  IRegisterParams,
} from '@/domain/auth/clients/api/AuthClientApi'
import { initFingerprintJS } from '@/utils/device'

export interface IAuthUseCase {
  login(credentials: ILoginParams): Promise<boolean>
  registration(credentials: IRegisterParams): Promise<boolean>
  refreshToken(): Promise<void>
  checkLogged(): Promise<boolean>
  logout(): Promise<void>
  autoLogout(): Promise<void>
  confirmed(code: string): Promise<boolean>
}

/**
 ******************************
 * @class AuthService
 ******************************
 */
@injectable()
export class AuthUseCase implements IAuthUseCase {
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

  @inject(baseTypes.IMessageService)
  private messageService!: IMessageService

  @inject(baseTypes.INotifyService)
  private notifyService!: INotifyService

  @inject(baseTypes.IRouterService)
  private routerService!: IRouterService

  public async login({ email, password }: ILoginParams): Promise<boolean> {
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

  public async registration({ email, name, password }: IRegisterParams): Promise<boolean> {
    try {
      await this.authClient.register({ email, name, password })

      this.messageService.success(`Спасибо ${name}. Регистрация прошла успешно.`)
      this.notifyService.success(`Сообщение на подтверждение, выслано на почту <u>${email}</u>.`)
      return true
    } catch (error) {
      this.errorHandler.handle(error).report()
      return false
    }
  }

  public async refreshToken(): Promise<void> {
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

  public async checkLogged(): Promise<boolean> {
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

  public async logout(): Promise<void> {
    try {
      await this.authClient.logout()
    } catch (error) {
      this.errorHandler.handle(error)
    } finally {
      this.tokenRepository.removeAccessToken()
      this.userRepository.clearPersonData()
    }
  }

  public async autoLogout(): Promise<void> {
    await this.logout()

    this.messageService.warning('Срок действия учетных данных истек, пожалуйста, войдите снова')

    if (!['/', '/login', '/register'].includes(this.routerService.currentPath)) {
      this.routerService.push({ name: 'login' })
    }
  }

  public async confirmed(code: string): Promise<boolean> {
    try {
      await this.authClient.confirmed(code)
      return true
    } catch (error) {
      this.errorHandler.handle(error)
    }
    return false
  }
}
