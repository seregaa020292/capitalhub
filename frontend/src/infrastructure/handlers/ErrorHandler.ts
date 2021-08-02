import { inject, injectable } from 'inversify'
import { baseTypes } from '@/infrastructure/di/types'
import { BaseHandler } from '@/types/infrastructure'
import { ILoggerService } from '@/services/logger/LoggerService'
import { IMessageService } from '@/services/message/MessageService'
import { ResponseException, ServerException } from '@/infrastructure/exceptions'

interface IHandleResult {
  report(): void
}

export interface IErrorHandler extends BaseHandler<unknown, IHandleResult> {}

@injectable()
export class ErrorHandler implements IErrorHandler {
  @inject(baseTypes.IMessageService)
  private messageService!: IMessageService

  @inject(baseTypes.ILoggerService)
  private loggerService!: ILoggerService

  public handle(error: unknown) {
    let messageError = 'Неизвестная ошибка. Повторите попытку позже. Ошибка: 000x001'
    if (error instanceof ResponseException) {
      messageError = `${error.result.status} ${error.result.error}`
    }
    if (error instanceof ServerException) {
      messageError = `Ошибка на сервере ${error.result}`
    }

    this.loggerService.report(error)

    return {
      report: () => this.messageService.error(messageError)
    }
  }
}
