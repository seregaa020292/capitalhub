import { DIContainer } from '@/infrastructure/di'
import { types } from '@/domain/asset/module/types'
import { IAssetPresenter } from '@/domain/asset/presenters/AssetPresenter'
import { IAssetAddUseCase } from '@/domain/asset/interactors/AssetAddUseCase'
import { IAssetRepository } from '@/domain/asset/repositories/AssetRepository'

export const AssetPresenterDI = (): IAssetPresenter => DIContainer.get(types.IAssetPresenter)
export const AssetAddUseCaseDI = (): IAssetAddUseCase => DIContainer.get(types.IAssetAddUseCase)
export const AssetRepositoryDI = (): IAssetRepository => DIContainer.get(types.IAssetRepository)
