import { injectable } from 'inversify'
import { IAccessToken } from '@/domain/user/entities/UserEntity'

export interface IStorageService {
  getAccessToken(): IAccessToken | null
  getAccessTokenWithPrefix(): string
  saveAccessToken(token: IAccessToken): void
  hasAccessToken(): boolean
  removeAccessToken(): void
}

const KEY_ACCESS_TOKEN = '_access-token_'

@injectable()
export class StorageService implements IStorageService {
  getAccessToken(): IAccessToken | null {
    return this.hasAccessToken() ? JSON.parse(<string>localStorage.getItem(KEY_ACCESS_TOKEN)) : null
  }

  getAccessTokenWithPrefix(): string {
    const accessToken = this.getAccessToken()
    return accessToken !== null ? `${accessToken.prefixToken} ${accessToken.token}` : ''
  }

  saveAccessToken(token: IAccessToken): void {
    if (!!token) {
      localStorage.setItem(KEY_ACCESS_TOKEN, JSON.stringify(token))
    }
  }

  hasAccessToken(): boolean {
    return !!localStorage.getItem(KEY_ACCESS_TOKEN)
  }

  removeAccessToken(): void {
    localStorage.removeItem(KEY_ACCESS_TOKEN)
  }
}
