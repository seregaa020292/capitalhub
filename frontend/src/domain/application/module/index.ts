import { ContainerModule } from 'inversify'
import { types } from '@/domain/application/module/types'
import {
  ApplicationDataUseCase,
  IApplicationDataUseCase,
} from '@/domain/application/interactors/ApplicationDataUseCase'
import {
  ApplicationClientApi,
  IApplicationClientApi,
} from '@/domain/application/clients/api/ApplicationClientApi'
import {
  ApplicationRepository,
  IApplicationRepository,
} from '@/domain/application/repositories/ApplicationRepository'
import {
  ApplicationPresenter,
  IApplicationPresenter,
} from '@/domain/application/presenter/ApplicationPresenter'

export const applicationModule = new ContainerModule((bind) => {
  bind<IApplicationClientApi>(types.IApplicationClientApi).to(ApplicationClientApi).inSingletonScope()
  bind<IApplicationDataUseCase>(types.IApplicationDataUseCase).to(ApplicationDataUseCase).inSingletonScope()
  bind<IApplicationRepository>(types.IApplicationRepository).to(ApplicationRepository).inSingletonScope()
  bind<IApplicationPresenter>(types.IApplicationPresenter).to(ApplicationPresenter).inSingletonScope()
})
