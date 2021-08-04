import { inject, injectable } from 'inversify'
import { baseTypes } from '@/infrastructure/di/types'
import { StoreRoot } from '@/app/store'
import { IDashboard } from '@/domain/application/entities/ApplicationEntity'
import { SaveAppDashboard } from '@/app/store/modules/application'

export interface IApplicationRepository {
  saveDashboard(dashboard: IDashboard): void
  getDashboard(): IDashboard
}

@injectable()
export class ApplicationRepository implements IApplicationRepository {
  @inject(baseTypes.IStoreRoot)
  private store!: StoreRoot

  getDashboard(): IDashboard {
    return this.store.state.application.dashboard;
  }

  saveDashboard(dashboard: IDashboard): void {
    this.store.commit(new SaveAppDashboard(dashboard))
  }
}
