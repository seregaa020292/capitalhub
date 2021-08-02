import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/auth/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IAuthClientApi } from '@/domain/auth/clients/api/AuthClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'

export interface IAuthConfirmedUseCase extends BaseUseCase<string, Promise<boolean>> {}

@injectable()
export class AuthConfirmedUseCase implements IAuthConfirmedUseCase {
  @inject(types.IAuthClientApi)
  private authClient!: IAuthClientApi

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(code: string): Promise<boolean> {
    try {
      await this.authClient.confirmed(code)
      return true
    } catch (error) {
      this.errorHandler.handle(error)
    }
    return false
  }
}
