import { injectable } from 'inversify'
import { ElMessage } from 'element-plus'
import { MessageParams } from 'element-plus/packages/message/src/types'

export interface IMessageService {
  error(message: string): void
  warning(message: string): void
  success(message: string): void
  make(messageParams: MessageParams): void
}

@injectable()
export class MessageService implements IMessageService {
  public static instance(): MessageService {
    return new MessageService()
  }

  public error(message: string): void {
    return this.make({ type: 'error', message })
  }

  public warning(message: string): void {
    return this.make({ type: 'warning', message })
  }

  public success(message: string): void {
    return this.make({ type: 'success', message })
  }

  public make(messageParams: MessageParams): void {
    ElMessage(messageParams)
  }
}
