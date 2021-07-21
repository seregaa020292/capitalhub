import { MutationTree } from 'vuex'
import { IAsset, IAssetPrice } from '@/domain/asset/entities/AssetEntity'

export interface IAssetState {
  assets: IAsset[]
}

/**
 ******************************
 * @State
 ******************************
 */
export const state = (): IAssetState => ({
  assets: [],
})

/**
 ******************************
 * @Mutation
 ******************************
 */
export enum Types {
  FETCH_ASSETS = 'asset/fetch',
  CLEAR_ASSETS = 'asset/clear',
  ASSET_ADD = 'asset/add',
  CHANGE_PRICE = 'asset/change_price',
}

export class FetchAssets implements FluxStandardAction {
  public type = Types.FETCH_ASSETS
  constructor(public payload: IAsset[]) {}
}

export class ClearAssets implements FluxStandardAction {
  public type = Types.CLEAR_ASSETS
}

export class AssetAdd implements FluxStandardAction {
  public type = Types.ASSET_ADD
  constructor(public payload: IAsset) {}
}

export class ChangePrice implements FluxStandardAction {
  public type = Types.CHANGE_PRICE
  constructor(public payload: IAssetPrice) {}
}

export const mutations: MutationTree<IAssetState> = {
  [Types.FETCH_ASSETS]: (state, action: FetchAssets) => {
    state.assets = action.payload
  },
  [Types.CLEAR_ASSETS]: (state) => {
    state.assets = []
  },
  [Types.ASSET_ADD]: (state, action: AssetAdd) => {
    const indexAsset = state.assets.findIndex((asset) => asset.marketId === action.payload.marketId)
    indexAsset === -1
      ? state.assets.push(action.payload)
      : state.assets[indexAsset] = action.payload
  },
  [Types.CHANGE_PRICE]: (state, action: ChangePrice) => {
    const indexAsset = state.assets.findIndex((asset) => asset.identify === action.payload.identify)
    if (indexAsset !== -1) {
      state.assets[indexAsset].currentPrice = action.payload.currentPrice
    }
  }
}
