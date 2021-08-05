import { inject, provide, Ref, ref } from 'vue'
import { IPortfolioEditFields } from '@/domain/portfolio/entities/PortfolioEntity'

interface IPortfolioEdited {
  (portfolioEditData: IPortfolioEditFields): void
}

export const usePortfolioEditProvide = () => {
  const portfolioEdit = ref<IPortfolioEditFields>()

  const portfolioEdited = (portfolioEditData: IPortfolioEditFields | undefined) => {
    portfolioEdit.value = portfolioEditData
  }

  provide('portfolioEdit', portfolioEdit)
  provide('portfolioEdited', portfolioEdited)

  return {
    portfolioEdit,
    portfolioEdited
  }
}

export const usePortfolioEditInject = () => {
  const portfolioEdit = inject('portfolioEdit') as Ref<IPortfolioEditFields | undefined>
  const portfolioEdited = inject('portfolioEdited') as IPortfolioEdited

  return {
    portfolioEdit,
    portfolioEdited
  }
}
