export interface IUser {
  userId: string
  name: string
  email: string
  role: string
  avatar?: string
}

export type Token = string

export interface IAccessToken {
  token: Token
  prefixToken: string
}

export interface IRefreshToken {
  token: Token
}

export interface IToken {
  accessToken: IAccessToken
  refreshToken: IRefreshToken
}
