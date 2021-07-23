import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import types from '@/infrastructure/di/types'
import { IAssetRepository } from '@/domain/asset/repositories/AssetRepository'
import { IAsset, IAssetNotation } from '@/domain/asset/entities/AssetEntity'
import { IAssetClientApi } from '@/services/api/AssetClientApi'
import { IErrorHandler } from '@/infrastructure/handlers/ErrorHandler'

export  interface IAssetAddUseCase extends BaseUseCase<IAssetNotation, Promise<boolean>> {}

@injectable()
export class AssetAddUseCase implements IAssetAddUseCase {
  @inject(types.IAssetClientApi)
  private assetClient!: IAssetClientApi

  @inject(types.IAssetRepository)
  private assetRepository!: IAssetRepository

  @inject(types.IErrorHandler)
  private errorHandler!: IErrorHandler

  async execute(assetNotation: IAssetNotation): Promise<boolean> {
    try {
      const response: IAsset = await this.assetClient.addAsset(assetNotation)
      this.assetRepository.addAsset(response)
      return true
    } catch (error) {
      this.errorHandler.handle(error).report()
      return false
    }
  }
}
