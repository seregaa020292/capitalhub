import { inject, injectable } from 'inversify'
import types from '@/infrastructure/di/types'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { IStorageService } from '@/services/auth/StorageService'
import { INotifyService } from '@/services/notify/NotifyService'
import { IMessageService } from '@/services/message/MessageService'
import { IRouterService } from '@/services/router/RouterService'
import { JwtService } from '@/services/auth/JwtService'
import { IUserRepository } from '@/domain/user/repositories/UserRepository'
import { IAuthRepository } from '@/services/auth/AuthRepository'
import {
  IAuthClientApi,
  ILoginParams,
  IRegisterParams,
} from '@/services/api/AuthClientApi'
import { initFingerprintJS } from '@/utils/device'

export interface IAuthService {
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
export class AuthService implements IAuthService {
  @inject(types.IAuthClientApi)
  private authClient!: IAuthClientApi

  @inject(types.IStorageService)
  private storageService!: IStorageService

  @inject(types.IUserRepository)
  private userRepository!: IUserRepository

  @inject(types.IAuthRepository)
  private authRepository!: IAuthRepository

  @inject(types.IErrorHandler)
  private errorHandler!: IErrorHandler

  @inject(types.IMessageService)
  private messageService!: IMessageService

  @inject(types.INotifyService)
  private notifyService!: INotifyService

  @inject(types.IRouterService)
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
      this.storageService.saveAccessToken(response.accessToken)
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

      this.storageService.saveAccessToken(response.accessToken)
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
      if (this.storageService.hasAccessToken()) {
        const response = await this.authClient.checkLogged()
        const accessToken = new JwtService(this.storageService.getAccessToken()!.token)

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
      this.storageService.removeAccessToken()
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
