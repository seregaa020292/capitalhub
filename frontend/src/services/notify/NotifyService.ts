import { injectable } from 'inversify'
import { ElNotification } from 'element-plus'
import { INotificationOptions } from 'element-plus/packages/notification/src/notification.type'

export interface INotifyService {
  error(message: string): void
  warning(message: string): void
  success(message: string): void
  make(notifyOptions: INotificationOptions): void
}

@injectable()
export class NotifyService implements INotifyService {
  public static instance(): NotifyService {
    return new NotifyService()
  }

  public error(message: string): void {
    return this.make({
      ...NotifyService.options(),
      type: 'error',
      message,
    })
  }

  public warning(message: string): void {
    return this.make({
      ...NotifyService.options(),
      type: 'warning',
      message,
    })
  }

  public success(message: string): void {
    return this.make({
      ...NotifyService.options(),
      type: 'success',
      message,
    })
  }

  public make(notifyOptions: INotificationOptions): void {
    ElNotification(notifyOptions)
  }

  private static options(): INotificationOptions {
    return {
      position: 'bottom-right',
      dangerouslyUseHTMLString: true,
      duration: 1e4,
    }
  }
}
