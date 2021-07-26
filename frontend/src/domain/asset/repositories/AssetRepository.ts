import { inject, injectable } from 'inversify'
import { StoreRoot } from '@/app/store'
import types from '@/infrastructure/di/types'
import { IAsset, IAssetPrice } from '@/domain/asset/entities/AssetEntity'
import {
  FetchAssets,
  ClearAssets,
  AssetAdd,
  ChangePrice,
  LoadingAssets,
} from '@/app/store/modules/asset'

export interface IAssetRepository {
  getAssets(): IAsset[]
  saveAssets(assets: IAsset[]): void
  addAsset(asset: IAsset): void
  clearAssets(): void
  changePriceAsset(assetPrice: IAssetPrice): void
  getLoadingAssets(): boolean
  loadingAssets(status: boolean): void
}

@injectable()
export class AssetRepository implements IAssetRepository {
  constructor(@inject(types.IStoreRoot) private store: StoreRoot) {}

  getAssets(): IAsset[] {
    return this.store.state.asset.assets
  }

  getLoadingAssets(): boolean {
    return this.store.state.asset.loadingItems
  }

  saveAssets(assets: IAsset[]): void {
    this.store.commit(new FetchAssets(assets))
  }

  loadingAssets(status: boolean) {
    this.store.commit(new LoadingAssets(status))
  }

  clearAssets(): void {
    this.store.commit(new ClearAssets())
  }

  addAsset(asset: IAsset): void {
    this.store.commit(new AssetAdd(asset))
  }

  changePriceAsset(assetPrice: IAssetPrice): void {
    this.store.commit(new ChangePrice(assetPrice))
  }
}
