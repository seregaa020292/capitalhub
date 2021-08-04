import { injectable } from 'inversify'
import { http } from '@/infrastructure/network/http'
import urls from '@/infrastructure/network/urls'
import { IDashboard } from '@/domain/application/entities/ApplicationEntity'

export interface IApplicationClientApi {
  dashboard(): Promise<IDashboard>
}

@injectable()
export class ApplicationClientApi implements IApplicationClientApi {
  dashboard(): Promise<IDashboard> {
    return http.get(urls.api_v1.APP_DASHBOARD)
  }
}
