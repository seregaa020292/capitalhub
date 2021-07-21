import { injectable } from 'inversify'
import { http, httpEasy, HttpResponse } from '@/infrastructure/network/http'
import urls from '@/infrastructure/network/urls'
import { IToken, IUser } from '@/domain/user/entities/UserEntity'
import { responseWithCSRF } from '@/utils/server'

export interface ILoginParams {
  email: string
  password: string
}

export interface IRegisterParams extends ILoginParams {
  name: string
}

export interface IAuthResponse extends IUserResponse, IToken {}

export interface CSRF {
  csrf: string
}

interface IUserResponse {
  user: IUser
}

interface ILogin extends ILoginParams {
  fingerprint: string
}

interface IRefreshParams {
  fingerprint: string
}

export interface IAuthClientApi {
  login(credentials: ILogin): Promise<IAuthResponse & CSRF>
  register(credentials: IRegisterParams): Promise<void>
  refresh(credentials: IRefreshParams): Promise<IToken & CSRF>
  logout(): Promise<HttpResponse>
  checkLogged(): Promise<IUserResponse & CSRF>
  confirmed(code: string): Promise<void>
}

@injectable()
export class AuthClientApi implements IAuthClientApi {
  login(credentials: ILogin): Promise<IAuthResponse & CSRF> {
    return http.post(urls.api_v1.LOGIN, credentials, responseWithCSRF())
  }

  register(credentials: IRegisterParams): Promise<void> {
    return http.post(urls.api_v1.REGISTER, credentials)
  }

  refresh(credentials: IRefreshParams): Promise<IToken & CSRF> {
    return httpEasy.post(urls.api_v1.REFRESH_TOKEN, credentials, responseWithCSRF())
  }

  logout(): Promise<HttpResponse> {
    return httpEasy.post(urls.api_v1.LOGOUT)
  }

  checkLogged(): Promise<IUserResponse & CSRF> {
    return http.get(urls.api_v1.CHECK_LOGGED, responseWithCSRF())
  }

  confirmed(code: string): Promise<void> {
    return http.get(urls.api_v1.CONFIRMED, { params: { code } })
  }
}
