import { inject, injectable } from 'inversify'
import { types } from '@/domain/application/module/types'
import { IApplicationRepository } from '@/domain/application/repositories/ApplicationRepository'
import { IDashboard } from '@/domain/application/entities/ApplicationEntity'

export interface IApplicationPresenter {
  getDashboard(): IDashboard
}

@injectable()
export class ApplicationPresenter implements IApplicationPresenter {
  @inject(types.IApplicationRepository)
  private applicationRepository!: IApplicationRepository

  getDashboard(): IDashboard {
    return this.applicationRepository.getDashboard()
  }
}
