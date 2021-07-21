import { injectable } from 'inversify'

export interface ILoggerService {
  report(message: any): void
}

@injectable()
export class LoggerService implements ILoggerService {
  report(message: any): void {
    console.info(message)
  }
}
