import { inject, injectable } from 'inversify'
import { BaseUseCase } from '@/types/domain'
import { types } from '@/domain/user/module/types'
import { IUserClientApi } from '@/domain/user/clients/api/UserClientApi'
import { IUserRepository } from '@/domain/user/repositories/UserRepository'
import { IUser } from '@/domain/user/entities/UserEntity'

export  interface IUserFetchUseCase extends BaseUseCase<unknown, Promise<void>> {}

@injectable()
export class UserFetchUseCase implements IUserFetchUseCase {
  @inject(types.IUserClientApi)
  private userClient!: IUserClientApi

  @inject(types.IUserRepository)
  private userRepository!: IUserRepository

  async execute(): Promise<void> {
    const response: IUser = await this.userClient.fetchUser()
    this.userRepository.savePersonData(response)
  }
}
