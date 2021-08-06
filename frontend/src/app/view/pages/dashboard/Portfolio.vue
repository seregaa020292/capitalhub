<template>
  <el-main>
    <h1 class="title text-center mb-3">Мои портфели</h1>
    <el-row :gutter="12">
      <transition-group v-if="portfolios.length" name="list">
        <el-col v-for="portfolio in portfolios" :key="portfolio.portfolioId" :md="6" class="mb-2">
          <portfolio-card :portfolio="portfolio" @editData="portfolioFormEditOpen" />
        </el-col>
      </transition-group>
      <el-col :md="6">
        <el-button @click="portfolioFormAddOpen" type="primary">
          <i class="el-icon-circle-plus" />
          Добавить еще портфель
        </el-button>
      </el-col>
    </el-row>
    <edit-modal />
  </el-main>
</template>

<script lang="ts">
import { defineComponent, onMounted, computed } from 'vue'
import { PortfolioPresenterDI, PortfoliosFetchUseCaseDI } from '@/domain/portfolio/module/di'
import EditModal from '@/app/view/containers/portfolio/edit/Modal.vue'
import PortfolioCard from '@/app/view/containers/portfolio/PortfolioCard.vue'
import { usePortfolioModalProvide } from '@/app/hooks/portfolio/usePortfolioModalProvideInject'
import { usePortfolioEditProvide } from '@/app/hooks/portfolio/usePortfolioEditProvideInject'
import { IPortfolioEditFields } from '@/domain/portfolio/entities/PortfolioEntity'

export default defineComponent({
  name: 'Portfolio',
  components: {
    EditModal,
    PortfolioCard,
  },
  setup() {
    const portfoliosFetchUseCase = PortfoliosFetchUseCaseDI()
    const portfolioPresenter = PortfolioPresenterDI()
    const { dialogOpenHandle } = usePortfolioModalProvide()
    const { portfolioEdited } = usePortfolioEditProvide()
    const loading = computed(() => portfolioPresenter.loadingPortfolios())
    const portfolios = computed(() => portfolioPresenter.portfolios())

    onMounted(async () => {
      await portfoliosFetchUseCase.execute()
    })

    const portfolioFormAddOpen = () => {
      portfolioEdited(undefined)
      dialogOpenHandle()
    }

    const portfolioFormEditOpen = (portfolioEdit: IPortfolioEditFields) => {
      portfolioEdited(portfolioEdit)
      dialogOpenHandle()
    }

    return {
      loading,
      portfolios,
      portfolioFormAddOpen,
      portfolioFormEditOpen,
    }
  },
})
</script>
