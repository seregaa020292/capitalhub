import { DIContainer } from '@/infrastructure/di'
import { types } from '@/domain/user/module/types'
import { IUserPresenter } from '@/domain/user/presenters/UserPresenter'

export const UserPresenterDI = (): IUserPresenter => DIContainer.get(types.IUserPresenter)
