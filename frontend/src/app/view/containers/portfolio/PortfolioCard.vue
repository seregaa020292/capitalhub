<template>
  <el-card shadow="hover">
    <template #header>
      <div class="card-header">
        <div @click="portfolioChoose" class="card-title">
          <i v-if="portfolio.active" class="el-icon-star-on" />
          {{ portfolio.title }}
          <i class="el-icon-right" />
        </div>
        <el-dropdown trigger="click">
          <span class="el-dropdown-link">
            <i class="el-icon-more" />
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="portfolioEdit" icon="el-icon-edit">
                Редактировать
              </el-dropdown-item>
              <el-dropdown-item @click="portfolioRemove" icon="el-icon-delete-solid">
                Удалить
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </template>
    <div class="mb-1">
      Стоимость:
      {{ currencyFormatter[portfolio.currencyTitle.toLowerCase()].format(portfolio.assetAmount) }}
    </div>
    <div>Кол-во активов: {{ portfolio.assetQuantity }} шт.</div>
  </el-card>
</template>

<script lang="ts">
import { defineComponent, inject, PropType } from 'vue'
import { useRouter } from 'vue-router'
import { PortfolioChooseUseCaseDI, PortfolioRemoveUseCaseDI } from '@/domain/portfolio/module/di'
import { currencyFormatter } from '@/utils/number'
import { IConfirmService } from '@/services/message/ConfirmService'
import { IPortfolio, IPortfolioEditFields } from '@/domain/portfolio/entities/PortfolioEntity'

export default defineComponent({
  name: 'PortfolioCard',
  props: {
    portfolio: {
      type: Object as PropType<IPortfolio>,
      required: true,
    },
  },
  emits: ['edit-data'],
  setup(props, { emit }) {
    const router = useRouter()
    const $confirm = inject('$confirm') as IConfirmService
    const portfolioChooseUseCase = PortfolioChooseUseCaseDI()
    const portfolioRemoveUseCase = PortfolioRemoveUseCaseDI()

    const portfolioChoose = async () => {
      const isChoose = await portfolioChooseUseCase.execute(props.portfolio.portfolioId)

      if (isChoose) {
        router.push({ name: 'dashboard' })
      }
    }

    const portfolioEdit = async () => {
      const portfolioEdit: IPortfolioEditFields = {
        portfolioId: props.portfolio.portfolioId,
        currencyId: props.portfolio.currencyId,
        title: props.portfolio.title,
      }

      emit('edit-data', portfolioEdit)
    }

    const portfolioRemove = () => {
      $confirm
        .warningVariant(
          'Сообщение',
          `Подтверждаете удаление портфеля - <u>${props.portfolio.title}</u>?`
        )
        .then(async () => {
          await portfolioRemoveUseCase.execute(props.portfolio.portfolioId)
        })
        .catch(() => {})
    }

    return {
      portfolioChoose,
      portfolioRemove,
      portfolioEdit,
      currencyFormatter,
    }
  },
})
</script>

<style lang="scss" scoped>
@import '@/app/themes/element/variables';

.card-header {
  display: flex;
  justify-content: space-between;
}

.card-title {
  display: inline-block;
  cursor: pointer;
  transition: opacity 0.3s ease;

  &:hover {
    opacity: 0.5;
  }
}

.el-dropdown-link {
  cursor: pointer;
  color: $--color-primary;
  transition: color 0.3s ease;

  &:hover {
    color: $--color-primary-light;
  }
}
</style>
