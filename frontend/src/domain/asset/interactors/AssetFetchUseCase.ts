import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/asset/module/types'
import { IAssetRepository } from '@/domain/asset/repositories/AssetRepository'
import { IAsset } from '@/domain/asset/entities/AssetEntity'
import { IAssetClientApi } from '@/domain/asset/clients/api/AssetClientApi'

export  interface IAssetFetchUseCase extends BaseUseCase<unknown, Promise<void>> {}

@injectable()
export class AssetFetchUseCase implements IAssetFetchUseCase {
  @inject(types.IAssetClientApi)
  private assetClient!: IAssetClientApi

  @inject(types.IAssetRepository)
  private assetRepository!: IAssetRepository

  async execute() {
    const response: IAsset[] = await this.assetClient.fetchTotalAssets()
    this.assetRepository.saveAssets(response)
  }
}
