import { injectable } from 'inversify'
import { http } from '@/infrastructure/network/http'
import urls from '@/infrastructure/network/urls'
import { IAsset, IAssetNotation } from '@/domain/asset/entities/AssetEntity'

export interface IAssetClientApi {
  fetchTotalAssets(): Promise<IAsset[]>
  addAsset(assetNotation: IAssetNotation): Promise<IAsset>
}

@injectable()
export class AssetClientApi implements IAssetClientApi {
  fetchTotalAssets(): Promise<IAsset[]> {
    return http.get(urls.api_v1.TOTAL_ASSETS)
  }

  addAsset(assetNotation: IAssetNotation): Promise<IAsset> {
    return http.post(urls.api_v1.ASSET_ADD, assetNotation)
  }
}
