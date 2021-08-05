import { inject, provide, Ref, ref } from 'vue'
import { IPortfolioEditable } from '@/domain/portfolio/entities/PortfolioEntity'

interface IPortfolioEdited {
  (portfolioEditData: IPortfolioEditable): void
}

export const usePortfolioEditProvide = () => {
  const portfolioEdit = ref<IPortfolioEditable>()

  const portfolioEdited = (portfolioEditData: IPortfolioEditable | undefined) => {
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
  const portfolioEdit = inject('portfolioEdit') as Ref<IPortfolioEditable | undefined>
  const portfolioEdited = inject('portfolioEdited') as IPortfolioEdited

  return {
    portfolioEdit,
    portfolioEdited
  }
}
