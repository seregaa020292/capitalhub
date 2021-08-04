import { DIContainer } from '@/infrastructure/di'
import { types } from '@/domain/auth/module/types'
import { IAuthPresenter } from '@/domain/auth/presenters/AuthPresenter'
import { ITokenRepository } from '@/domain/auth/repositories/TokenRepository'
import { IAuthRefreshTokenUseCase } from '@/domain/auth/interactors/AuthRefreshTokenUseCase'
import { IAuthLogoutUseCase } from '@/domain/auth/interactors/AuthLogoutUseCase'
import { IAuthLoginUseCase } from '@/domain/auth/interactors/AuthLoginUseCase'
import { IAuthConfirmedUseCase } from '@/domain/auth/interactors/AuthConfirmedUseCase'
import { IAuthCheckUseCase } from '@/domain/auth/interactors/AuthCheckUseCase'
import { IAuthAutoLogoutUseCase } from '@/domain/auth/interactors/AuthAutoLogoutUseCase'
import { IAuthRegistrationUseCase } from '@/domain/auth/interactors/AuthRegistrationUseCase'

export const AuthPresenterDI = (): IAuthPresenter => DIContainer.get(types.IAuthPresenter)

export const AuthRefreshTokenUseCaseDI = (): IAuthRefreshTokenUseCase => DIContainer.get(types.IAuthRefreshTokenUseCase)
export const AuthLogoutUseCaseDI = (): IAuthLogoutUseCase => DIContainer.get(types.IAuthLogoutUseCase)
export const AuthLoginUseCaseDI = (): IAuthLoginUseCase => DIContainer.get(types.IAuthLoginUseCase)
export const AuthConfirmedUseCaseDI = (): IAuthConfirmedUseCase => DIContainer.get(types.IAuthConfirmedUseCase)
export const AuthCheckUseCaseDI = (): IAuthCheckUseCase => DIContainer.get(types.IAuthCheckUseCase)
export const AuthAutoLogoutUseCaseDI = (): IAuthAutoLogoutUseCase => DIContainer.get(types.IAuthAutoLogoutUseCase)
export const AuthRegistrationUseCaseDI = (): IAuthRegistrationUseCase => DIContainer.get(types.IAuthRegistrationUseCase)

export const TokenRepositoryDI = (): ITokenRepository => DIContainer.get(types.ITokenRepository)
