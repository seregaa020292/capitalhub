import { injectable } from 'inversify'
import { ElMessageBox } from 'element-plus'
import { ElMessageBoxOptions } from 'element-plus/lib/el-message-box/src/message-box.type'
import { MessageBoxData } from 'element-plus/packages/message-box/src/message-box.type'

export interface IConfirmService {
  warningVariant(title: string, message: string): Promise<MessageBoxData>
  make(title: string, message: string, options: ElMessageBoxOptions): Promise<MessageBoxData>
}

@injectable()
export class ConfirmService implements IConfirmService {
  public static instance(): ConfirmService {
    return new ConfirmService()
  }

  public warningVariant(title: string, message: string): Promise<MessageBoxData> {
    return this.make(title, message, {
      confirmButtonText: 'Да',
      cancelButtonText: 'Нет',
      type: 'warning',
      dangerouslyUseHTMLString: true,
    })
  }

  public make(title: string, message: string, options: ElMessageBoxOptions) {
    return ElMessageBox.confirm(message, title, options)
  }
}
