<template>
  <el-main>
    <h1 class="title text-center mb-3">Мои портфели</h1>
    <el-row :gutter="12">
      <el-col
        v-for="{
          title,
          portfolioId,
          active,
          currencyTitle,
          assetAmount,
          assetQuantity,
        } in portfolios"
        :key="portfolioId"
        :md="6"
        class="mb-2"
      >
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <div @click="() => choosePortfolio(portfolioId)" class="card-title">
                <i v-if="active" class="el-icon-star-on" />
                {{ title }}
                <i class="el-icon-right" />
              </div>
              <el-dropdown trigger="click">
                <span class="el-dropdown-link">
                  <i class="el-icon-more" />
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item icon="el-icon-edit">Редактировать</el-dropdown-item>
                    <el-dropdown-item
                      @click="() => removePortfolio(portfolioId, title)"
                      icon="el-icon-delete-solid"
                    >
                      Удалить
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
          </template>
          <div class="mb-1">
            Стоимость: {{ currencyFormatter[currencyTitle.toLowerCase()].format(assetAmount) }}
          </div>
          <div>Кол-во активов: {{ assetQuantity }} шт.</div>
        </el-card>
      </el-col>
      <el-col :md="6">
        <edit-modal />
      </el-col>
    </el-row>
  </el-main>
</template>

<script lang="ts">
import { defineComponent, onMounted, computed, inject } from 'vue'
import { useRouter } from 'vue-router'
import {
  PortfolioPresenterDI,
  PortfoliosFetchUseCaseDI,
  PortfolioChooseUseCaseDI,
} from '@/domain/portfolio/module/di'
import { currencyFormatter } from '@/utils/number'
import EditModal from '@/app/view/containers/portfolio/edit/Modal.vue'
import { IConfirmService } from '@/services/message/ConfirmService'

export default defineComponent({
  name: 'Portfolio',
  components: {
    EditModal,
  },
  setup() {
    const router = useRouter()
    const $confirm = inject('$confirm') as IConfirmService
    const portfoliosFetchUseCase = PortfoliosFetchUseCaseDI()
    const portfolioPresenter = PortfolioPresenterDI()
    const portfolioChooseUseCase = PortfolioChooseUseCaseDI()
    const loading = computed(() => portfolioPresenter.loadingPortfolios())
    const portfolios = computed(() => portfolioPresenter.portfolios())

    onMounted(async () => {
      await portfoliosFetchUseCase.execute()
    })

    const choosePortfolio = async (portfolioId: string) => {
      const isChoose = await portfolioChooseUseCase.execute(portfolioId)

      if (isChoose) {
        router.push({ name: 'dashboard' })
      }
    }

    const removePortfolio = async (portfolioId: string, title: string) => {
      $confirm
        .warningVariant('Сообщение', `Подтверждаете удаление портфеля - <u>${title}</u>?`)
        .then(() => {})
        .catch(() => {})
    }

    return {
      loading,
      portfolios,
      choosePortfolio,
      removePortfolio,
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
