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
            <div @click="() => choosePortfolio(portfolioId)" class="card-header">
              <i v-if="active" class="el-icon-star-on" />
              {{ title }}
              <i class="el-icon-right" />
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
import { defineComponent, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import {
  PortfolioPresenterDI,
  PortfoliosFetchUseCaseDI,
  PortfolioChooseUseCaseDI,
} from '@/domain/portfolio/module/di'
import { currencyFormatter } from '@/utils/number'
import EditModal from '@/app/view/containers/portfolio/edit/Modal.vue'

export default defineComponent({
  name: 'Portfolio',
  components: {
    EditModal,
  },
  setup() {
    const router = useRouter()
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

    return {
      loading,
      portfolios,
      choosePortfolio,
      currencyFormatter,
    }
  },
})
</script>

<style lang="scss" scoped>
.card-header {
  display: inline-block;
  cursor: pointer;
  transition: opacity 0.3s ease;

  &:hover {
    opacity: 0.5;
  }
}
</style>
