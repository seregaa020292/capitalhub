import { inject, injectable } from 'inversify'
import { StoreRoot } from '@/app/store'
import types from '@/infrastructure/di/types'
import { IAsset, IAssetPrice } from '@/domain/asset/entities/AssetEntity'
import { FetchAssets, ClearAssets, AssetAdd, ChangePrice } from '@/app/store/modules/asset'

export interface IAssetRepository {
  getAssets(): IAsset[]
  saveAssets(assets: IAsset[]): void
  addAsset(asset: IAsset): void
  clearAssets(): void
  changePriceAsset(assetPrice: IAssetPrice): void
}

@injectable()
export class AssetRepository implements IAssetRepository {
  constructor(@inject(types.IStoreRoot) private store: StoreRoot) {}

  getAssets(): IAsset[] {
    return this.store.state.asset.assets
  }

  saveAssets(assets: IAsset[]): void {
    this.store.commit(new FetchAssets(assets))
  }

  clearAssets(): void {
    this.store.commit(new ClearAssets())
  }

  addAsset(asset: IAsset): void {
    this.store.commit(new AssetAdd(asset))
  }

  changePriceAsset(assetPrice: IAssetPrice): void{
    this.store.commit(new ChangePrice(assetPrice))
  }
}
