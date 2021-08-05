import { useRouter } from 'vue-router'
import { AuthLogoutUseCaseDI } from '@/domain/auth/module/di'

interface ILogout {
  onLogout: () => Promise<void>
}

export const useLogout = (): ILogout => {
  const router = useRouter()
  const authLogoutUseCase = AuthLogoutUseCaseDI()

  const onLogout = async () => {
    await authLogoutUseCase.execute()
    router.push({ name: 'login' })
  }

  return {
    onLogout
  }
}
