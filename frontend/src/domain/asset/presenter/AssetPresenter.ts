import { inject, injectable } from 'inversify'
import types from '@/infrastructure/di/types'
import { IAssetRepository } from '@/domain/asset/repositories/AssetRepository'
import { IAsset } from '@/domain/asset/entities/AssetEntity'

export interface IAssetPresenter {
  assets(): IAsset[]
  loadingAssets(): boolean
}

@injectable()
export class AssetPresenter implements IAssetPresenter {
  @inject(types.IAssetRepository)
  private assetRepository!: IAssetRepository

  assets(): IAsset[] {
    return this.assetRepository.getAssets()
  }

  loadingAssets(): boolean {
    return this.assetRepository.getLoadingAssets()
  }
}
