import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import types from '@/infrastructure/di/types'
import { IAssetRepository } from '@/domain/asset/repositories/AssetRepository'
import { IAsset, IAssetNotation } from '@/domain/asset/entities/AssetEntity'
import { IAssetClientApi } from '@/services/api/AssetClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'

export  interface IAssetAddUseCase extends BaseUseCase<IAssetNotation, Promise<void>> {}

@injectable()
export class AssetAddUseCase implements IAssetAddUseCase {
  @inject(types.IAssetClientApi)
  private assetClient!: IAssetClientApi

  @inject(types.IAssetRepository)
  private assetRepository!: IAssetRepository

  @inject(types.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(assetNotation: IAssetNotation) {
    try {
      const response: IAsset = await this.assetClient.addAsset(assetNotation)
      this.assetRepository.addAsset(response)
    } catch (error) {
      this.errorHandler.handle(error).report()
    }
  }
}
