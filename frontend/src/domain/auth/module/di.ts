import { DIContainer } from '@/infrastructure/di'
import { types } from '@/domain/auth/module/types'
import { IAuthPresenter } from '@/domain/auth/presenter/AuthPresenter'
import { IAuthUseCase } from '@/domain/auth/interactors/AuthUseCase'
import { ITokenRepository } from '@/domain/auth/repositories/TokenRepository'

export const AuthPresenterDI = (): IAuthPresenter => DIContainer.get(types.IAuthPresenter)
export const AuthUseCaseDI = (): IAuthUseCase => DIContainer.get(types.IAuthUseCase)
export const TokenRepositoryDI = (): ITokenRepository => DIContainer.get(types.ITokenRepository)
