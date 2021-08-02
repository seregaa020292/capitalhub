import { http } from '@/infrastructure/network/http'
import urls from '@/infrastructure/network/urls'
import { IUser } from '@/domain/user/entities/UserEntity'
import { injectable } from 'inversify'

export interface IUserClientApi {
  fetchUser(): Promise<IUser>
}

@injectable()
export class UserClientApi implements IUserClientApi {
  fetchUser(): Promise<IUser> {
    return http.get(urls.api_v1.USER)
  }
}
