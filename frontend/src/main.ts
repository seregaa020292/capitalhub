import 'reflect-metadata'
import { DIContainer } from '@/infrastructure/di'
import Main from '@/app/main'

DIContainer.get<Main>(Main)
