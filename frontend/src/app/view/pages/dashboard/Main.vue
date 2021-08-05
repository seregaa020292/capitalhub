<template>
  <el-main>
    <h1 class="title text-center mb-3">{{ portfolio.title }}</h1>
    <div class="table-wrapper">
      <el-row :gutter="10" class="table-menu">
        <el-col :md="16">
          <router-link to="/dashboard/asset-edit" custom v-slot="{ navigate }">
            <el-button type="primary" @click="navigate" icon="el-icon-circle-plus-outline">
              Добавить актив
            </el-button>
          </router-link>
        </el-col>
        <el-col :md="8">
          <el-row :gutter="5">
            <el-col :md="16">
              <el-input
                v-model="state.search"
                type="search"
                placeholder="Поиск по названию"
                class="w-100"
              />
            </el-col>
            <el-col :md="8">
              <el-button icon="el-icon-s-operation" class="w-100">Фильтры</el-button>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
      <br />
      <asset-table />
    </div>
    <router-view name="AssetEdit" />
  </el-main>
</template>

<script lang="ts">
import { defineComponent, reactive, onUnmounted, onMounted, computed } from 'vue'
import AssetTable from '@/app/view/containers/dashboard/AssetTable.vue'
import QuoteClientSocket from '@/services/socket/QuoteClientSocket'
import { PortfolioFetchUseCaseDI, PortfolioPresenterDI } from '@/domain/portfolio/module/di'

export default defineComponent({
  name: 'Main',
  components: {
    AssetTable,
  },
  setup: () => {
    const state = reactive({
      search: '',
    })
    const portfolioFetchUseCase = PortfolioFetchUseCaseDI()
    const portfolioPresenter = PortfolioPresenterDI()
    const socket = new QuoteClientSocket()

    const portfolio = computed(() => portfolioPresenter.portfolio())

    onMounted(async () => {
      await portfolioFetchUseCase.execute()
      socket.subscribe()
    })

    onUnmounted(() => {
      socket.terminate()
    })

    return {
      state,
      portfolio,
    }
  },
})
</script>

<style lang="scss" scoped>
@import '@/app/themes/element/variables';

.table-wrapper {
  background-color: $--color-main-darker;
  border-radius: 8px;
  overflow: hidden;
}

.table-menu {
  padding: 15px 10px 0;
}
</style>
