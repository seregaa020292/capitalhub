type JwtData = {
  header: IHeaderJwt
  payload: IPayloadJwt
  sign: unknown
}

interface IHeaderJwt {
  alg: string
  typ: string
}

export interface IPayloadJwt {
  email: string
  id: string
  exp: number
}

export class JwtService {
  private data: JwtData

  constructor(private token: string) {
    this.data = JwtService.parseJWT(token)
  }

  public static parse(token: string): JwtService {
    return new JwtService(token)
  }

  public get header(): IHeaderJwt {
    return this.data.header
  }

  public get payload(): IPayloadJwt {
    return this.data.payload
  }

  public static parseJWT(token: string): JwtData {
    const [header, payload, sign] = token.split('.')

    return {
      header: JwtService.parsePart(header),
      payload: JwtService.parsePart(payload),
      sign: sign,
    }
  }

  public isExpired(): boolean {
    const expDate = this.data.payload.exp - 10
    const nowTime = Math.floor(new Date().getTime() / 1000)

    return expDate <= nowTime
  }

  private static parsePart(payload: string) {
    return JSON.parse(atob(payload))
  }
}
