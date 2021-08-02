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
      >
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <i v-if="active" class="el-icon-star-on" />
              {{ title }}
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
import { PortfolioPresenterDI, PortfoliosFetchUseCaseDI } from '@/domain/portfolio/module/di'
import { currencyFormatter } from '@/utils/number'
import EditModal from '@/app/view/containers/portfolio/edit/Modal.vue'

export default defineComponent({
  name: 'Portfolio',
  components: {
    EditModal,
  },
  setup() {
    const portfoliosFetchUseCase = PortfoliosFetchUseCaseDI()
    const portfolioPresenter = PortfolioPresenterDI()

    const loading = computed(() => portfolioPresenter.loadingPortfolios())
    const portfolios = computed(() => portfolioPresenter.portfolios())

    onMounted(async () => {
      await portfoliosFetchUseCase.execute()
    })

    return {
      loading,
      portfolios,
      currencyFormatter,
    }
  },
})
</script>
