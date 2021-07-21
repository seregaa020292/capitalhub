import AxiosMockAdapter from 'axios-mock-adapter'
import { http } from '@/infrastructure/network/http'
import userMock from '@/data/mock/userMock'
import assetMock from '@/data/mock/assetMock'

export default {
  init(): void {
    const mock = new AxiosMockAdapter(http, { delayResponse: 300 })

    userMock(mock)
    assetMock(mock)
  },
}
