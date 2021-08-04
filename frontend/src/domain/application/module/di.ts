import { DIContainer } from '@/infrastructure/di'
import { types } from '@/domain/application/module/types'
import { IApplicationDataUseCase } from '@/domain/application/interactors/ApplicationDataUseCase'
import { IApplicationPresenter } from '@/domain/application/presenter/ApplicationPresenter'

export const ApplicationDataUseCaseDI = (): IApplicationDataUseCase => DIContainer.get(types.IApplicationDataUseCase)
export const ApplicationPresenterDI = (): IApplicationPresenter => DIContainer.get(types.IApplicationPresenter)
