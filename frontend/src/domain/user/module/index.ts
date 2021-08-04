import { ContainerModule } from 'inversify'
import { types } from '@/domain/user/module/types'
import { IUserClientApi, UserClientApi } from '@/domain/user/clients/api/UserClientApi'
import { IUserFetchUseCase, UserFetchUseCase } from '@/domain/user/interactors/UserFetchUseCase'
import { IUserPresenter, UserPresenter } from '@/domain/user/presenters/UserPresenter'
import { IUserRepository, UserRepository } from '@/domain/user/repositories/UserRepository'

export const userModule = new ContainerModule((bind) => {
  bind<IUserClientApi>(types.IUserClientApi).to(UserClientApi).inSingletonScope()

  bind<IUserFetchUseCase>(types.IUserFetchUseCase).to(UserFetchUseCase).inSingletonScope()

  bind<IUserPresenter>(types.IUserPresenter).to(UserPresenter).inSingletonScope()

  bind<IUserRepository>(types.IUserRepository).to(UserRepository).inSingletonScope()
})
