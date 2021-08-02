import { ContainerModule } from 'inversify'
import { types } from '@/domain/asset/module/types'
import { AssetClientApi, IAssetClientApi } from '@/domain/asset/clients/api/AssetClientApi'
import { AssetFetchUseCase, IAssetFetchUseCase } from '@/domain/asset/interactors/AssetFetchUseCase'
import { AssetAddUseCase, IAssetAddUseCase } from '@/domain/asset/interactors/AssetAddUseCase'
import { AssetPresenter, IAssetPresenter } from '@/domain/asset/presenter/AssetPresenter'
import { AssetRepository, IAssetRepository } from '@/domain/asset/repositories/AssetRepository'

export const assetModule = new ContainerModule((bind) => {
  bind<IAssetClientApi>(types.IAssetClientApi).to(AssetClientApi).inSingletonScope()

  bind<IAssetFetchUseCase>(types.IAssetFetchUseCase).to(AssetFetchUseCase).inSingletonScope()
  bind<IAssetAddUseCase>(types.IAssetAddUseCase).to(AssetAddUseCase).inSingletonScope()

  bind<IAssetPresenter>(types.IAssetPresenter).to(AssetPresenter).inSingletonScope()

  bind<IAssetRepository>(types.IAssetRepository).to(AssetRepository).inSingletonScope()
})
