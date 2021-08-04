import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/application/module/types'
import { baseTypes } from '@/infrastructure/di/types'
import { IApplicationClientApi } from '@/domain/application/clients/api/ApplicationClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'
import { IApplicationRepository } from '@/domain/application/repositories/ApplicationRepository'

export interface IApplicationDataUseCase extends BaseUseCase<undefined, Promise<void>>{}

@injectable()
export class ApplicationDataUseCase implements IApplicationDataUseCase {
  @inject(types.IApplicationClientApi)
  private applicationClientApi!: IApplicationClientApi

  @inject(types.IApplicationRepository)
  private applicationRepository!: IApplicationRepository

  @inject(baseTypes.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(): Promise<void> {
    try {
      const dashboard = await this.applicationClientApi.dashboard()
      this.applicationRepository.saveDashboard(dashboard)
    } catch (error) {
      this.errorHandler.handle(error)
    }
  }
}
