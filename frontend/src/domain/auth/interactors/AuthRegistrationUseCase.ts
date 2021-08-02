import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/auth/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IAuthClientApi, IRegisterParams } from '@/domain/auth/clients/api/AuthClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { IMessageService } from '@/services/message/MessageService'
import { INotifyService } from '@/services/notify/NotifyService'

export interface IAuthRegistrationUseCase extends BaseUseCase<IRegisterParams, Promise<boolean>> {}

@injectable()
export class AuthRegistrationUseCase implements IAuthRegistrationUseCase {
  @inject(types.IAuthClientApi)
  private authClient!: IAuthClientApi

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  @inject(baseTypes.IMessageService)
  private messageService!: IMessageService

  @inject(baseTypes.INotifyService)
  private notifyService!: INotifyService

  async execute({ email, name, password }: IRegisterParams): Promise<boolean> {
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
}
