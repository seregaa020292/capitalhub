import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/auth/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IRouterService } from '@/services/router/RouterService'
import { IMessageService } from '@/services/message/MessageService'
import { IAuthLogoutUseCase } from '@/domain/auth/interactors/AuthLogoutUseCase'

export interface IAuthAutoLogoutUseCase extends BaseUseCase<void, Promise<void>> {}

@injectable()
export class AuthAutoLogoutUseCase implements IAuthAutoLogoutUseCase {
  @inject(types.IAuthLogoutUseCase)
  private authLogoutUseCase!: IAuthLogoutUseCase

  @inject(baseTypes.IMessageService)
  private messageService!: IMessageService

  @inject(baseTypes.IRouterService)
  private routerService!: IRouterService

  async execute(): Promise<void> {
    await this.authLogoutUseCase.execute()

    this.messageService.warning('Срок действия учетных данных истек, пожалуйста, войдите снова')

    if (!['/', '/login', '/register'].includes(this.routerService.currentPath)) {
      this.routerService.push({ name: 'login' })
    }
  }
}
