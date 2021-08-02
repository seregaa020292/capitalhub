import { ContainerModule } from 'inversify'
import { types } from '@/domain/auth/module/types'
import { AuthClientApi, IAuthClientApi } from '@/domain/auth/clients/api/AuthClientApi'
import { AuthPresenter, IAuthPresenter } from '@/domain/auth/presenter/AuthPresenter'
import { AuthRepository, IAuthRepository } from '@/domain/auth/repositories/AuthRepository'
import { AuthUseCase, IAuthUseCase } from '@/domain/auth/interactors/AuthUseCase'
import { ITokenRepository, TokenRepository } from '@/domain/auth/repositories/TokenRepository'

export const authModule = new ContainerModule((bind) => {
  bind<IAuthClientApi>(types.IAuthClientApi).to(AuthClientApi).inSingletonScope()

  bind<IAuthUseCase>(types.IAuthUseCase).to(AuthUseCase).inSingletonScope()

  bind<IAuthPresenter>(types.IAuthPresenter).to(AuthPresenter).inSingletonScope()

  bind<IAuthRepository>(types.IAuthRepository).to(AuthRepository).inSingletonScope()
  bind<ITokenRepository>(types.ITokenRepository).to(TokenRepository).inSingletonScope()
})
