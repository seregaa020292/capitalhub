import { ErrorException } from '@/infrastructure/exceptions/ErrorException'
import { HttErrorResponse } from '@/infrastructure/network/http'

export class ResponseException extends ErrorException {
  constructor(private data: HttErrorResponse) {
    super(JSON.stringify(data))
  }

  public get result(): HttErrorResponse {
    return this.data
  }
}
