import { ErrorException } from '@/infrastructure/exceptions/ErrorException'

export class ServerException extends ErrorException {
  constructor(private code: number) {
    super(code.toString())
  }

  public get result(): number {
    return this.code
  }
}
