import { ContainerModule } from 'inversify'
import { types } from '@/domain/auth/module/types'
import { AuthClientApi, IAuthClientApi } from '@/domain/auth/clients/api/AuthClientApi'
import { AuthPresenter, IAuthPresenter } from '@/domain/auth/presenter/AuthPresenter'
import { AuthRepository, IAuthRepository } from '@/domain/auth/repositories/AuthRepository'
import { ITokenRepository, TokenRepository } from '@/domain/auth/repositories/TokenRepository'
import {
  AuthAutoLogoutUseCase,
  IAuthAutoLogoutUseCase,
} from '@/domain/auth/interactors/AuthAutoLogoutUseCase'
import { AuthCheckUseCase, IAuthCheckUseCase } from '@/domain/auth/interactors/AuthCheckUseCase'
import {
  AuthConfirmedUseCase,
  IAuthConfirmedUseCase,
} from '@/domain/auth/interactors/AuthConfirmedUseCase'
import { AuthLoginUseCase, IAuthLoginUseCase } from '@/domain/auth/interactors/AuthLoginUseCase'
import { AuthLogoutUseCase, IAuthLogoutUseCase } from '@/domain/auth/interactors/AuthLogoutUseCase'
import {
  AuthRefreshTokenUseCase,
  IAuthRefreshTokenUseCase,
} from '@/domain/auth/interactors/AuthRefreshTokenUseCase'
import {
  AuthRegistrationUseCase,
  IAuthRegistrationUseCase,
} from '@/domain/auth/interactors/AuthRegistrationUseCase'

export const authModule = new ContainerModule((bind) => {
  bind<IAuthClientApi>(types.IAuthClientApi).to(AuthClientApi).inSingletonScope()

  bind<IAuthAutoLogoutUseCase>(types.IAuthAutoLogoutUseCase).to(AuthAutoLogoutUseCase).inSingletonScope()
  bind<IAuthCheckUseCase>(types.IAuthCheckUseCase).to(AuthCheckUseCase).inSingletonScope()
  bind<IAuthConfirmedUseCase>(types.IAuthConfirmedUseCase).to(AuthConfirmedUseCase).inSingletonScope()
  bind<IAuthLoginUseCase>(types.IAuthLoginUseCase).to(AuthLoginUseCase).inSingletonScope()
  bind<IAuthLogoutUseCase>(types.IAuthLogoutUseCase).to(AuthLogoutUseCase).inSingletonScope()
  bind<IAuthRefreshTokenUseCase>(types.IAuthRefreshTokenUseCase).to(AuthRefreshTokenUseCase).inSingletonScope()
  bind<IAuthRegistrationUseCase>(types.IAuthRegistrationUseCase).to(AuthRegistrationUseCase).inSingletonScope()

  bind<IAuthPresenter>(types.IAuthPresenter).to(AuthPresenter).inSingletonScope()

  bind<IAuthRepository>(types.IAuthRepository).to(AuthRepository).inSingletonScope()
  bind<ITokenRepository>(types.ITokenRepository).to(TokenRepository).inSingletonScope()
})
