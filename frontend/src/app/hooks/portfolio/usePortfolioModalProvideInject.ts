import {
  useModalHandleInject,
  useModalHandleProvide,
} from '@/app/hooks/common/useModalHandleProvideInject'

export const usePortfolioModalProvide = () => useModalHandleProvide('portfolio')
export const usePortfolioModalInject = () => useModalHandleInject('portfolio')
